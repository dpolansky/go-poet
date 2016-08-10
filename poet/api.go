package poet

// CodeBlock represent a block of code that can be included in a File
type CodeBlock interface {
	String() string
	GetImports() []Import
}

// statement represent a templated line of code.
type statement struct {
	Format       string        // Format specifies the format for the code
	Arguments    []interface{} // Arguments are used within the format string
	BeforeIndent int           // BeforeIndent augments the indent for the current statement.
	AfterIndent  int           // AfterIndent specifies ondentation for subsequent statements.
}

// Identifier represent an instance of a variable
type Identifier struct {
	Name string        // Name of the instance of a variable (e.g. in "var a int", a)
	Type TypeReference // Type of the variable
}

// IdentifierParameter represent a parameter in a function/method
type IdentifierParameter struct {
	Identifier
	Variadic bool // Variadic specifies whether the parameter is a variadic
}

// IdentifierField represent a field in a struct
type IdentifierField struct {
	Identifier
	Tag string // Tag is a
}

// Import represent an indivdual import
type Import interface {
	GetPackage() string
	GetAlias() string
}

// TypeReference represent a specific reference (either an interface, struct or a global)
type TypeReference interface {
	GetImports() []Import
	GetName() string
}
