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
	"fmt"
	"html/template"
	"os"
	"strings"
)

type GoCodeTemplate struct {
	*DataType
	StructTpl  string
	StructName string
	Columuns   []*StructColumun
}

type StructColumun struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

func NewGoTpl() *GoCodeTemplate {
	return &GoCodeTemplate{
		StructTpl: `type {{.TableName | ToCamelCase}} struct {
			{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
				{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
			{{end}}}
			func (model {{.TableName | ToCamelCase}}) TableName() string {
				return "{{.TableName}}"
			}`,
		DataType: NewDataType(Golang),
	}
}

func (tpl *GoCodeTemplate) Parse(tableName string, tplColumns []*StructColumun) error {
	tmpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": func(s string) string {
			s = strings.Replace(s, "_", " ", -1)
			s = strings.Title(s)
			return strings.Replace(s, " ", "", -1)
		},
	}).Parse(tpl.StructTpl))

	err := tmpl.Execute(os.Stdout, tpl)
	if err != nil {
		return err
	}
	return nil
}

func (tpl *GoCodeTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumun {
	tplColumns := make([]*StructColumun, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumun{
			Name:    column.ColumnName,
			Type:    tpl.DataType.Table[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}
