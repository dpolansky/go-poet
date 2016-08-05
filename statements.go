package gopoet

type Statement struct {
	// CodeBlock

	Format       string
	Arguments    []interface{}
	BeforeIndent int
	AfterIndent  int
}
