package gopoet

import . "gopkg.in/check.v1"

type ImportsSuite struct{}

var _ = Suite(&ImportsSuite{})

func (f *ImportsSuite) TestQualifiedImport(c *C) {
	expected := "bytes."
	imp := &ImportSpec{
		Package:   "bytes",
		Qualified: true,
	}
	actual := imp.getQualifier()
	c.Assert(actual, Equals, expected)
}

func (f *ImportsSuite) TestUnqualifiedImport(c *C) {
	expected := ""
	imp := &ImportSpec{
		Package:   "bytes",
		Qualified: false,
	}
	actual := imp.getQualifier()
	c.Assert(actual, Equals, expected)
}
