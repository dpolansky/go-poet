package gopoet

import (
	"bytes"

	. "gopkg.in/check.v1"
)

type StructsSuite struct{}

var _ = Suite(&StructsSuite{})

func (s *StructsSuite) TestStruct(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"}\n"

	st := NewStructSpec("foo")

	actual := st.String()
	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructWithFields(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"\tbar string\n" +
		"\tbaz *bytes.Buffer\n" +
		"}\n"

	st := NewStructSpec("foo")
	st.Field("bar", TypeReferenceFromInstance(""))
	st.Field("baz", TypeReferenceFromInstance(&bytes.Buffer{}))

	actual := st.String()
	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructWithFieldsWithTags(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"\tbar string `json:\"bar\"`\n" +
		"}\n"

	st := NewStructSpec("foo")
	st.FieldWithTag("bar", TypeReferenceFromInstance(""), "json:\"bar\"")

	actual := st.String()
	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructGetImports(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "bytes",
			Qualified: true,
		},
	}

	st := NewStructSpec("foo")
	st.Field("baz", TypeReferenceFromInstance(&bytes.Buffer{}))

	actual := st.GetImports()
	c.Assert(actual, DeepEquals, expected)
}

func (s *StructsSuite) TestStructWithVarTag(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"\tbar string `json:\"bar\"`\n" +
		"}\n"

	st := NewStructSpec("foo")
	st.FieldWithTag("bar", TypeReferenceFromInstance(""), "json:\"bar\"")

	actual := st.String()

	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructWithComment(c *C) {
	expected := "" +
		"// Test comment\n" +
		"type foo struct {\n" +
		"}\n"

	st := NewStructSpec("foo")
	st.StructComment("Test comment")

	actual := st.String()

	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestMethodFromStruct(c *C) {
	expected := "" +
		"func (f foo) bar() {\n" +
		"}\n"

	st := NewStructSpec("foo")
	m := st.Method("bar", "f")

	actual := m.String()

	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructWithMethodAttached(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"}\n" +
		"\n" +
		"func (f foo) bar() {\n" +
		"}\n" +
		"\n"

	st := NewStructSpec("foo")
	st.MethodAndAttach("bar", "f")

	actual := st.String()

	c.Assert(actual, Equals, expected)
}

func (s *StructsSuite) TestStructName(c *C) {
	expected := "foo"
	st := NewStructSpec("foo")
	actual := st.GetName()
	c.Assert(actual, Equals, expected)
}
