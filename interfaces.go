package gopoet

type InterfaceSpec struct {
	// CodeBlock

	Name               string
	Comment            string
	EmbeddedInterfaces []ImportSpec
	Methods            []FuncSpec
}

func NewInterfaceSpec(name string) *InterfaceSpec {
	return &InterfaceSpec{
		Name: name,
	}
}

func (i *InterfaceSpec) String() string {
	writer := NewCodeWriter()
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

func (i *InterfaceSpec) Packages() []ImportSpec {
	packages := []ImportSpec{}

	for _, method := range i.Methods {
		packages = append(packages, method.Packages()...)
	}

	packages = append(packages, i.EmbeddedInterfaces...)

	return packages
}

func (i *InterfaceSpec) Method(spec FuncSpec) *InterfaceSpec {
	i.Methods = append(i.Methods, spec)
	return i
}

func (i *InterfaceSpec) EmbedInterface(spec ImportSpec) *InterfaceSpec {
	i.EmbeddedInterfaces = append(i.EmbeddedInterfaces, spec)
	return i
}
