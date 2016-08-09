package poet

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

func (f *ImportsSuite) TestAliasedImport(c *C) {
	expected := "blah."
	imp := &ImportSpec{
		Package:   "bytes",
		Qualified: true,
		Alias:     "blah",
	}
	actual := imp.getQualifier()
	c.Assert(imp.GetAlias(), Equals, "blah")
	c.Assert(imp.GetPackage(), Equals, "bytes")
	c.Assert(actual, Equals, expected)
}

func (f *ImportsSuite) TestNils(c *C) {
	nilInst := ((*ImportSpec)(nil))
	c.Assert(nilInst.GetAlias(), Equals, "")
	c.Assert(nilInst.GetPackage(), Equals, "")
}
