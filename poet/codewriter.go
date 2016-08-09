package poet

import (
	"bytes"
	"strings"
)

// CodeWriter keeps track of the current indentation and writes code to a buffer
type CodeWriter struct {
	buffer        bytes.Buffer
	currentIndent int
}

// NewCodeWriter constructs a new CodeWriter
func NewCodeWriter() *CodeWriter {
	return &CodeWriter{
		buffer: bytes.Buffer{},
	}
}

// WriteCode writes code at the given indentation
func (c *CodeWriter) WriteCode(code string) {
	c.buffer.WriteString(strings.Repeat("\t", c.currentIndent))
	c.buffer.WriteString(code)
}

// WriteStatement writes a new line of code with the current indentation and augments the identation per the statement.
func (c *CodeWriter) WriteStatement(s Statement) {
	c.currentIndent += s.BeforeIndent
	c.WriteCode(Template(s.Format, s.Arguments...) + "\n")
	c.currentIndent += s.AfterIndent
}

// String gives a string with the code
func (c *CodeWriter) String() string {
	return c.buffer.String()
}
