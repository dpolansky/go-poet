package gopoet

import "bytes"

type CodeWriter struct {
	buffer        bytes.Buffer
	currentIndent int
}

func NewCodeWriter() *CodeWriter {
	return &CodeWriter{
		buffer: bytes.Buffer{},
	}
}

func (c *CodeWriter) WriteCode(code string) {
	for i := 0; i < c.currentIndent; i++ {
		c.buffer.WriteString("\t")
	}
	c.buffer.WriteString(code)
}

func (c *CodeWriter) WriteStatement(s Statement) {
	c.currentIndent += s.BeforeIndent
	c.WriteCode(Template(s.Format, s.Arguments...) + "\n")
	c.currentIndent += s.AfterIndent
}

func (c *CodeWriter) String() string {
	return c.buffer.String()
}
