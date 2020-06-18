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
  {{range .Object.Fields}}{{CamelCase .Name}} {{TypeName .Type}}               // {{.Description}}
  {{end}}
}

func From{{CamelCase .Object.Name}}ToXml({{.Object.Name}} *{{CamelCase .Object.Name}}) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)
	{{range .Object.Fields}}
	{{if (IsMap .Type)}}
		{{.Name}} := make(xmlrpc.Struct)
		for key, value := range {{$o.Name}}.{{CamelCase .Name}} {
			{{.Name}}[{{if (eq (MapKeyType .Type) "int")}}strconv.Itoa(key){{else}}key{{if (IsEnum (MapKeyType .Type))}}.String(){{end}}{{end}}] = value{{if (IsEnum .Type)}}.String(){{end}}
		}
		result["{{.Name}}"] = {{.Name}}
	{{else}}
		result["{{.Name}}"] = {{if (eq (TypeName .Type) "int")}}strconv.Itoa({{$o.Name}}.{{CamelCase .Name}}){{else}}{{$o.Name}}.{{CamelCase .Name}}{{if (IsSet .Type)}}{{else}}{{if (IsEnum .Type)}}.String(){{end}}{{end}}{{end}}
	{{end}}
	{{end}}
	return result
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
				{{ObjectFieldConversion .}}{{end}}
		}
	}

	return resultObj{{end}}
}

