package gopoet

type CodeBlock interface {
	String() string
	GetImports() []Import
}

type Identifier struct {
	Name string
	Type TypeReference
}

type IdentifierParameter struct {
	Identifier
}

type IdentifierField struct {
	Identifier
	Tag string
}

type Import interface {
	GetPackage() string
	GetAlias() string
}

type TypeReference interface {
	GetImports() []Import
	GetName() string
}
