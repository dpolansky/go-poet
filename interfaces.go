package gopoet

// InterfaceSpec represents an interface
type InterfaceSpec struct {
	// CodeBlock

	Name               string
	Comment            string
	EmbeddedInterfaces []TypeReference
	Methods            []*FuncSpec
}

// NewInterfaceSpec constructs a new interface with the given name
func NewInterfaceSpec(name string) *InterfaceSpec {
	return &InterfaceSpec{
		Name: name,
	}
}

// Method adds a new method to the interface
func (i *InterfaceSpec) Method(spec *FuncSpec) *InterfaceSpec {
	i.Methods = append(i.Methods, spec)
	return i
}

// EmbedInterface specifies an interface to embed in the interface
func (i *InterfaceSpec) EmbedInterface(interfaceType TypeReference) *InterfaceSpec {
	i.EmbeddedInterfaces = append(i.EmbeddedInterfaces, interfaceType)
	return i
}

var _ CodeBlock = (*InterfaceSpec)(nil)

// GetImports returns Import's used by the interface
func (i *InterfaceSpec) GetImports() []Import {
	packages := []Import{}

	for _, method := range i.Methods {
		packages = append(packages, method.GetImports()...)
	}

	for _, embedded := range i.EmbeddedInterfaces {
		packages = append(packages, embedded.GetImports()...)
	}

	return packages
}

// String outputs the interface declaration
func (i *InterfaceSpec) String() string {
	writer := NewCodeWriter()
	if i.Comment != "" {
		writer.WriteStatement(Statement{
			Format:    "// $L",
			Arguments: []interface{}{i.Comment},
		})
	}
	writer.WriteStatement(Statement{
		Format:      "type $L interface {",
		Arguments:   []interface{}{i.Name},
		AfterIndent: 1,
	})

	for _, interf := range i.EmbeddedInterfaces {
		writer.WriteStatement(Statement{
			Format:    "$L",
			Arguments: []interface{}{interf.GetName()},
		})
	}

	for _, method := range i.Methods {
		if method.Comment != "" {
			writer.WriteStatement(Statement{
				Format:    "// $L",
				Arguments: []interface{}{method.Comment},
			})
		}
		signature, args := method.Signature()
		writer.WriteStatement(Statement{
			Format:    signature,
			Arguments: args,
		})
	}

	writer.WriteStatement(Statement{
		Format:       "}",
		BeforeIndent: -1,
	})

	return writer.String()
}
