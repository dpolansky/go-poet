package gopoet

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
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

var _ = (*FuncSpec)(nil)

func NewFuncSpec(name string) *FuncSpec {
	return &FuncSpec{
		Name:             name,
		Parameters:       []IdentifierParameter{},
		ResultParameters: []IdentifierParameter{},
		Statements:       []Statement{},
	}
}

func (f *FuncSpec) String() string {
	writer := NewCodeWriter()

	writer.WriteStatement(f.createSignature())

	for _, st := range f.Statements {
		writer.WriteStatement(st)
	}

	writer.WriteStatement(Statement{
		BeforeIndent: -1,
		Format:       "}",
	})

	return writer.String()
}

func (f *FuncSpec) createSignature() Statement {
	formatStr := bytes.Buffer{}
	arguments := []interface{}{}

	// add comment to front of function
	if f.Comment != "" {
		formatttedComment := strings.Replace(f.Comment, "\n", "\n// ", -1)
		formatStr.WriteString("// ")
		formatStr.WriteString(formatttedComment)
		formatStr.WriteString("\n")
	}
	formatStr.WriteString("func ")
	formatStr.WriteString(f.Name)
	formatStr.WriteString("(")

	for i, param := range f.Parameters {
		formatStr.WriteString("$L $T")
		arguments = append(arguments, param.Name, param.Type)

		if i != len(f.Parameters)-1 {
			formatStr.WriteString(", ")
		}
	}

	formatStr.WriteString(") ")

	if len(f.ResultParameters) == 1 && f.ResultParameters[0].Name == "" {
		formatStr.WriteString("$T")
		arguments = append(arguments, f.ResultParameters[0].Type)
	} else if len(f.ResultParameters) >= 1 {

		formatStr.WriteString("(")
		for i, resultParameter := range f.ResultParameters {
			if resultParameter.Name == "" {
				panic(fmt.Sprintf("Result parameters need a name when there is more than one (got %v)", resultParameter))
			}

			formatStr.WriteString("$L $T")
			arguments = append(arguments, resultParameter.Name, resultParameter.Type)

			if i != len(f.ResultParameters)-1 {
				formatStr.WriteString(", ")
			}
		}
		formatStr.WriteString(") ")
	}

	formatStr.WriteString("{")

	return Statement{
		AfterIndent: 1,
		Format:      formatStr.String(),
		Arguments:   arguments,
	}
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
	})

	return f
}

// BlockStart is a convenient method to append a statement that marks the start of a
// block of code.
func (f *FuncSpec) BlockStart(format string, args ...interface{}) *FuncSpec {
	f.Statements = append(f.Statements, Statement{
		Format:      format + " {",
		Arguments:   args,
		AfterIndent: 1,
	})

	return f
}

// BlockEnd is a convenient method to append a statement that marks the end of a
// block of code.
func (f *FuncSpec) BlockEnd() *FuncSpec {
	f.Statements = append(f.Statements, Statement{
		Format:       "}",
		BeforeIndent: -1,
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

// FunctionComment adds a comment to the function
func (f *FuncSpec) FunctionComment(comment string) *FuncSpec {
	f.Comment = comment

	return f
}
