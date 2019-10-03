package main

import (
	"fmt"
	"os"
	"text/template"
)

const template_obj_enums = `// This is a generated file. DO NOT EDIT manually.

package {{.Package}}
{{range $i, $e := .Enums}}
//{{CamelCase $e.Name}}: {{$e.Name}} enum type
type {{CamelCase $e.Name}} int
const({{range $j, $v := $e.Values}}
  {{if (eq $j 0)}} {{CamelCase $e.Name}}{{CamelCase $v.Name}} {{CamelCase $e.Name}} = iota //{{$v.Doc}} {{else}} {{CamelCase $e.Name}}{{CamelCase $v.Name}} //{{$v.Doc}} {{end}} {{end}}
)

func (e {{CamelCase $e.Name}}) String() string {
	switch e {
		{{range $j, $v := $e.Values}}
	case {{$j}}:
			return "{{$v.Name}}"
		{{end}}
		default:
			return ""
	}

}

func To{{CamelCase $e.Name}}(strValue string) {{CamelCase $e.Name}} {
	switch strValue { {{range $j, $v := $e.Values}}
		case "{{$v.Name}}":
			return {{$j}}{{end}}
		default:
			return -1
	}
}
{{end}}
`

func genObjectEnums(packageName string, objectName string, enums []ObjectEnum) (err error) {
	tmpl, err := template.New("struct").Funcs(
		map[string]interface{}{
			"CamelCase": ToCamelCase,
		},
	).Parse(template_obj_enums)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("../%s.enums.gen.go", ToCamelCase(objectName)))
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, map[string]interface{}{
		"Package": packageName,
		"Enums":   enums,
	})
	return err
}
