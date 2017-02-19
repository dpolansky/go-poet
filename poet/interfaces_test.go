package poet

import (
	"io"

	. "gopkg.in/check.v1"
)

type InterfaceSuite struct{}

var _ = Suite(&InterfaceSuite{})

func (f *InterfaceSuite) TestBlankInterface(c *C) {
	expected := "type NewInterface interface {\n" +
		"}\n"

	i := NewInterfaceSpec("NewInterface")
	actual := i.String()
	c.Assert(actual, Equals, expected)
	c.Assert(i.GetName(), Equals, "NewInterface")
}

func (f *InterfaceSuite) TestBlankInterfaceWithComment(c *C) {
	expected := "// NewInterface is a cool interface\n" +
		"type NewInterface interface {\n" +
		"}\n"

	i := InterfaceSpec{Name: "NewInterface", Comment: "NewInterface is a cool interface"}
	actual := i.String()
	c.Assert(actual, Equals, expected)
}

func (f *InterfaceSuite) TestInterfaceWithMethods(c *C) {
	expected := "type NewInterface interface {\n" +
		"\t// TestA does stuff\n" +
		"\tTestA(paramA string)\n" +
		"\tTestB()\n" +
		"}\n"

	i := NewInterfaceSpec("NewInterface")
	i.
		Method(
			NewFuncSpec("TestA").
				FunctionComment("TestA does stuff").
				Parameter("paramA", TypeReferenceFromInstance("")),
		).
		Method(NewFuncSpec("TestB"))

	actual := i.String()
	c.Assert(actual, Equals, expected)
}

func (f *InterfaceSuite) TestInterfaceEmbeddedInterface(c *C) {
	expected := "type NewInterface interface {\n" +
		"\tio.Writer\n" +
		"}\n"

	i := NewInterfaceSpec("NewInterface").EmbedInterface(TypeReferenceFromInstance((*io.Writer)(nil)))

	actual := i.String()
	c.Assert(actual, Equals, expected)
}

func (f *InterfaceSuite) TestInterfaceImportsFromEmbeddedInterfaces(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "io",
			Qualified: true,
		},
	}

	i := NewInterfaceSpec("NewInterface").EmbedInterface(TypeReferenceFromInstance((*io.Writer)(nil)))

	actual := i.GetImports()
	c.Assert(actual, DeepEquals, expected)
}

func (f *InterfaceSuite) TestInterfaceImportsFromMethods(c *C) {
	expected := []Import{
		&ImportSpec{
			Package:   "io",
			Qualified: true,
		},
	}

	i := NewInterfaceSpec("NewInterface")
	i.
		Method(
			NewFuncSpec("TestA").
				FunctionComment("TestA does stuff").
				Parameter("paramA", TypeReferenceFromInstance((*io.Writer)(nil))),
		).
		Method(NewFuncSpec("TestB"))

	actual := i.GetImports()
	c.Assert(actual, DeepEquals, expected)
}