{{range .Object.Messages}}
/* {{CamelCase .Name}}: {{.Description}} */{{$resultType := index .Result 0}}
func (client *XenClient) {{CamelCase $o.Name}}{{CamelCase .Name}}({{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{.Name}}{{end}} {{TypeName .Type}},{{end}}{{end}}) ({{if (eq $resultType "void")}}{{else}}result {{TypeName $resultType}}, {{end}}err error){
	{{if (eq $resultType "void")}}_, err ={{else}}obj, err :={{end}} client.APICall("{{$o.Name}}.{{.Name}}", {{range .Params}}{{if (eq .Name "session_id")}}{{else}}{{if (eq .Name "type")}}xtype{{else}}{{if IsRecord .Type}}From{{TypeName .Type}}ToXml(&{{.Name}}){{else}}{{if (IsEnum .Type)}}{{.Name}}.String(){{else}}{{.Name}}{{end}}{{end}}{{end}},{{end}}{{end}})
	if err != nil {
		return
	}
	{{GenResult $resultType}}
	return
}
{{end}}
`

func genObjectFieldConversion(field ObjectField) (returnCode string) {
	returnCode = ""
	typedef := string(field.Type)
	fieldType := MapType(typedef)
	camelFieldName := ToCamelCase(field.Name)
	if isSet(typedef) {
		setElementType := strings.Replace(typedef, " set", "", -1)
		returnCode += fmt.Sprintf(`if interim, ok := keyValue.([]interface{}); ok {
			resultObj.%s = make(%s, len(interim))
			for i, interimValue := range interim {
				if v, ok := interimValue.(%s); ok {
					resultObj.%s[i] = v
				}
			}
		}`, camelFieldName, fieldType, MapType(setElementType), camelFieldName)
	} else if isMap(typedef) {
		mapTypedef := strings.Replace(typedef, "(", "", -1)
		mapTypedef = strings.Replace(mapTypedef, ")", "", -1)
		mapTypedef = strings.Replace(mapTypedef, " map", "", -1)
		map_types := strings.Split(mapTypedef, " -> ")
		mapKeyType := MapType(strings.TrimSpace(map_types[0]))
		mapValueType := MapType(strings.TrimSpace(map_types[1]))

		if mapKeyType == "string" {
			mappingNullValue := `""`
			if isPrimitive(map_types[1]) && mapValueType != "string" {
				mappingNullValue = "0"
			}
			mappingcode := fmt.Sprintf(`if v, ok := mapKeyValue.(%s); ok {
					resultObj.%s[mapKeyName] = v
				} else {
					resultObj.%s[mapKeyName] = %s
				}`, mapValueType, camelFieldName, camelFieldName, mappingNullValue)
			if isEnum(map_types[1]) {
				mappingcode = fmt.Sprintf(`if v, ok := mapKeyValue.(string); ok {
						resultObj.%s[mapKeyName] = To%s(v)
					} else {
						resultObj.%s[mapKeyName] = 0
					}`, camelFieldName, mapValueType, camelFieldName)
			}

			returnCode += fmt.Sprintf(`
				resultObj.%s = map[string]%s{}
				interimMap := reflect.ValueOf(keyValue).MapKeys()
				for _, mapKey := range interimMap {
					mapKeyName := mapKey.String()
					mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
					%s
				}`, camelFieldName, mapValueType, mappingcode)
		}
	} else if isPrimitive(typedef) || isEnum(typedef) {
		returnCode += fmt.Sprintf(`if v, ok := keyValue.(%s); ok {
			resultObj.%s = v
		}`, fieldType, camelFieldName)
	} else {
		returnCode += fmt.Sprintf(`if v, ok := keyValue.(%s); ok {
			resultObj.%s = v
		}`, fieldType, camelFieldName)
	}

	return returnCode
}

func genReturn(resultType string) (returnCode string) {
	returnCode = "//result conversion not implemented yet"
	returnCode += "\n" + `log.Printf("%+v", obj)`

	if resultType == "void" {
		returnCode = "// no return result"
	} else if MapType(resultType) == "interface{}" {
		returnCode = "result = obj"
	} else if isPrimitive(MapType(resultType)) || isPrimitive(resultType) {
		returnCode = fmt.Sprintf("result = obj.(%s)", MapType(resultType))
	} else if isSet(resultType) {
		setElementType := MapType(strings.TrimSpace(strings.Replace(resultType, " set", "", -1)))
		conversionCode := fmt.Sprintf(`value.(%s)`, setElementType)
		if !isPrimitive(setElementType) {
			if isEnum(resultType) {
				conversionCode = fmt.Sprintf("To%s(value.(string))", setElementType)
			} else {
				conversionCode = fmt.Sprintf("*To%s(value)", setElementType)
			}
		}
		returnCode = fmt.Sprintf(`
			result = make(%s, len(obj.([]interface{})))
			for i, value := range obj.([]interface{}) {
				result[i] = %s
			}
		`, MapType(resultType), conversionCode)
	} else if isMap(resultType) {
		mapTypedef := strings.Replace(resultType, "(", "", -1)
		mapTypedef = strings.Replace(mapTypedef, ")", "", -1)
		mapTypedef = strings.Replace(mapTypedef, " map", "", -1)
		map_types := strings.Split(mapTypedef, " -> ")
		mapKeyType := MapType(strings.TrimSpace(map_types[0]))
		mapValueType := MapType(strings.TrimSpace(map_types[1]))

		if mapKeyType == "string" || mapKeyType == "int" {
			mapKeyReflect := "key.String()"
			if mapKeyType == "int" {
				mapKeyReflect = "int(key.Int())"
			}
			if mapValueType == "string" || mapValueType == "int" {
				mapValueReflect := "obj.String()"
				if mapValueType == "int" {
					mapValueReflect = "int(obj.Int())"
				} else if mapValueType == "float32" {
					mapValueReflect = "float32(obj.Float())"
				}
				returnCode = fmt.Sprintf(`
					interim := reflect.ValueOf(obj)
					result = map[%s]%s{}
					for _, key := range interim.MapKeys() {
						obj := interim.MapIndex(key)
						result[%s] = %s
					}
				`, mapKeyType, mapValueType, mapKeyReflect, mapValueReflect)
			} else if isSet(map_types[1]) {
				setElementType := MapType(strings.TrimSpace(strings.Replace(map_types[1], " set", "", -1)))
				if isEnum(setElementType) {
					returnCode = fmt.Sprintf(`
						interim := reflect.ValueOf(obj)
						result = map[%s]%s{}
						for _, key := range interim.MapKeys() {
							obj := interim.MapIndex(key)
							mapObj := To%s(obj.String())
							result[%s] = mapObj
						}
					`, mapKeyType, mapValueType, mapValueType, mapKeyReflect)
				} else if setElementType == "string" {
					returnCode = fmt.Sprintf(`
						interim := reflect.ValueOf(obj)
						result = map[%s][]string{}
						for _, key := range interim.MapKeys() {
							obj := interim.MapIndex(key)
							interimObj := obj.Interface().([]interface{})
							interimResult := make([]string, len(interimObj))
							for i, interimValue := range interimObj {
								interimResult[i] = interimValue.(string)
							}
							result[%s] = interimResult
						}
					`, mapKeyType, mapKeyReflect)
				}
			} else if isEnum(map_types[1]) {
				returnCode = fmt.Sprintf(`
						interim := reflect.ValueOf(obj)
						result = map[%s]%s{}
						for _, key := range interim.MapKeys() {
							obj := interim.MapIndex(key)
							mapObj := To%s(obj.String())
							result[%s] = mapObj
						}
					`, mapKeyType, mapValueType, mapValueType, mapKeyReflect)
			} else if !isPrimitive(map_types[1]) {
				returnCode = fmt.Sprintf(`
						interim := reflect.ValueOf(obj)
						result = map[%s]%s{}
						for _, key := range interim.MapKeys() {
							obj := interim.MapIndex(key)
							mapObj := To%s(obj.Interface())
							result[%s] = *mapObj
						}
					`, mapKeyType, mapValueType, mapValueType, mapKeyReflect)
			}
		}
	} else if isEnum(resultType) {
		returnCode = fmt.Sprintf(`
			result = To%s(obj.(string))
		`, MapType(resultType))
	} else {
		returnCode = fmt.Sprintf(`
			result = *To%s(obj.(interface{}))
		`, MapType(resultType))
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
			"CamelCase":             ToCamelCase,
			"TypeName":              MapType,
			"GenResult":             genReturn,
			"ObjectFieldConversion": genObjectFieldConversion,
			"IsEnum":                isEnum,
			"IsRecord":              isRecord,
			"IsSet":                 isSet,
			"IsMap":                 isMap,
			"IsPrimitive":           isPrimitive,
			"MapKeyType":            MapKeyType,
			"MapValueType":          MapValueType,
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
