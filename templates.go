package main

import (
	"text/template"
)

var header string = `
package {{.PkgName}}

import (
  "lms/lib/mysql"
)


`

var modelStruct string = `
var (
	{{.TbName}}s       map[{{.MapKey}}]*{{.TbName}}    
)
type {{.TbName}} struct {
	{{range .Fields}}{{.Name}} {{.Type}} {{if .Comment}} // {{.Comment}}{{end}}
	{{end}}
}
`

var objApi string = ``

var queryApi string = `
func Load(db *mysql.Connection) {
	load{{.TbName}}(db)
}

func load{{.TbName}}(db *mysql.Connection) {

	res, err := db.ExecuteFetch([]byte("select * from {{.TableName}};"), -1)
	if err != nil {
		panic(err)
	}
	length := len(res.Rows)
	{{.TbName}}s = make(map[{{.MapKey}}]*{{.TbName}}, length)
	if length > 0 {
    {{range .Fields}}i{{.Name}}:=res.Map("{{.ColumnName}}")
    {{end}}for _, row := range res.Rows {
			{{.TbName}}s[row.{{.MapKey}}(iId)] = &{{.TbName}}{
				{{range .Fields}}{{.Name}}: row.{{.ConverterFuncName}}(i{{.Name}}),
			    {{end}} 
			}
		}
	}
}
`

var managedApi string = ``

var (
	tmHeader        *template.Template
	tmStruct        *template.Template
	tmObjApi        *template.Template
	tmQueryApi      *template.Template
	tmManagedObjApi *template.Template
)

func init() {
	tmHeader = template.Must(template.New("header").Parse(header))
	tmStruct = template.Must(template.New("modelStruct").Parse(modelStruct))
	tmObjApi = template.Must(template.New("objApi").Parse(objApi))
	tmQueryApi = template.Must(template.New("queryApi").Parse(queryApi))
	tmManagedObjApi = template.Must(template.New("managedApi").Parse(managedApi))
}
