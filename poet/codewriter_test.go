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
	s := Statement{
		Format: "this is a test",
	}
	writer := &codeWriter{}
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

func (f *CodeWriterSuite) TestCodeWriterCodeBlock(c *C) {
	expected := "" +
		"type foo struct {\n" +
		"}\n"
	writer := &codeWriter{}
	writer.WriteCodeBlock(NewStructSpec("foo"))
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}

func (f *CodeWriterSuite) TestCodeWriterNewStatement(c *C) {
	expected := "this is a test\n"
	writer := &codeWriter{}
	writer.WriteStatement(newStatement(0, 0, "$L $L $L $L", "this", "is", "a", "test"))
	actual := writer.String()

	c.Assert(actual, Equals, expected)
}
