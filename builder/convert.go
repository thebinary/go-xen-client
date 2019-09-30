package main

import (
	"fmt"
	"log"
	"strings"
)

func isPrimitive(typedef string) bool {
	switch typedef {
	case "bool",
		"datetime",
		"float",
		"int",
		"string":
		return true
	}
	return false
}

func isMap(typedef string) bool {
	return strings.Contains(typedef, "->")
}

func isEnum(typedef string) bool {
	return strings.HasPrefix(typedef, "enum")
}

func isRef(typedef string) bool {
	return strings.HasSuffix(typedef, "ref")
}

func isRecord(typedef string) bool {
	return strings.HasSuffix(typedef, "record")
}

func isSet(typedef string) bool {
	return strings.HasSuffix(typedef, "set")
}

func convertPrimitive(typedef string) (gotype string) {
	gotype = "UNKWN"
	switch typedef {
	case "string", "int", "bool":
		gotype = typedef
	case "float":
		return "float32"
	case "datetime":
		return "time.Time"
	default:
		log.Printf("[UNKWN_PRIMITIVE_TYPE] Unknown type: %s", typedef)
		gotype = "interface{}"
	}
	return gotype
}

func ToCamelCase(org string) string {
	orgArr := strings.Split(org, "_")
	camelStr := ""
	for _, str := range orgArr {
		camelStr += fmt.Sprintf("%s%s", strings.ToUpper(string(str[0])), str[1:])
	}

	orgArr = strings.Split(camelStr, "-")
	camelStr = ""
	for _, str := range orgArr {
		camelStr += fmt.Sprintf("%s%s", strings.ToUpper(string(str[0])), str[1:])
	}
	return camelStr
}

func MapType(typedef string) string {
	//typedef_arr := strings.Split(typedef, " ")
	if isPrimitive(typedef) {
		return convertPrimitive(typedef)
	} else if isSet(typedef) {
		typedefSet := strings.Replace(typedef, " set", "", -1)
		setValueType := MapType(typedefSet)
		return fmt.Sprintf("[]%s", setValueType)
	} else if isEnum(typedef) {
		typedefEnum := strings.Replace(typedef, "enum ", "", -1)
		//return "Enum" + ToCamelCase(typedefEnum)
		return ToCamelCase(typedefEnum)
	} else if isRef(typedef) {
		//typedefRef := strings.Replace(typedef, " ref", "", -1)
		//return "Ref" + ToCamelCase(typedefRef)
		return "string"
	} else if isRecord(typedef) {
		typedefRecord := strings.Replace(typedef, "record", "", -1)
		return ToCamelCase(typedefRecord)
	} else if isMap(typedef) {
		mapTypedef := strings.Replace(typedef, "(", "", -1)
		mapTypedef = strings.Replace(mapTypedef, ")", "", -1)
		mapTypedef = strings.Replace(mapTypedef, " map", "", -1)
		map_types := strings.Split(mapTypedef, " -> ")
		mapKeyType := MapType(map_types[0])
		mapValueType := MapType(map_types[1])
		return fmt.Sprintf("map[%s]%s", mapKeyType, mapValueType)
	} else {
		return "unknown"
	}
}
