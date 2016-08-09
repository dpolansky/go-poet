package gopoet

import (
	"bytes"

	. "gopkg.in/check.v1"
)

type MethodSuite struct{}

var _ = Suite(&MethodSuite{})

func (s *TemplateSuite) TestMethod(c *C) {
	expected := "" +
		"func (b *bytes.Buffer) foo() {\n" +
		"}\n"

	m := NewMethodSpec("foo", "b", TypeReferenceFromInstance(&bytes.Buffer{}))

	actual := m.String()
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestMethodValueReceiver(c *C) {
	expected := "" +
		"func (b bytes.Buffer) foo() {\n" +
		"}\n"

	m := NewMethodSpec("foo", "b", TypeReferenceFromInstance(bytes.Buffer{}))

	actual := m.String()
	c.Assert(actual, Equals, expected)
}
