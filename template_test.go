package gopoet

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestGoPoet(t *testing.T) { TestingT(t) }

type TemplateSuite struct{}

var _ = Suite(&TemplateSuite{})

func (s *TemplateSuite) TestTemplateTypeWithName(c *C) {
	expected := "fmt.Println()"

	fmtSpec := TypeReferenceFromInstance(fmt.Println)

	actual := Template("$T()", fmtSpec)
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestTemplateWithString(c *C) {
	expected := "fmt.Println(\"Hello World\")"

	fmtSpec := TypeReferenceFromInstance(fmt.Println)

	actual := Template("$T($S)", fmtSpec, "Hello World")
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestTemplatePanicsWithNotEnoughArgs(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	Template("$T()")
	c.Fail()
}
