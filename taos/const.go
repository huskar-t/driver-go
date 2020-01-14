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
	statusInTrans StatusFlag = 1 << iota
	statusInAutocommit
	statusReserved // Not in documentation
	StatusMoreResultsExists
	statusNoGoodIndexUsed
	statusNoIndexUsed
	statusCursorExists
	statusLastRowSent
	statusDbDropped
	StatusNoBackslashEscapes
	statusMetadataChanged
	statusQueryWasSlow
	statusPsOutParams
	statusInTransReadonly
	statusSessionStateChanged
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
