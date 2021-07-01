/*
Copyright © 2021 Jarvib Ding <ding@ibyte.me>
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package core

import (
	"database/sql"
	"fmt"
)

type DBInfo struct {
	DBType   string
	HostIP   string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

type DBMode struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

func NewDB(info *DBInfo) *DBMode {
	return &DBMode{
		DBInfo: info,
	}
}

func (db *DBMode) Connection() error {
	var err error
	str := "%s:%s@tcp(%s)/infomation_schema?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(str,
		db.DBInfo.UserName,
		db.DBInfo.Password,
		db.DBInfo.DBType,
		db.DBInfo.Charset,
	)
	// open sql connection
	db.DBEngine, err = sql.Open(db.DBInfo.DBType, dsn)
	return err
}
