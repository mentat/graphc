package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flosch/pongo2"
	"github.com/mentat/graphc/pkg/common"

	_ "embed"
)

// map of GraphQL types to Typescript types
var typeNameMap = map[common.GQLType]string{
	common.TypeString:  "string",
	common.TypeInt:     "number",
	common.TypeFloat:   "number",
	common.TypeBoolean: "boolean",
}

//go:embed templates/types.jinja2
var typesTmplStr string

func GetTSType(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	inter := in.Interface()
	real, ok := inter.(common.TypeDetail)
	if !ok {
		fmt.Println("Error converting interface to TypeDetail")
		return pongo2.AsValue(""), nil
	}
	typeName := typeNameMap[common.GQLType(real.Kind)]
	if typeName == "" {
		return pongo2.AsValue(real.Name), nil
	}
	return pongo2.AsValue(typeName), nil
}

type GraphCTS struct {
}

func (my *GraphCTS) GetName() string {
	return ""
}
func (my *GraphCTS) GetAuthor() string {
	return ""
}
func (my *GraphCTS) GenerateTypes(schema *common.Schema, packageName string, target string) error {

	var typesTmpl = pongo2.Must(pongo2.FromString(typesTmplStr))

	file, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}

	err = typesTmpl.ExecuteWriter(pongo2.Context{
		"schema": schema,
		"cap":    common.CapitalizeFirstLetter,
		"getType": func(name string) *pongo2.Value {
			fmt.Printf("Get type: %s\n", name)
			return pongo2.AsValue(schema.Types[name])
		},
	}, file)

	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := pongo2.RegisterFilter("ts", GetTSType)
	if err != nil {
		fmt.Printf("Error registering filter: %s\n", err)
	}
}

func main() {
}
