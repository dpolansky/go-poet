package poet

import (
	"bytes"
	"fmt"
	IoAlias "io"
	"os"

	"golang.org/x/net/context"
	. "gopkg.in/check.v1"
)

type TypeSuite struct{}

var _ = Suite(&TypeSuite{})

func (s *TypeSuite) TestFunctionRef(c *C) {
	expected := "fmt.Println"
	typeRef := TypeReferenceFromInstance(fmt.Println)
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestFunctionRefFromExternal(c *C) {
	expected := "context.Background"
	typeRef := TypeReferenceFromInstance(context.Background)
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestExternalStructPointer(c *C) {
	expected := "*bytes.Buffer"
	typeRef := TypeReferenceFromInstance(&bytes.Buffer{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestExternalStruct(c *C) {
	expected := "bytes.Buffer"
	typeRef := TypeReferenceFromInstance(bytes.Buffer{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestAliasedExternalStruct(c *C) {
	expected := "io.SectionReader"
	typeRef := TypeReferenceFromInstance(IoAlias.SectionReader{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestMap(c *C) {
	expected := "map[string]string"
	typeRef := TypeReferenceFromInstance(map[string]string{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestMapPointer(c *C) {
	expected := "*map[string]string"
	typeRef := TypeReferenceFromInstance(&map[string]string{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestMapPointerPointer(c *C) {
	expected := "**map[string]string"
	m := &map[string]string{}
	typeRef := TypeReferenceFromInstance(&m)
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestMapImports(c *C) {
	expected := []Import{
		(*ImportSpec)(nil),
		&ImportSpec{
			Package:   "bytes",
			Qualified: true,
		},
	}
	typeRef := TypeReferenceFromInstance(map[string]*bytes.Buffer{})
	actual := typeRef.GetImports()

	c.Assert(actual, DeepEquals, expected)
}

func (s *TypeSuite) TestArray(c *C) {
	expected := "[]int"
	typeRef := TypeReferenceFromInstance([]int{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestInterface(c *C) {
	expected := "os.Signal"
	typeRef := TypeReferenceFromInstance((*os.Signal)(nil))
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestUnqualified(c *C) {
	type _unqualifiedBuffer bytes.Buffer
	expected := "Buffer"
	typeRef := TypeReferenceFromInstance(_unqualifiedBuffer{})
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestChannel(c *C) {
	expected := "chan *bytes.Buffer"
	typeRef := TypeReferenceFromInstance(make(chan *bytes.Buffer))
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestChannelOneDirection(c *C) {
	expected := "chan<- *bytes.Buffer"
	typeRef := TypeReferenceFromInstance(make(chan<- *bytes.Buffer))
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestTypeReferencePanicsWithNilInstance(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	TypeReferenceFromInstance(nil)
	c.Fail()
}

func (s *TypeSuite) TestTypePrimitives(c *C) {
	for _, test := range []struct {
		ref  TypeReference
		name string
	}{
		{
			String,
			"string",
		},
		{
			Bool,
			"bool",
		},
		{
			Int,
			"int",
		},
		{
			Int8,
			"int8",
		},
		{
			Int16,
			"int16",
		},
		{
			Int32,
			"int32",
		},
		{
			Int64,
			"int64",
		},
		{
			Uint,
			"uint",
		},
		{
			Uint8,
			"uint8",
		},
		{
			Uint16,
			"uint16",
		},
		{
			Uint32,
			"uint32",
		},
		{
			Uint64,
			"uint64",
		},
		{
			Uintptr,
			"uintptr",
		},
		{
			Float32,
			"float32",
		},
		{
			Float64,
			"float64",
		},
		{
			Complex64,
			"complex64",
		},
		{
			Complex128,
			"complex128",
		},
		{
			Byte,
			"byte",
		},
		{
			Rune,
			"rune",
		},
		{
			Error,
			"error",
		},
	} {
		c.Check(test.ref.GetName(), Equals, test.name)
	}
}
