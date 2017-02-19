package poet

import (
	"bytes"
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type AliasSuite struct{}

var _ = Suite(&AliasSuite{})

func (f *AliasSuite) TestAlias(c *C) {
	expected := "type foo string\n"
	actual := NewAliasSpec("foo", String).String()

	c.Assert(actual, Equals, expected)
}

func (f *AliasSuite) TestAliasWithComment(c *C) {
	expected := "// foo is an alias\ntype foo string\n"
	actual := NewAliasSpec("foo", String).AliasComment("foo is an alias").String()

	c.Assert(actual, Equals, expected)
}

func (f *AliasSuite) TestAliasWithImport(c *C) {
	spec := NewAliasSpec("foo", TypeReferenceFromInstance(&bytes.Buffer{}))
	c.Assert(spec.String(), Equals, "type foo *bytes.Buffer\n")
	c.Assert(spec.GetImports()[0].GetPackage(), Equals, "bytes")
}
