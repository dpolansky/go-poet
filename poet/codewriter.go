package poet

import (
	"bytes"
	"strings"
)

// codeWriter keeps track of the current indentation and writes code to a buffer
type codeWriter struct {
	buffer        bytes.Buffer
	currentIndent int
}

// newCodeWriter constructs a new codeWriter
func newCodeWriter() *codeWriter {
	return &codeWriter{
		buffer: bytes.Buffer{},
	}
}

// WriteCode writes code at the given indentation
func (c *codeWriter) WriteCode(code string) {
	c.buffer.WriteString(strings.Repeat("\t", c.currentIndent))
	c.buffer.WriteString(code)
}

// WriteCodeBlock writes a code block at the given indentation
func (c *codeWriter) WriteCodeBlock(block CodeBlock) {
	c.WriteCode(block.String())
}

// WriteStatement writes a new line of code with the current indentation and augments
// the indentation per the statement. A newline is appended at the end of the statement.
func (c *codeWriter) WriteStatement(s Statement) {
	c.currentIndent += s.BeforeIndent
	c.WriteCode(template(s.Format, s.Arguments...) + "\n")
	c.currentIndent += s.AfterIndent
}

// String gives a string with the code
func (c *codeWriter) String() string {
	return c.buffer.String()
}

func newStatement(beforeIndent, afterIndent int, format string, args ...interface{}) Statement {
	return Statement{
		BeforeIndent: beforeIndent,
		AfterIndent:  afterIndent,
		Format:       format,
		Arguments:    args,
	}
}
