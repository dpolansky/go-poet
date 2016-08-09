package gopoet

import (
	"bytes"
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type FunctionsSuite struct{}

var _ = Suite(&FunctionsSuite{})

func (f *FunctionsSuite) TestFunction(c *C) {
	expected := "" +
		"func foo() {\n" +
		"}\n"
	actual := NewFuncSpec("foo").String()

	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestFunctionWithParameter(c *C) {
	expected := "" +
		"func foo(a string) {\n" +
		"}\n"
	actual := NewFuncSpec("foo").Parameter("a", TypeReferenceFromInstance("")).String()

	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestFunctionWithParameters(c *C) {
	expected := "" +
		"func foo(a string, b *bytes.Buffer, c int) {\n" +
		"}\n"
	fnc := NewFuncSpec("foo")
	fnc.Parameter("a", TypeReferenceFromInstance(""))
	fnc.Parameter("b", TypeReferenceFromInstance(&bytes.Buffer{}))
	fnc.Parameter("c", TypeReferenceFromInstance(1))

	actual := fnc.String()
	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestFunctionWithReturnParameters(c *C) {
	expected := "" +
		"func foo() (string, int) {\n" +
		"}\n"
	fnc := NewFuncSpec("foo")
	fnc.ResultParameter("", TypeReferenceFromInstance(""))
	fnc.ResultParameter("", TypeReferenceFromInstance(1))

	actual := fnc.String()
	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestFunctionWithNamedReturnParameters(c *C) {
	expected := "" +
		"func foo() (a string, b int) {\n" +
		"}\n"
	fnc := NewFuncSpec("foo")
	fnc.ResultParameter("a", TypeReferenceFromInstance(""))
	fnc.ResultParameter("b", TypeReferenceFromInstance(1))

	actual := fnc.String()
	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestVariadicFunctionParameter(c *C) {
	expected := "" +
		"func foo(bar string...) {\n" +
		"}\n"

	actual := NewFuncSpec("foo").VariadicParameter("bar", TypeReferenceFromInstance("")).String()

	c.Assert(actual, Equals, expected)
}

func (f *FunctionsSuite) TestFunctionImportsFromStatement(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "fmt",
			Qualified: true,
		},
	}

	fnc := NewFuncSpec("foo")
	fnc.Statement("$T($S)", TypeReferenceFromInstance(fmt.Println), "Test")

	actual := fnc.GetImports()
	c.Assert(actual, DeepEquals, expected)
}

func (f *FunctionsSuite) TestFunctionImportsFromParameter(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "bytes",
			Qualified: true,
		},
	}

	fnc := NewFuncSpec("foo")
	fnc.Parameter("a", TypeReferenceFromInstance(&bytes.Buffer{}))

	actual := fnc.GetImports()
	c.Assert(actual, DeepEquals, expected)
}

func (f *FunctionsSuite) TestFunctionImportsFromReturnParameter(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "bytes",
			Qualified: true,
		},
	}

	fnc := NewFuncSpec("foo")
	fnc.ResultParameter("a", TypeReferenceFromInstance(&bytes.Buffer{}))

	actual := fnc.GetImports()
	c.Assert(actual, DeepEquals, expected)
}

func (f *FunctionsSuite) TestFunctionImportsWithAlias(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "bytes",
			Qualified: true,
			Alias:     "test",
		},
	}

	fnc := NewFuncSpec("foo")
	fnc.ResultParameter("a", TypeReferenceFromInstanceWithAlias(&bytes.Buffer{}, "test"))

	actual := fnc.GetImports()
	c.Assert(actual, DeepEquals, expected)
}
