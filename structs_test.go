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
