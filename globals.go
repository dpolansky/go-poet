package gopoet

import "bytes"

type VariableGrouping struct {
	Variables []Variable
	Constant  bool
}

var _ CodeBlock = (*VariableGrouping)(nil)

func (g *VariableGrouping) GetImports() []Import {
	imports := []Import{}
	for _, vari := range g.Variables {
		imports = append(imports, vari.GetImports()...)
	}
	return imports
}

func (g *VariableGrouping) String() string {
	buff := bytes.Buffer{}
	if g.Constant {
		buff.WriteString("const (\n")
	} else {
		buff.WriteString("var (\n")
	}

	for _, vari := range g.Variables {
		buff.WriteString("\t")
		buff.WriteString(vari.GetDeclaration())
		buff.WriteString("\n")
	}

	buff.WriteString(")\n")

	return buff.String()
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
