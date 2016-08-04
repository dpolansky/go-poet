package gopoet

import (
	"bytes"
	"reflect"
)

// FuncSpec represents information needed to write a function
type FuncSpec struct {
	// CodeBlock

	Name             string
	Comment          string
	Parameters       []IdentifierParameter
	ResultParameters []IdentifierParameter
	Statements       []Statement
}

func NewFuncSpec(name string) *FuncSpec {
	return &FuncSpec{
		Name:             name,
		Parameters:       []IdentifierParameter{},
		ResultParameters: []IdentifierParameter{},
		Statements:       []Statement{},
	}
}

func (f *FuncSpec) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("func " + f.Name + "() {}")

	return buffer.String()
}

func (f *FuncSpec) Packages() []ImportSpec {
	packages := []ImportSpec{}

	for _, st := range f.Statements {
		for _, arg := range st.Arguments {
			if reflect.TypeOf(arg) == reflect.TypeOf((*ImportSpec)(nil)) {
				packages = append(packages, arg.(ImportSpec))
			}
		}
	}

	for _, param := range f.Parameters {
		packages = append(packages, param.Type)
	}

	for _, param := range f.ResultParameters {
		packages = append(packages, param.Type)
	}

	return packages
}

// Statement is a convenient method to append a statement to the function
func (f *FuncSpec) Statement(format string, args ...interface{}) *FuncSpec {
	f.Statements = append(f.Statements, Statement{
		Format:    format,
		Arguments: args,
		Indent:    0,
	})

	return f
}

// BlockStart is a convenient method to append a statement that marks the start of a
// block of code.
func (f *FuncSpec) BlockStart(format string, args ...interface{}) *FuncSpec {
	f.Statements = append(f.Statements, Statement{
		Format:    format + " {",
		Arguments: args,
		Indent:    1,
	})

	return f
}

// BlockEnd is a convenient method to append a statement that marks the end of a
// block of code.
func (f *FuncSpec) BlockEnd() *FuncSpec {
	f.Statements = append(f.Statements, Statement{
		Format: "}",
		Indent: -1,
	})

	return f
}

// Parameter is a convenient method to append a parameter to the function
func (f *FuncSpec) Parameter(name string, spec ImportSpec) *FuncSpec {
	f.Parameters = append(f.Parameters, IdentifierParameter{
		Identifier{
			Name: name,
			Type: spec,
		},
	})

	return f
}

// ResultParameter is a convenient method to append a result parameter to the function
func (f *FuncSpec) ResultParameter(name string, spec ImportSpec) *FuncSpec {
	f.ResultParameters = append(f.ResultParameters, IdentifierParameter{
		Identifier{
			Name: name,
			Type: spec,
		},
	})

	return f
}

func (f *FuncSpec) FunctionComment(comment string) *FuncSpec {
	f.Comment = comment

	return f
}
