package gopoet

import "bytes"

type VariableGrouping struct {
	Variables []*Variable
}

var _ CodeBlock = (*VariableGrouping)(nil)

func (g *VariableGrouping) Variable(name string, typ TypeReference, format string, args ...interface{}) *VariableGrouping {
	v := &Variable{
		Identifier: Identifier{
			Name: name,
			Type: typ,
		},
		Constant: false,
		Format:   format,
		Args:     args,
	}
	g.Variables = append(g.Variables, v)
	return g
}

func (g *VariableGrouping) Constant(name string, typ TypeReference, format string, args ...interface{}) *VariableGrouping {
	v := &Variable{
		Identifier: Identifier{
			Name: name,
			Type: typ,
		},
		Constant: true,
		Format:   format,
		Args:     args,
	}
	g.Variables = append(g.Variables, v)
	return g
}

func (g *VariableGrouping) GetImports() []Import {
	imports := []Import{}
	for _, vari := range g.Variables {
		imports = append(imports, vari.GetImports()...)
	}
	return imports
}

func (g *VariableGrouping) String() string {
	buff := bytes.Buffer{}

	constants := []*Variable{}
	vars := []*Variable{}

	for _, v := range g.Variables {
		if v.Constant {
			constants = append(constants, v)
		} else {
			vars = append(vars, v)
		}
	}

	if len(constants) > 0 {
		buff.WriteString(writeGroupAsString("const", constants))
	}

	// if both groups are populated, add a newline between them
	if len(constants) > 0 && len(vars) > 0 {
		buff.WriteString("\n")
	}

	if len(vars) > 0 {
		buff.WriteString(writeGroupAsString("var", vars))
	}

	return buff.String()
}

func writeGroupAsString(groupName string, vars []*Variable) string {
	buf := bytes.Buffer{}

	buf.WriteString(groupName + " (\n")
	for _, v := range vars {
		buf.WriteString("\t" + v.GetDeclaration() + "\n")
	}
	buf.WriteString(")\n")
	return buf.String()
}

type Variable struct {
	Identifier
	Constant bool
	Format   string
	Args     []interface{}
}

var _ CodeBlock = (*Variable)(nil)

func (v *Variable) GetImports() []Import {
	return v.Type.GetImports()
}

func (v *Variable) GetDeclaration() string {
	buff := bytes.Buffer{}
	buff.WriteString(Template("$L $T", v.Name, v.Type))
	buff.WriteString(" = ")
	buff.WriteString(Template(v.Format, v.Args...))
	return buff.String()
}

func (v *Variable) String() string {
	buff := bytes.Buffer{}
	if v.Constant {
		buff.WriteString("const ")
	} else {
		buff.WriteString("var ")
	}
	buff.WriteString(v.GetDeclaration())
	buff.WriteString("\n")

	return buff.String()
}
