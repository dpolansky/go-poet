package poet

import (
	. "gopkg.in/check.v1"
)

type CommentSuite struct{}

var _ = Suite(&CommentSuite{})

func (s *CommentSuite) TestCommentEmpty(c *C) {
	comment := Comment("")
	expected := ""

	c.Assert(comment.String(), Equals, expected)
	c.Assert(comment.GetImports(), IsNil)
}

func (s *CommentSuite) TestCommentSimgleLine(c *C) {
	comment := Comment("this is a comment")
	expected := "// this is a comment\n"

	c.Assert(comment.String(), Equals, expected)
	c.Assert(comment.GetImports(), IsNil)
}

func (s *CommentSuite) TestCommentMultiLine(c *C) {
	comment := Comment("this\nis\n a\n  comment")
	expected := "// this\n// is\n//  a\n//   comment\n"

	c.Assert(comment.String(), Equals, expected)
	c.Assert(comment.GetImports(), IsNil)
}

func (s *CommentSuite) TestCommentEmptyLines(c *C) {
	comment := Comment("this is\n\na comment")
	expected := "// this is\n//\n// a comment\n"

	c.Assert(comment.String(), Equals, expected)
	c.Assert(comment.GetImports(), IsNil)
}
