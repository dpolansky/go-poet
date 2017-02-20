package poet

import (
	"bytes"
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type TypeAliasSuite struct{}

var _ = Suite(&TypeAliasSuite{})

func (s *TypeAliasSuite) TestAlias(c *C) {
	expected := "type foo string\n"
	actual := NewTypeAliasSpec("foo", String).String()
	c.Assert(actual, Equals, expected)
}

func (s *TypeAliasSuite) TestAliasWithComment(c *C) {
	expected := "// foo is an alias\ntype foo string\n"
	actual := NewTypeAliasSpec("foo", String).AliasComment("foo is an alias").String()
	c.Assert(actual, Equals, expected)
}

func (s *TypeAliasSuite) TestAliasWithImport(c *C) {
	spec := NewTypeAliasSpec("foo", TypeReferenceFromInstance(&bytes.Buffer{}))
	c.Assert(spec.String(), Equals, "type foo *bytes.Buffer\n")
	c.Assert(spec.GetImports()[0].GetPackage(), Equals, "bytes")
}
