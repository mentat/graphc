package common

type Plugin interface {
	GetName() string
	GetAuthor() string
	GenerateTypes(schema *Schema, packageName string, target string) error
}
