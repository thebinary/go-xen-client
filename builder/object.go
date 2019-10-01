package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const template_obj = `// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w {{CamelCase .Object.Name}}.gen.go

package {{.Package}}
{{$o := .Object}}
//{{CamelCase .Object.Name}}: {{.Object.Description}}
type {{CamelCase .Object.Name}} struct {
  {{range .Object.Fields}}{{CamelCase .Name}} {{.Type}}               // {{.Description}}
  {{end}}
}

func To{{CamelCase .Object.Name}}(obj interface{}) (resultObj *{{CamelCase .Object.Name}}) {
	{{if not .Object.Fields}}return &{{CamelCase .Object.Name}}{}{{else}}
	objValue := reflect.ValueOf(obj)
	resultObj = &{{CamelCase .Object.Name}}{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName { {{range .Object.Fields}}
			case "{{.Name}}":
				if v, ok := keyValue.({{.Type}}); ok {
					resultObj.{{CamelCase .Name}} = v
				}{{end}}
		}
	}

	return resultObj{{end}}
}

{{range .Object.Messages}}
/* {{CamelCase .Name}}: {{.Description}} */{{$resultType := index .Result 0}}
func (client *XenClient) {{CamelCase $o.Name}}{{CamelCase .Name}}({{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{.Name}}{{end}} {{TypeName .Type}},{{end}}{{end}}) ({{if (eq $resultType "void")}}{{else}}result {{TypeName $resultType}}, {{end}}err error){
	{{if (eq $resultType "void")}}_, err ={{else}}obj, err :={{end}} client.APICall("{{$o.Name}}.{{.Name}}", {{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{.Name}}{{end}},{{end}}{{end}})
	if err != nil {
		return
	}
	{{GenResult $resultType}}
	return
}
{{end}}
`

func genReturn(resultType string) (returnCode string) {
	if resultType == "void" {
		returnCode = "// no return result"
	} else if isPrimitive(resultType) {
		returnCode = fmt.Sprintf("result = obj.(%s)", MapType(resultType))
	} else if isSet(resultType) {
		returnCode = fmt.Sprintf(`
			result = make(%s, len(obj.([]interface{})))
			for i, value := range obj.([]interface{}) {
				result[i] = value.(%s)
			}
		`, MapType(resultType), strings.Replace(MapType(resultType), "[]", "", -1))
	} else {
		returnCode = "//not implemented yet"
		returnCode += "\n" + `log.Printf("%+v", obj)`
	}
	return returnCode
}

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
			"GenResult": genReturn,
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
