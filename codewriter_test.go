package gopoet

import . "gopkg.in/check.v1"

type CodeWriterSuite struct{}

var _ = Suite(&CodeWriterSuite{})

func (f *CodeWriterSuite) TestCodeWriterSingleCode(c *C) {
	expected := "this is a test"
	writer := CodeWriter{}
	writer.WriteCode("this is a test")
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterSingleStatement(c *C) {
	expected := "this is a test\n"
	s := Statement{
		Format: "this is a test",
	}
	writer := CodeWriter{}
	writer.WriteStatement(s)
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterPreindentStatement(c *C) {
	expected := "\t\tthis is a test\n"
	s := Statement{
		Format:       "this is a test",
		BeforeIndent: 2,
	}
	writer := CodeWriter{}
	writer.WriteStatement(s)
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterMixedIndentStatement(c *C) {
	expected := "\t\tthis is a test\n" +
		"\tstill going\n" +
		"gone\n" +
		"\tbut back\n"

	writer := CodeWriter{}
	writer.WriteStatement(Statement{
		Format:       "this is a test",
		BeforeIndent: 2,
	})
	writer.WriteStatement(Statement{
		Format:       "still going",
		BeforeIndent: -1,
	})
	writer.WriteStatement(Statement{
		Format:       "gone",
		BeforeIndent: -1,
	})
	writer.WriteStatement(Statement{
		Format:       "but back",
		BeforeIndent: 1,
	})
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}
