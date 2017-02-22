package poet

import (
	. "gopkg.in/check.v1"
)

type CodeWriterSuite struct{}

var _ = Suite(&CodeWriterSuite{})

func (f *CodeWriterSuite) TestCodeWriterSingleCode(c *C) {
	expected := "this is a test"
	writer := &codeWriter{}
	writer.WriteCode("this is a test")
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterSingleStatement(c *C) {
	expected := "this is a test\n"
	s := statement{
		Format: "this is a test",
	}
	writer := &codeWriter{}
	writer.WriteStatement(s)
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterPreindentStatement(c *C) {
	expected := "\t\tthis is a test\n"
	s := statement{
		Format:       "this is a test",
		BeforeIndent: 2,
	}
	writer := &codeWriter{}
	writer.WriteStatement(s)
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterMixedIndentStatement(c *C) {
	expected := "\t\tthis is a test\n" +
		"\tstill going\n" +
		"gone\n" +
		"\tbut back\n"

	writer := &codeWriter{}
	writer.WriteStatement(statement{
		Format:       "this is a test",
		BeforeIndent: 2,
	})
	writer.WriteStatement(statement{
		Format:       "still going",
		BeforeIndent: -1,
	})
	writer.WriteStatement(statement{
		Format:       "gone",
		BeforeIndent: -1,
	})
	writer.WriteStatement(statement{
		Format:       "but back",
		BeforeIndent: 1,
	})
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterComment(c *C) {
	expected := "// This is a comment\n"

	writer := &codeWriter{}
	writer.WriteComment("This is a comment")
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterMultiLineComment(c *C) {
	expected := "// This is\n" +
		"// a multi\n" +
		"// line comment\n"

	writer := &codeWriter{}
	writer.WriteComment("This is\na multi\nline comment")
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}
