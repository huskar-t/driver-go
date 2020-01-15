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

package taos

const (
	TimeFormat     = "2006-01-02 15:04:05"
	MaxTaosSqlLen  = 65380
	DefaultBufSize = MaxTaosSqlLen + 32
)

type FieldType byte

type FieldFlag uint16

const (
	FlagNotNULL FieldFlag = 1 << iota
)

type StatusFlag uint16

const (
	StatusInTrans StatusFlag = 1 << iota
	StatusInAutocommit
	StatusReserved // Not in documentation
	StatusMoreResultsExists
	StatusNoGoodIndexUsed
	StatusNoIndexUsed
	StatusCursorExists
	StatusLastRowSent
	StatusDbDropped
	StatusNoBackslashEscapes
	StatusMetadataChanged
	StatusQueryWasSlow
	StatusPsOutParams
	StatusInTransReadonly
	StatusSessionStateChanged
)
const (
	TSDB_UNI_LEN        = 24
	TSDB_USER_LEN       = TSDB_UNI_LEN
	TSDB_ACCT_LEN       = TSDB_UNI_LEN
	TSDB_PASSWORD_LEN   = TSDB_UNI_LEN
	TSDB_METER_NAME_LEN = 64
	TSDB_DB_NAME_LEN    = 32
	TSDB_COL_NAME_LEN   = 64
)

const (
	TSDB_SQL_SELECT = iota
	TSDB_SQL_FETCH
	TSDB_SQL_INSERT

	TSDB_SQL_MGMT // the SQL below is for mgmt node
	TSDB_SQL_CREATE_DB
	TSDB_SQL_CREATE_TABLE
	TSDB_SQL_DROP_DB
	TSDB_SQL_DROP_TABLE
	TSDB_SQL_CREATE_ACCT
	TSDB_SQL_CREATE_USER
	TSDB_SQL_DROP_ACCT // 10
	TSDB_SQL_DROP_USER
	TSDB_SQL_ALTER_USER
	TSDB_SQL_ALTER_ACCT
	TSDB_SQL_ALTER_TABLE
	TSDB_SQL_ALTER_DB
	TSDB_SQL_CREATE_MNODE
	TSDB_SQL_DROP_MNODE
	TSDB_SQL_CREATE_DNODE
	TSDB_SQL_DROP_DNODE
	TSDB_SQL_CFG_DNODE // 20
	TSDB_SQL_CFG_MNODE
	TSDB_SQL_SHOW
	TSDB_SQL_RETRIEVE
	TSDB_SQL_KILL_QUERY
	TSDB_SQL_KILL_STREAM
	TSDB_SQL_KILL_CONNECTION

	TSDB_SQL_READ // SQL below is for read operation
	TSDB_SQL_CONNECT
	TSDB_SQL_USE_DB
	TSDB_SQL_META // 30
	TSDB_SQL_METRIC
	TSDB_SQL_MULTI_META
	TSDB_SQL_HB

	TSDB_SQL_LOCAL // SQL below for client local
	TSDB_SQL_DESCRIBE_TABLE
	TSDB_SQL_RETRIEVE_METRIC
	TSDB_SQL_METRIC_JOIN_RETRIEVE
	TSDB_SQL_RETRIEVE_TAGS
	/*
	 * build empty result instead of accessing dnode to fetch result
	 * reset the client cache
	 */
	TSDB_SQL_RETRIEVE_EMPTY_RESULT

	TSDB_SQL_RESET_CACHE // 40
	TSDB_SQL_SERV_STATUS
	TSDB_SQL_CURRENT_DB
	TSDB_SQL_SERV_VERSION
	TSDB_SQL_CLI_VERSION
	TSDB_SQL_CURRENT_USER
	TSDB_SQL_CFG_LOCAL

	TSDB_SQL_MAX
)
