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

import (
	"errors"
	"log"
	"os"
)

// Various errors the driver might return.
var (
	ErrInvalidConn = errors.New("invalid connection")
	ErrConnNoExist = errors.New("no existent connection ")
)

type Logger interface {
	Print(v ...interface{})
}

var errLog = Logger(log.New(os.Stderr, "[taos] ", log.Ldate|log.Ltime|log.Lshortfile))

func SetLogger(logger Logger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}
	errLog = logger
	return nil
}
