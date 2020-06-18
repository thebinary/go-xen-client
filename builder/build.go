package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type EntityLifecycle struct {
	Transition  string `json:"transition"`
	Release     string `json:"release"`
	Description string `json:"description"`
}

type ObjectField struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Qualifier   string            `json:"qualifier"`
	Tag         string            `json:"tag"`
	Lifecycle   []EntityLifecycle `json:"lifecycle"`
}

type ObjectMessage struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Result      []string             `json:"result"`
	Params      []ObjectMessageParam `json:"params"`
	Errors      []interface{}        `json:"errors"`
	Roles       []string             `json:"string"`
	Tag         string               `json:"tag"`
	Lifecycle   []EntityLifecycle    `json:"lifecycle"`
	Implicit    bool                 `json:"implicit"`
}
type ObjectMessageParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}

type ObjectEnumValue struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
}
type ObjectEnum struct {
	Name   string            `json:"name"`
	Values []ObjectEnumValue `json:"values"`
}
type ObjectDef struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Fields      []ObjectField   `json:"fields"`
	Messages    []ObjectMessage `json:"messages"`
	Enums       []ObjectEnum    `json:"enums"`
}

func (o ObjectDef) Strings() string {
	printer := fmt.Sprintf(`
Object: %s
Description: %s

Fields:
%+v

Messages:
%+v
`, o.Name, o.Description, o.Fields, o.Messages,
	)
	return printer
}

func (field ObjectField) Strings() string {
	return fmt.Sprintf("%s | %s | %v\n", field.Name, field.Description, field.Type)
}

func (message ObjectMessage) Strings() string {
	return fmt.Sprintf("%s | %s | %v | %v\n", message.Name, message.Description, message.Params, message.Result)
}

func processXenAPI(filename string) (objMap map[string]ObjectDef) {
	xenapi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer xenapi.Close()

	objDefs := []ObjectDef{}
	err = json.NewDecoder(xenapi).Decode(&objDefs)
	if err != nil {
		panic(err)
	}

	objMap = map[string]ObjectDef{}
	for _, obj := range objDefs {
		objMap[obj.Name] = obj
	}

	return objMap
}

func main() {
	objMap := processXenAPI("xenapi.json")

	/*
		err := genObject("go_xen_client", objMap["host"])
		if err != nil {
			log.Fatal(err)
		}

		err = genObjectEnums("go_xen_client", "host", objMap["host"].Enums)
		if err != nil {
			log.Fatal(err)
		}

		err = genObject("go_xen_client", objMap["data_source"])
		if err != nil {
			log.Fatal(err)
		}
	*/

	for k := range objMap {
		log.Printf("Generating object definition for: %s", k)
		err := genObject("go_xen_client", objMap[k])
		if err != nil {
			panic(err)
		}

		// generate only if enums exist in the object
		// VM_metrics has only one enum that is already included in VM
		if enums := objMap[k].Enums; len(enums) > 0 && objMap[k].Name != "VM_metrics" {
			log.Printf("Generating enums definition for: %s", k)
			err = genObjectEnums("go_xen_client", k, enums)
			if err != nil {
				panic(err)
			}
		} else {
			log.Printf("Ignoring enums definition for: %s, enums_count=0", k)
		}
	}

}
