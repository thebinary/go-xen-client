package main

import (
	"fmt"
	"os"
	"text/template"
)

const template_obj = `// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w {{CamelCase .Object.Name}}.gen.go

package {{.Package}}
{{$o := .Object}}
//{{CamelCase .Object.Name}}Record: {{.Object.Description}}
type {{CamelCase .Object.Name}}Record struct {
  {{range .Object.Fields}}{{CamelCase .Name}} {{.Type}}               // {{.Description}}
  {{end}}
}

{{range .Object.Messages}}
/* {{CamelCase .Name}}: {{.Description}} */
func (client *XenClient) {{CamelCase $o.Name}}{{CamelCase .Name}}({{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{.Name}}{{end}} {{TypeName .Type}},{{end}}{{end}}) (result xmlrpc.Struct, err error){
	log.Println("called {{$o.Name}}.{{.Name}}")
	return
  //return client.Call("{{$o.Name}}.{{.Name}}", {{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{.Name}}{{end}},{{end}}{{end}})
}
{{end}}
`

func genObject(packageName string, objDef ObjectDef) (err error) {
	// fix parameter name "interface"
	for mi, m := range objDef.Messages {
		for fi, f := range m.Params {
			if f.Name == "interface" {
				objDef.Messages[mi].Params[fi].Name = "iface"
			}
		}
	}

	tmpl, err := template.New("object").Funcs(
		map[string]interface{}{
			"CamelCase": ToCamelCase,
			"TypeName":  MapType,
		},
	).Parse(template_obj)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("../%s.gen.go", ToCamelCase(objDef.Name)))
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, map[string]interface{}{
		"Package": packageName,
		"Object":  objDef,
	})
	return err
}
