package main

import (
	"testing"

	"github.com/mentat/graphc/pkg/common"
)

func TestTemplate(t *testing.T) {

	schema := &common.Schema{
		Types: map[string]common.Type{
			"Group": {
				Name: "Group",
				Fields: map[string]common.Field{
					"name": {
						Type: common.TypeDetail{
							Kind: common.TypeString,
						},
						Repeated: false,
						Nullable: false,
					},
					"attributes": {
						Type: common.TypeDetail{
							Kind: common.TypeString,
						},
						Repeated: true,
						Nullable: false,
						Args: []common.Argument{
							{
								Name: "limit",
								Type: common.TypeDetail{
									Kind: common.TypeInt,
								},
							},
						},
					},
				},
			},
			"User": {
				Name: "User",
				Fields: map[string]common.Field{
					"isActive": {
						Type: common.TypeDetail{
							Kind: common.TypeBoolean,
						},
						Repeated: false,
						Nullable: false,
					},
					"firstName": {
						Type: common.TypeDetail{
							Kind: common.TypeString,
						},
						Repeated: false,
						Nullable: false,
					},
					"groups": {
						Type: common.TypeDetail{
							Kind: common.TypeInt,
						},
						Repeated: true,
						Nullable: false,
					},
					"mainGroup": {
						Type: common.TypeDetail{
							Kind: common.TypeType,
							Name: "Group",
						},
					},
				},
			},
		},
		Queries: map[string]common.Field{
			"users": {
				Type: common.TypeDetail{
					Kind: common.TypeType,
					Name: "User",
				},
				Repeated: true,
				Nullable: false,
				Args: []common.Argument{
					{
						Name: "sort",
						Type: common.TypeDetail{
							Kind: common.TypeString,
						},
					},
					{
						Name:     "filter",
						Nullable: true,
						Type: common.TypeDetail{
							Kind: common.TypeString,
						},
					},
				},
			},
		},
	}

	ts := GraphCTS{}
	err := ts.GenerateTypes(schema, "myPackage", "testing.ts")
	if err != nil {
		t.Fatalf("Cannot generate types: %s", err)
	}

}
