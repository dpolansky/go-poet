package gopoet

import (
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type FunctionsSuite struct{}

var _ = Suite(&FunctionsSuite{})

func (f *FunctionsSuite) TestVariadicFunctionParameter(c *C) {
	expected :=
		`func foo(bar string...) {
}
`
	actual := NewFuncSpec("foo").VariadicParameter("bar", TypeReferenceFromInstance("")).String()

	c.Assert(actual, Equals, expected)
}
