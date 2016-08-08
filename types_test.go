package gopoet

import (
	"bytes"
	"fmt"
	IoAlias "io"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestTypes(t *testing.T) { TestingT(t) }

type TypeSuite struct{}

var _ = Suite(&TypeSuite{})

func (s *TypeSuite) TestFunctionRef(c *C) {
	expected := "fmt.Println"
	typeRef := TypeReferenceFromInstance(fmt.Println)
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

func (s *TypeSuite) TestPrimitive(c *C) {
	expected := "int"
	typeRef := TypeReferenceFromInstance(1)
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestBoolean(c *C) {
	expected := "bool"
	typeRef := TypeReferenceFromInstance(false)
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
}

func (s *TypeSuite) TestByte(c *C) {
	expected := "byte"
	typeRef := TypeReferenceFromInstanceWithAlias(byte('A'), "byte")
	actual := typeRef.GetName()

	c.Assert(actual, Equals, expected)
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
