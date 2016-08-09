package gopoet

import (
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type StructSuite struct{}

var _ = Suite(&StructSuite{})

func (f *StructSuite) TestStructWithVarTag(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"\tbar string `json:\"bar\"`\n" +
		"}\n"

	s := NewStructSpec("foo")
	s.FieldWithTag("bar", TypeReferenceFromInstance(""), "json:\"bar\"")

	actual := s.String()

	c.Assert(actual, Equals, expected)
}
