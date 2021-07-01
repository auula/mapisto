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

// MEngine is mapisto database mapper engine
type MEngine struct {
	Tpl Template
	*DataBase
}

func (me *MEngine) SetDB(db *DataBase) {
	me.DataBase = db
}

func New(lang Languages) *MEngine {
	var me MEngine
	switch lang {
	case Golang:
		me.Tpl = NewGoTpl()
	}
	return &me
}

func (me *MEngine) Generate(dbName, tabName string) error {

	// 连接数据库
	if err := me.DataBase.Connection(); err != nil {
		return err
	}

	// 通过数据库和表查询表的字段
	tabColumns, err := me.DataBase.Columns(dbName, tabName)

	if err != nil {
		return err
	}

	// 表字段转换成结构体字段
	structCulumns := me.Tpl.AssemblyColumns(tabColumns)

	return me.Tpl.Parse(tabName, structCulumns)
}
