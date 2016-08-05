package gopoet

type CodeBlock interface {
	String() string
	Packages() []ImportSpec
}

type Identifier struct {
	Name string
	Type ImportSpec
}

type IdentifierParameter struct {
	Identifier
}

type IdentifierField struct {
	Identifier
	Tag string
}

type ImportSpec interface {
	GetPackage() string
	GetPackageAlias() string
	NeedsQualifier() bool
	GetName() string
}
