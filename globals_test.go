package gopoet

import (
	"testing"

	. "gopkg.in/check.v1"
)

func _(t *testing.T) { TestingT(t) }

type VariablesSuite struct{}

var _ = Suite(&VariablesSuite{})

func (f *VariablesSuite) TestVariable(c *C) {
	expected := "var c int = 1\n"
	variable := &Variable{
		Identifier: Identifier{
			Name: "c",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1},
	}
	actual := variable.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestVariableGrouping(c *C) {
	expected := "var (\n" +
		"\tc int = 1\n" +
		"\td int = 1\n" +
		")\n"
	variableA := &Variable{
		Identifier: Identifier{
			Name: "c",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1},
	}
	variableB := &Variable{
		Identifier: Identifier{
			Name: "d",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1},
	}
	variableGrouping := VariableGrouping{Variables: []*Variable{variableA, variableB}}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestConstant(c *C) {
	expected := "const c int = 1\n"
	variable := &Variable{
		Identifier: Identifier{
			Name: "c",
			Type: TypeReferenceFromInstance(1),
		},
		Format: "$L", Args: []interface{}{1}, Constant: true,
	}
	actual := variable.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestConstantGrouping(c *C) {
	expected := "const (\n" +
		"\tc int = 1\n" +
		"\td int = 1\n" +
		")\n"
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
		Format: "$L", Args: []interface{}{1}, Constant: true,
	}
	variableGrouping := VariableGrouping{Variables: []*Variable{variableA, variableB}}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestConstantGroupingMixed(c *C) {
	expected := "const (\n" +
		"\tc int = 1\n" +
		")\n" +
		"\n" +
		"var (\n" +
		"\td int = 1\n" +
		")\n"
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
	variableGrouping := VariableGrouping{Variables: []*Variable{variableA, variableB}}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestGroupingEmpty(c *C) {
	expected := ""

	variableGrouping := VariableGrouping{}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}
