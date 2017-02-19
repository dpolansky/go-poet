package poet

import (
	"strings"
)

var _ CodeBlock = (*Comment)(nil)

// Comment represents a comment and implements CodeBlock. Multiline comments
// are supported, with '// ' being prepended to each line.
type Comment string

// GetImports returns nil.
func (c Comment) GetImports() []Import {
	return nil
}

// String returns a string with all lines prepended with '// '
func (c Comment) String() string {
	w := newCodeWriter()
	for _, s := range c.GetStatements() {
		w.WriteStatement(s)
	}
	return w.String()
}

// GetStatements returns the comment as statements prepended with '// '.
func (c Comment) GetStatements() []Statement {
	if string(c) == "" {
		return nil
	}

	lines := strings.Split(string(c), "\n")
	statements := make([]Statement, len(lines))
	for i, line := range lines {
		if line != "" {
			statements[i] = newStatement(0, 0, "// $L", line)
		} else {
			statements[i] = newStatement(0, 0, "//")
		}
	}
	return statements
}
