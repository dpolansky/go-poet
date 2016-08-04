package gopoet

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestTemplateType(c *C) {
	expected := "fmt.Println()"

	fmtSpec := Import{
		Package: "fmt",
	}

	actual := Template("$T.Println()", fmtSpec)
	c.Assert(actual, Equals, expected)
}

func (s *MySuite) TestTemplateTypeWithName(c *C) {
	expected := "fmt.Println()"

	fmtSpec := Import{
		Package: "fmt",
		Name:    "Println",
	}

	actual := Template("$T()", fmtSpec)
	c.Assert(actual, Equals, expected)
}

func (s *MySuite) TestTemplateWithString(c *C) {
	expected := "fmt.Println(\"Hello World\")"

	fmtSpec := Import{
		Package: "fmt",
		Name:    "Println",
	}

	actual := Template("$T($S)", fmtSpec, "Hello World")
	c.Assert(actual, Equals, expected)
}

func (s *MySuite) TestTemplatePanicsWithNotEnoughArgs(c *C) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	Template("$T()")
	c.Fail()
}
