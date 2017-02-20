package poet

import (
	"bytes"
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
	"strings"
)

func _(t *testing.T) { TestingT(t) }

type FilesSuite struct{}

var _ = Suite(&FilesSuite{})

func (f *FilesSuite) TestFilePackage(c *C) {
	expected := "" +
		"package foo\n\n"
	actual := NewFileSpec("foo").String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileImports(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"import (\n" +
		"\t\"bytes\"\n" +
		")\n" +
		"\n" +
		"func blah(a *bytes.Buffer) {\n" +
		"}\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fnc := NewFuncSpec("blah").Parameter("a", TypeReferenceFromInstance(&bytes.Buffer{}))

	fspec.CodeBlock(fnc)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileEmptyImports(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"func blah(a string) {\n" +
		"}\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fnc := NewFuncSpec("blah").Parameter("a", String)

	fspec.CodeBlock(fnc)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileOnlyIncludesExternalImports(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"import (\n" +
		"\t\"bytes\"\n" +
		")\n" +
		"\n" +
		"func blah(a string, b *bytes.Buffer) {\n" +
		"}\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fnc := NewFuncSpec("blah").Parameter("a", String).Parameter("b", TypeReferenceFromInstance(&bytes.Buffer{}))

	fspec.CodeBlock(fnc)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileImportsWithAlias(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"import (\n" +
		"\tblah \"bytes\"\n" +
		")\n" +
		"\n" +
		"func blah(a *blah.Buffer) {\n" +
		"}\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fnc := NewFuncSpec("blah").Parameter("a", TypeReferenceFromInstanceWithAlias(&bytes.Buffer{}, "blah"))

	fspec.CodeBlock(fnc)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileInitializationImports(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"import (\n" +
		"\t_ \"bytes\"\n" +
		")\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileRepeatedImports(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"import (\n" +
			"\t_ \"bytes\"\n" +
		")\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileMultipleImports(c *C) {
	fspec := NewFileSpec("foo")
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})
	fspec.InitializationPackage(&ImportSpec{
		Package: "bytes",
	})
	fspec.InitializationPackage(&ImportSpec{
		Package: "context",
	})
	fspec.CodeBlock(NewFuncSpec("blah").
		Parameter("a", TypeReferenceFromInstance(&bytes.Buffer{})).
		Parameter("b", TypeReferenceFromInstanceWithAlias(&bytes.Buffer{}, "blah")))

	actual := fspec.String()
	c.Assert(strings.Count(actual, "\t_ \"bytes\"\n"), Equals, 1)
	c.Assert(strings.Count(actual, "\t_ \"context\"\n"), Equals, 1)
	c.Assert(strings.Count(actual, "\tblah \"bytes\"\n"), Equals, 1)
	c.Assert(strings.Count(actual, "\t\"bytes\"\n"), Equals, 1)
}

func (f *FilesSuite) TestFileGlobalVariable(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"var a int = 5\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fspec.GlobalVariable("a", TypeReferenceFromInstance(1), "$L", 5)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileGlobalConstant(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"const a int = 5\n" +
		"\n"

	fspec := NewFileSpec("foo")
	fspec.GlobalConstant("a", TypeReferenceFromInstance(1), "$L", 5)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileGlobalVariableGrouping(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"const (\n" +
		"\tc int = 1\n" +
		")\n" +
		"\n" +
		"var (\n" +
		"\td int = 1\n" +
		")\n" +
		"\n"

	variableA := &Variable{
		Identifier: Identifier{
			Name: "c",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1}, Constant: true,
	}
	variableB := &Variable{
		Identifier: Identifier{
			Name: "d",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1}, Constant: false,
	}

	fspec := NewFileSpec("foo")
	grouping := fspec.VariableGrouping()
	grouping.Variables = append(grouping.Variables, variableA, variableB)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileInitFunction(c *C) {
	expected := "" +
		"package foo\n" +
		"\n" +
		"func init() {\n" +
		"\tfmt.Println(\"Init\")\n" +
		"}\n" +
		"\n"

	fspec := NewFileSpec("foo")
	initFunc := NewFuncSpec("init").Statement("$T($S)", TypeReferenceFromInstance(fmt.Println), "Init")
	fspec.InitFunction(initFunc)

	actual := fspec.String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileInitFunctionPanicsWithWrongName(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	fspec := NewFileSpec("foo")
	initFunc := NewFuncSpec("bar")
	fspec.InitFunction(initFunc)

	c.Fail()
}

func (f *FilesSuite) TestFileComment(c *C) {
	expected := "" +
		"// This is a comment.\n" +
		"package foo\n" +
		"\n"

	actual := NewFileSpec("foo").FileComment("This is a comment.").String()
	c.Assert(actual, Equals, expected)
}

func (f *FilesSuite) TestFileDoesNotShowEmptyComment(c *C) {
	expected := "" +
		"package foo\n" +
		"\n"

	actual := NewFileSpec("foo").String()
	c.Assert(actual, Equals, expected)
}
