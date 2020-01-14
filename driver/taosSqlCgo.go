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
	"errors"
	"github.com/taosdata/driver-go/taos"
	"net"
	"strconv"
	"strings"
	"unsafe"
)

type SConnectMsg struct {
	ClientVersion string
	DB            string
}

func (mc *taosConn) taosConnect(ctx context.Context) error {
	//check ip
	ip := mc.cfg.addr
	user := mc.cfg.user
	pass := mc.cfg.password
	db := mc.cfg.dbName
	if ip == "" || ip == "127.0.0.1" || ip == "localhost" {
		mc.cfg.addr = "127.0.0.1"
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

	//check port
	if mc.cfg.port == 0 {
		mc.cfg.port = 6030
	}
	//check db
	if db != "" {
		if len(db) > taos.TSDB_DB_NAME_LEN {
			return taos.GetErrorStr(taos.TSDB_CODE_INVALID_DB)
		}
		db = strings.ToLower(db)
		mc.cfg.dbName = db
	}
	mc.rpcProtocol = taos.NewRpcProtocol(user, db, encodePassword)
	// Connect to Server
	nd := net.Dialer{}
	var err error
	//mc.netConn, err = nd.DialContext(ctx, mc.cfg.net, fmt.Sprintf("%s:%d",mc.cfg.addr, &mc.cfg.port))
	mc.netConn, err = nd.DialContext(ctx, "udp", net.JoinHostPort(mc.cfg.addr, strconv.Itoa(mc.cfg.port)))
	if err != nil {
		return err
	}
	data, err := mc.rpcProtocol.GetReqMsg("connect", nil)
	if err != nil {
		return err
	}
	_, err = mc.netConn.Write(data)
	return err
}

func (mc *taosConn) taosQuery(sqlstr string) (int, error) {
	//taosLog.Printf("taosQuery() input sql:%s\n", sqlstr)

	csqlstr := C.CString(sqlstr)
	defer C.free(unsafe.Pointer(csqlstr))
	code := int(C.taos_query(mc.taos, csqlstr))

	if 0 != code {
		mc.taosError()
		errStr := C.GoString(C.taos_errstr(mc.taos))
		taos.Log.Println("taos_query() failed:", errStr)
		taos.Log.Printf("taosQuery() input sql:%s\n", sqlstr)
		return 0, errors.New(errStr)
	}

	// read result and save into mc struct
	numFields := int(C.taos_field_count(mc.taos))
	if 0 == numFields { // there are no select and show kinds of commands
		mc.affectedRows = int(C.taos_affected_rows(mc.taos))
		mc.insertId = 0
	}

	return numFields, nil
}

func (mc *taosConn) taosClose() {
	C.taos_close(mc.taos)
}

func (mc *taosConn) taosError() {
	// free local resouce: allocated memory/metric-meta refcnt
	//var pRes unsafe.Pointer
	pRes := C.taos_use_result(mc.taos)
	C.taos_free_result(pRes)
}
