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
	TSDB_TRUE  = 1
	TSDB_FALSE = 0
	TSDB_OK    = 0
	TSDB_ERR   = -1

	TS_PATH_DELIMITER = "."

	TSDB_TIME_PRECISION_MILLI = 0
	TSDB_TIME_PRECISION_MICRO = 1

	TSDB_TIME_PRECISION_MILLI_STR = "ms"
	TSDB_TIME_PRECISION_MICRO_STR = "us"

	TSDB_DATA_TYPE_BOOL      = 1  // 1 bytes
	TSDB_DATA_TYPE_TINYINT   = 2  // 1 byte
	TSDB_DATA_TYPE_SMALLINT  = 3  // 2 bytes
	TSDB_DATA_TYPE_INT       = 4  // 4 bytes
	TSDB_DATA_TYPE_BIGINT    = 5  // 8 bytes
	TSDB_DATA_TYPE_FLOAT     = 6  // 4 bytes
	TSDB_DATA_TYPE_DOUBLE    = 7  // 8 bytes
	TSDB_DATA_TYPE_BINARY    = 8  // string
	TSDB_DATA_TYPE_TIMESTAMP = 9  // 8 bytes
	TSDB_DATA_TYPE_NCHAR     = 10 // unicode string

	TSDB_RELATION_INVALID     = 0
	TSDB_RELATION_LESS        = 1
	TSDB_RELATION_LARGE       = 2
	TSDB_RELATION_EQUAL       = 3
	TSDB_RELATION_LESS_EQUAL  = 4
	TSDB_RELATION_LARGE_EQUAL = 5
	TSDB_RELATION_NOT_EQUAL   = 6
	TSDB_RELATION_LIKE        = 7

	TSDB_RELATION_AND = 8
	TSDB_RELATION_OR  = 9
	TSDB_RELATION_NOT = 10

	TSDB_BINARY_OP_ADD       = 11
	TSDB_BINARY_OP_SUBTRACT  = 12
	TSDB_BINARY_OP_MULTIPLY  = 13
	TSDB_BINARY_OP_DIVIDE    = 14
	TSDB_BINARY_OP_REMAINDER = 15
	TSDB_USERID_LEN          = 9
	TS_PATH_DELIMITER_LEN    = 1

	TSDB_METER_ID_LEN_MARGIN = 10
	TSDB_METER_ID_LEN        = TSDB_DB_NAME_LEN + TSDB_METER_NAME_LEN + 2*TS_PATH_DELIMITER_LEN + TSDB_USERID_LEN + TSDB_METER_ID_LEN_MARGIN //TSDB_DB_NAME_LEN+TSDB_METER_NAME_LEN+2*strlen(TS_PATH_DELIMITER)+strlen(USERID)
	TSDB_UNI_LEN             = 24
	TSDB_USER_LEN            = TSDB_UNI_LEN
	TSDB_ACCT_LEN            = TSDB_UNI_LEN
	TSDB_PASSWORD_LEN        = TSDB_UNI_LEN

	TSDB_MAX_COLUMNS = 256
	TSDB_MIN_COLUMNS = 2 //PRIMARY COLUMN(timestamp) + other columns

	TSDB_METER_NAME_LEN      = 64
	TSDB_DB_NAME_LEN         = 32
	TSDB_COL_NAME_LEN        = 64
	TSDB_MAX_SAVED_SQL_LEN   = TSDB_MAX_COLUMNS * 16
	TSDB_MAX_SQL_LEN         = TSDB_PAYLOAD_SIZE
	TSDB_MAX_ALLOWED_SQL_LEN = 8 * 1024 * 1024 // sql length should be less than 6mb

	TSDB_MAX_BYTES_PER_ROW = TSDB_MAX_COLUMNS * 16
	TSDB_MAX_TAGS_LEN      = 512
	TSDB_MAX_TAGS          = 32

	TSDB_AUTH_LEN       = 16
	TSDB_KEY_LEN        = 16
	TSDB_VERSION_LEN    = 12
	TSDB_STREET_LEN     = 64
	TSDB_CITY_LEN       = 20
	TSDB_STATE_LEN      = 20
	TSDB_COUNTRY_LEN    = 20
	TSDB_VNODES_SUPPORT = 6
	TSDB_MGMT_SUPPORT   = 4
	TSDB_LOCALE_LEN     = 64
	TSDB_TIMEZONE_LEN   = 64

	TSDB_IPv4ADDR_LEN     = 16
	TSDB_FILENAME_LEN     = 128
	TSDB_METER_VNODE_BITS = 20
	TSDB_METER_SID_MASK   = 0xFFFFF
	TSDB_SHELL_VNODE_BITS = 24
	TSDB_SHELL_SID_MASK   = 0xFF
	TSDB_HTTP_TOKEN_LEN   = 20
	TSDB_SHOW_SQL_LEN     = 32

	TSDB_METER_STATE_OFFLINE = 0
	TSDB_METER_STATE_ONLLINE = 1

	TSDB_DEFAULT_PKT_SIZE = 65480 //same as RPC_MAX_UDP_SIZE

	TSDB_PAYLOAD_SIZE         = (TSDB_DEFAULT_PKT_SIZE - 100)
	TSDB_DEFAULT_PAYLOAD_SIZE = 1024 // default payload size
	TSDB_EXTRA_PAYLOAD_SIZE   = 128  // extra bytes for auth
	TSDB_SQLCMD_SIZE          = 1024
	TSDB_MAX_VNODES           = 256
	TSDB_MIN_VNODES           = 50
	TSDB_INVALID_VNODE_NUM    = 0

	TSDB_DNODE_ROLE_ANY   = 0
	TSDB_DNODE_ROLE_MGMT  = 1
	TSDB_DNODE_ROLE_VNODE = 2

	TSDB_MAX_MPEERS   = 5
	TSDB_MAX_MGMT_IPS = (TSDB_MAX_MPEERS + 1)

	TSDB_REPLICA_MIN_NUM = 1
	/*
	 * this is defined in CMakeList.txt
	 */
	//TSDB_REPLICA_MAX_NUM      3

	TSDB_TBNAME_COLUMN_INDEX     = (-1)
	TSDB_MULTI_METERMETA_MAX_NUM = 100000 // maximum batch size allowed to load metermeta

	//default value == 10
	TSDB_FILE_MIN_PARTITION_RANGE = 1    //minimum partition range of vnode file in days
	TSDB_FILE_MAX_PARTITION_RANGE = 3650 //max partition range of vnode file in days

	TSDB_DATA_MIN_RESERVE_DAY     = 1    // data in db to be reserved.
	TSDB_DATA_DEFAULT_RESERVE_DAY = 3650 // ten years

	TSDB_MIN_COMPRESSION_LEVEL = 0
	TSDB_MAX_COMPRESSION_LEVEL = 2

	TSDB_MIN_COMMIT_TIME_INTERVAL = 30
	TSDB_MAX_COMMIT_TIME_INTERVAL = 40960

	TSDB_MIN_ROWS_IN_FILEBLOCK = 200
	TSDB_MAX_ROWS_IN_FILEBLOCK = 500000

	TSDB_MIN_CACHE_BLOCK_SIZE = 100
	TSDB_MAX_CACHE_BLOCK_SIZE = 104857600

	TSDB_MIN_CACHE_BLOCKS = 100
	TSDB_MAX_CACHE_BLOCKS = 409600

	TSDB_MIN_AVG_BLOCKS     = 2
	TSDB_MAX_AVG_BLOCKS     = 2048
	TSDB_DEFAULT_AVG_BLOCKS = 4

	/*
	 * There is a bug in function taosAllocateId.
	 * When "create database tables 1" is executed, the wrong sid is assigned, so the minimum value is set to 2.
	 */
	TSDB_MIN_TABLES_PER_VNODE = 2
	TSDB_MAX_TABLES_PER_VNODE = 220000

	TSDB_MAX_JOIN_TABLE_NUM = 5

	//TSDB_MAX_BINARY_LEN            = (TSDB_MAX_BYTES_PER_ROW - TSDB_KEYSIZE)
	//TSDB_MAX_NCHAR_LEN             = (TSDB_MAX_BYTES_PER_ROW - TSDB_KEYSIZE)
	PRIMARYKEY_TIMESTAMP_COL_INDEX = 0

	TSDB_DATA_BOOL_NULL     = 0x02
	TSDB_DATA_TINYINT_NULL  = 0x80
	TSDB_DATA_SMALLINT_NULL = 0x8000
	TSDB_DATA_INT_NULL      = 0x80000000
	TSDB_DATA_BIGINT_NULL   = 0x8000000000000000

	TSDB_DATA_FLOAT_NULL  = 0x7FF00000         // it is an NAN
	TSDB_DATA_DOUBLE_NULL = 0x7FFFFF0000000000 // an NAN
	TSDB_DATA_NCHAR_NULL  = 0xFFFFFFFF
	TSDB_DATA_BINARY_NULL = 0xFF

	TSDB_DATA_NULL_STR   = "NULL"
	TSDB_DATA_NULL_STR_L = "null"

	TSDB_MAX_RPC_THREADS = 5

	TSDB_QUERY_TYPE_QUERY         = 0    // normal query
	TSDB_QUERY_TYPE_FREE_RESOURCE = 0x01 // free qhandle at vnode

	/*
	 * 1. ordinary sub query for select * from super_table
	 * 2. all sqlobj generated by createSubqueryObj with this flag
	 */
	TSDB_QUERY_TYPE_SUBQUERY        = 0x02
	TSDB_QUERY_TYPE_STABLE_SUBQUERY = 0x04 // two-stage subquery for super table

	TSDB_QUERY_TYPE_TABLE_QUERY      = 0x08 // query ordinary table; below only apply to client side
	TSDB_QUERY_TYPE_STABLE_QUERY     = 0x10 // query on super table
	TSDB_QUERY_TYPE_JOIN_QUERY       = 0x20 // join query
	TSDB_QUERY_TYPE_PROJECTION_QUERY = 0x40 // select *,columns... query
	TSDB_QUERY_TYPE_JOIN_SEC_STAGE   = 0x80 // join sub query at the second stage

	TSQL_SO_ASC  = 1
	TSQL_SO_DESC = 0
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
