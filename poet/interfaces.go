package poet

var _ CodeBlock = (*InterfaceSpec)(nil)
var _ TypeReference = (*InterfaceSpec)(nil)

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

// GetImports returns Imports used by the interface
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

// GetName returns the name and fulfills TypeReference.
func (i *InterfaceSpec) GetName() string {
	return i.Name
}

// String outputs the interface declaration
func (i *InterfaceSpec) String() string {
	writer := newCodeWriter()
	if i.Comment != "" {
		writer.WriteCodeBlock(Comment(i.Comment))
	}
	writer.WriteStatement(newStatement(0, 1, "type $L interface {", i.Name))

	for _, interf := range i.EmbeddedInterfaces {
		writer.WriteStatement(newStatement(0, 0, "$L", interf.GetName()))
	}

	for _, method := range i.Methods {
		if method.Comment != "" {
			writer.WriteStatement(newStatement(0, 0, "// $L", method.Comment))
		}
		signature, args := method.Signature()
		writer.WriteStatement(newStatement(0, 0, signature, args...))
	}

	writer.WriteStatement(newStatement(-1, 0, "}"))

	return writer.String()
}
