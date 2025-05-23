package common

import (
	"testing"
)

type MyPlugin struct {
}

func (my *MyPlugin) GetName() string {
	return ""
}
func (my *MyPlugin) GetAuthor() string {
	return ""
}
func (my *MyPlugin) GenerateTypes(schema *Schema, packageName string, target string) error {
	return nil
}

func TestTypes(t *testing.T) {

	schema := Schema{
		Types: map[string]Type{
			"blah": {Name: "blah", Fields: map[string]Field{
				"yo": {
					Type: TypeDetail{
						Kind: TypeBoolean,
					},
					Repeated: true,
					Nullable: false,
				},
			}},
		},
	}
	_ = schema
}
