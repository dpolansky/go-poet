package poet

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

	actual := template("$T()", fmtSpec)
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestTemplateWithString(c *C) {
	expected := "fmt.Println(\"Hello World\")"

	fmtSpec := TypeReferenceFromInstance(fmt.Println)

	actual := template("$T($S)", fmtSpec, "Hello World")
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestTemplateWithIntLiterals(c *C) {
	expected := "literal 1 string \"2\" type int"
	actual := template("literal $L string $S type $T", 1, 2, TypeReferenceFromInstance(3))
	c.Assert(actual, Equals, expected)
}

func (s *TemplateSuite) TestTemplatePanicsWithNotEnoughArgs(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	template("$T()")
	c.Fail()
}

func (s *TemplateSuite) TestTemplatePanicsWithNonTypeReference(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	template("$T()", 1)
	c.Fail()
}

func (s *TemplateSuite) TestTemplatePanicsWithInvalidTemplatingString(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	template("$D()", TypeReferenceFromInstance(fmt.Println))
	c.Fail()
}
