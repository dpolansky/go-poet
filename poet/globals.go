package poet

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
		InGroup:  true,
		Value:    newStatement(0, 0, format, args...),
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
		InGroup:  true,
		Value:    newStatement(0, 0, format, args...),
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
	w := newCodeWriter()
	for _, s := range g.GetStatements() {
		w.WriteStatement(s)
	}
	return w.String()
}

func (g *VariableGrouping) GetStatements() []Statement {
	var constants []*Variable
	var vars []*Variable
	var statements []Statement

	for _, v := range g.Variables {
		if v.Constant {
			constants = append(constants, v)
		} else {
			vars = append(vars, v)
		}
	}

	statements = append(statements, globalsAsStatements("const", constants)...)
	// if both groups are populated, add a newline between them
	if len(constants) > 0 && len(vars) > 0 {
		statements = append(statements, Statement{})
	}
	statements = append(statements, globalsAsStatements("var", vars)...)

	return statements
}

func globalsAsStatements(groupName string, vars []*Variable) []Statement {
	if len(vars) == 0 {
		return nil
	}
	var s []Statement
	s = append(s, newStatement(0, 1, "$L (", groupName))
	for _, v := range vars {
		s = append(s, v.GetStatements()...)
	}
	s = append(s, newStatement(-1, 0, ")"))
	return s
}

// Variable represents a variable, with name, type, and value.
type Variable struct {
	Identifier
	Comment  string
	Value    Statement
	Constant bool
	InGroup  bool
}

var _ CodeBlock = (*Variable)(nil)

// GetImports returns a slice of imports that this variable and its value uses.
func (v *Variable) GetImports() []Import {
	return v.Type.GetImports()
}

// GetStatements returns Value.GetStatements() with the first
// statement prepended with the variable declaration.
func (v *Variable) GetStatements() []Statement {
	var s []Statement
	s = append(s, Comment(v.Comment).GetStatements()...)
	s = append(s, v.statement())
	return s
}

func (v *Variable) String() string {
	w := newCodeWriter()
	for _, s := range v.GetStatements() {
		w.WriteStatement(s)
	}
	return w.String()
}

func (v *Variable) statement() Statement {
	if v.Value.Format == "" {
		return newStatement(0, 0, "$L$L $T", v.prefix(), v.Name, v.Type)
	}
	assignment := newStatement(0, 0, "$L$L $T = ", v.prefix(), v.Name, v.Type)

	return appendStatements(assignment, v.Value)
}

// prefix returns (var |const ). Note trailing space!
func (v *Variable) prefix() string {
	if v.InGroup {
		return ""
	}
	if v.Constant {
		return "const "
	}
	return "var "
}
