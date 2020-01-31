/*
 * Copyright (c) 2019 TAOS Data, Inc. <jhtao@taosdata.com>
 *
 * This program is free software: you can use, redistribute, and/or modify
 * it under the terms of the GNU Affero General Public License, version 3
 * or later ("AGPL"), as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */
package driver

/*
#cgo CFLAGS : -I/usr/include
#cgo LDFLAGS: -L/usr/lib -ltaos
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <taos.h>
*/
import "C"

import (
	"context"
	"crypto/md5"
	"database/sql/driver"
	"errors"
	"github.com/taosdata/driver-go/taos"
	"net"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type SConnectMsg struct {
	ClientVersion string
	DB            string
}

func (tc *taosConn) taosConnect(ctx context.Context) error {
	//check ip
	ip := tc.cfg.addr
	user := tc.cfg.user
	pass := tc.cfg.password
	db := tc.cfg.dbName
	if ip == "" || ip == "127.0.0.1" || ip == "localhost" {
		tc.cfg.addr = "127.0.0.1"
	}
	//check user
	if user == "" || len(user) > taos.TSDB_USER_LEN {
		return taos.GetErrorStr(taos.TSDB_CODE_INVALID_ACCT)
	}
	//check pass
	if pass == "" || len(pass) > taos.TSDB_PASSWORD_LEN {
		return taos.GetErrorStr(taos.TSDB_CODE_INVALID_ACCT)
	}
	//md5 pass
	h := md5.New()
	h.Write([]byte(pass))
	encodePassword := h.Sum(nil)
	if len(encodePassword) != taos.TSDB_AUTH_LEN {
		return taos.GetErrorStr(taos.TSDB_CODE_INVALID_ACCT)
	}
	//check port
	if tc.cfg.port == 0 {
		tc.cfg.port = 6030
	}
	//check db
	if db != "" {
		if len(db) > taos.TSDB_DB_NAME_LEN {
			return taos.GetErrorStr(taos.TSDB_CODE_INVALID_DB)
		}
		db = strings.ToLower(db)
		tc.cfg.dbName = db
	}
	tc.rpcProtocol = taos.NewRpcProtocol(user, db, encodePassword)
	// Connect to Server
	nd := net.Dialer{}
	var err error
	//tc.netConn, err = nd.DialContext(ctx, tc.cfg.net, fmt.Sprintf("%s:%d",tc.cfg.addr, &tc.cfg.port))
	tc.netConn, err = nd.DialContext(ctx, "udp", net.JoinHostPort(tc.cfg.addr, strconv.Itoa(tc.cfg.port)))
	if err != nil {
		return err
	}
	data, err := tc.rpcProtocol.GetReqMsg(taos.TSDB_SQL_CONNECT, nil)
	if err != nil {
		return err
	}
	_, err = tc.netConn.Write(data)
	return err
}

func (tc *taosConn) taosQuery(sqlStr string) (int, error) {
	cSqlStr := C.CString(sqlStr)
	defer C.free(unsafe.Pointer(cSqlStr))
	code := int(C.taos_query(tc.netConn, cSqlStr))

	if 0 != code {
		tc.taosError()
		errStr := C.GoString(C.taos_errstr(tc.netConn))
		errLog.Print("taos_query() failed:", errStr)
		errLog.Print("taosQuery() input sql:%s\n", sqlStr)
		return 0, errors.New(errStr)
	}

	// read result and save into tc struct
	numFields := int(C.taos_field_count(tc.netConn))
	if 0 == numFields { // there are no select and show kinds of commands
		tc.affectedRows = int(C.taos_affected_rows(tc.netConn))
		tc.insertId = 0
	}

	return numFields, nil
}

func (tc *taosConn) taosError() {
	// free local resouce: allocated memory/metric-meta refcnt
	//var pRes unsafe.Pointer
	pRes := C.taos_use_result(tc.netConn)
	C.taos_free_result(pRes)
}

func (tc *taosConn) heartbeat() {
	heartbeatTicker := time.NewTicker(time.Millisecond * 500 * 3)
	for range heartbeatTicker.C {
		if tc.netConn == nil {
			errLog.Print(driver.ErrBadConn)
			return
		}
		data, err := tc.rpcProtocol.GetReqMsg(taos.TSDB_SQL_HB, nil)
		if err != nil {
			return
		}
		_, _ = tc.netConn.Write(data)
	}
}
