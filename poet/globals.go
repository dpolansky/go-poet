package poet

import (
	"bytes"
	"fmt"
)

// VariableGrouping represents a collection of variables and/or constants that will
// be separated into groups on output.
type VariableGrouping struct {
	Variables []*Variable
}

var _ CodeBlock = (*VariableGrouping)(nil)

// Variable adds a new variable to this variable grouping.
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

// Constant adds a new constant to this variable grouping.
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

// GetImports returns a slice of imports that this variable grouping uses.
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
		buf.WriteString("\t" + v.GetDeclaration())
	}
	buf.WriteString(")\n")
	return buf.String()
}

// Variable represents a variable, with name, type, and value.
type Variable struct {
	Identifier
	Constant bool
	Format   string
	Args     []interface{}
}

var _ CodeBlock = (*Variable)(nil)

// GetImports returns a slice of imports that this variable uses.
func (v *Variable) GetImports() []Import {
	return v.Type.GetImports()
}

// GetDeclaration returns the name and type of this variable, for example: 'foo string'.
func (v *Variable) GetDeclaration() string {
	w := newCodeWriter()
	args := append([]interface{}{v.Name, v.Type}, v.Args...)
	w.WriteStatement(newStatement(0, 0, "$L $T = "+v.Format, args...))
	return w.String()
}

func (v *Variable) String() string {
	var prefix string
	if v.Constant {
		prefix = "const"
	} else {
		prefix = "var"
	}
	return fmt.Sprintf("%s %s", prefix, v.GetDeclaration())
}
