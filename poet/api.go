// Package poet provides an API for generating Go constructs in a type-safe and
// reusable way.
package poet

// CodeBlock represent a block of code that can be included in a File
type CodeBlock interface {
	// String is the literal string serialization of the code.
	String() string
	// GetImports returns the imports required to use this code.
	GetImports() []Import
}

// statement represent a templated line of code.
type statement struct {
	Format       string        // Format specifies the format for the code
	Arguments    []interface{} // Arguments are used within the format string
	BeforeIndent int           // BeforeIndent augments the indent for the current statement.
	AfterIndent  int           // AfterIndent specifies indentation for subsequent statements.
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
	Tag string // Tag is a struct field tag, e.g. `json:"foo"`
}

// Import represent an individual imported package.
type Import interface {
	// GetPackage returns the go import package, like reflect.Type.PkgPath()
	GetPackage() string
	// GetAlias returns an alias string to refer to the import package, or the
	// empty string to omit an import alias.
	GetAlias() string
}

// TypeReference represent a specific reference (either an interface, function, struct or global)
type TypeReference interface {
	// GetImports returns the imports required to use this type. A struct, for example,
	// collects all the imports for its fields and itself.
	GetImports() []Import
	// GetName returns the go-syntax name of the type.
	GetName() string
}
