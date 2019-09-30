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
//go:generate stringer -type {{CamelCase $e.Name}}
type {{CamelCase $e.Name}} int
const({{range $j, $v := $e.Values}}
  {{if (eq $j 0)}} {{CamelCase $e.Name}}{{CamelCase $v.Name}} {{CamelCase $e.Name}} = iota //{{$v.Doc}} {{else}} {{CamelCase $e.Name}}{{CamelCase $v.Name}} //{{$v.Doc}} {{end}} {{end}}
)
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
