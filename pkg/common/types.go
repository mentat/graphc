package common

type GQLType int

const (
	TypeString GQLType = iota
	TypeInt
	TypeFloat
	TypeBoolean
	TypeID
	TypeScalar
	TypeType
	TypeEnum
	TypeUnion
	TypeInterface
	TypeInput
)

func (t GQLType) IsScalar() bool {
	switch t {
	case TypeType:
		fallthrough
	case TypeInterface:
		fallthrough
	case TypeUnion:
		fallthrough
	case TypeInput:
		fallthrough
	case TypeEnum:
		return false
	}
	return true
}

func (t GQLType) String() string {
	switch t {
	case TypeString:
		return "String"
	case TypeBoolean:
		return "Boolean"
	case TypeFloat:
		return "Float"
	case TypeInt:
		return "Int"
	case TypeID:
		return "ID"
	}
	return ""
}

type TypeDetail struct {
	Kind GQLType
	Name string
}

type Argument struct {
	Name         string
	Type         TypeDetail
	Repeated     bool
	Nullable     bool
	ValueDefault interface{}
}

type Field struct {
	Type     TypeDetail
	Repeated bool
	Nullable bool
	Args     []Argument
}

type Type struct {
	Name   string
	Fields map[string]Field
}

type Schema struct {
	Types         map[string]Type
	Queries       map[string]Field
	Mutations     map[string]Field
	Subscriptions map[string]Field
}
