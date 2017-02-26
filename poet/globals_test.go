package poet

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
			Type: Int,
		},
		Value: newStatement(0, 0, "$L", 1),
	}
	actual := variable.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestVariableNoValue(c *C) {
	expected := "var c int\n"
	variable := &Variable{
		Identifier: Identifier{
			Name: "c",
			Type: Int,
		},
	}
	actual := variable.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestVariableWithComment(c *C) {
	expected := "// c is an int\n" +
		"// with value 1\n" +
		"var c int = 1\n"
	variable := &Variable{
		Comment: "c is an int\nwith value 1",
		Identifier: Identifier{
			Name: "c",
			Type: Int,
		},
		Value: newStatement(0, 0, "$L", 1),
	}
	actual := variable.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestVariableNoValueWithComment(c *C) {
	expected := "// c is an int\n" +
		"var c int\n"
	variable := &Variable{
		Comment: "c is an int",
		Identifier: Identifier{
			Name: "c",
			Type: Int,
		},
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
			Type: Int,
		},
		InGroup: true,
		Value:   newStatement(0, 0, "$L", 1),
	}
	variableB := &Variable{
		Identifier: Identifier{
			Name: "d",
			Type: Int,
		},
		InGroup: true,
		Value:   newStatement(0, 0, "$L", 1),
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
			Type: Int,
		},
		Constant: true,
		Value:    newStatement(0, 0, "$L", 1),
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
			Type: Int,
		},
		Constant: true,
		InGroup:  true,
		Value:    newStatement(0, 0, "$L", 1),
	}
	variableB := &Variable{
		Identifier: Identifier{
			Name: "d",
			Type: Int,
		},
		Constant: true,
		InGroup:  true,
		Value:    newStatement(0, 0, "$L", 1),
	}
	variableGrouping := VariableGrouping{Variables: []*Variable{variableA, variableB}}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestConstantGroupingWithComment(c *C) {
	expected := "const (\n" +
		"\t// c has a value of 1\n" +
		"\tc int = 1\n" +
		"\t// d has\n" +
		"\t// no value yet\n" +
		"\td int\n" +
		")\n"

	variableA := &Variable{
		Comment: "c has a value of 1",
		Identifier: Identifier{
			Name: "c",
			Type: Int,
		},
		Constant: true,
		InGroup:  true,
		Value:    newStatement(0, 0, "$L", 1),
	}
	variableB := &Variable{
		Comment: "d has\n" +
			"no value yet",
		Identifier: Identifier{
			Name: "d",
			Type: Int,
		},
		Constant: true,
		InGroup:  true,
	}
	variableGrouping := VariableGrouping{Variables: []*Variable{variableA, variableB}}
	actual := variableGrouping.String()

	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestGroupingWithAttachedConstants(c *C) {
	expected := "const (\n" +
		"\tc int = 1\n" +
		"\td int = 1\n" +
		")\n"

	variableGrouping := &VariableGrouping{}
	variableGrouping.Constant("c", Int, "$L", 1)
	variableGrouping.Constant("d", Int, "$L", 1)

	actual := variableGrouping.String()
	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestGroupingWithAttachedVariables(c *C) {
	expected := "var (\n" +
		"\tc int = 1\n" +
		"\td int = 1\n" +
		")\n"

	variableGrouping := &VariableGrouping{}
	variableGrouping.Variable("c", Int, "$L", 1)
	variableGrouping.Variable("d", Int, "$L", 1)

	actual := variableGrouping.String()
	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestGroupingWithAttachedMixed(c *C) {
	expected := "const (\n" +
		"\tc int = 1\n" +
		")\n" +
		"\n" +
		"var (\n" +
		"\td int = 1\n" +
		")\n"

	variableGrouping := &VariableGrouping{}
	variableGrouping.Constant("c", Int, "$L", 1)
	variableGrouping.Variable("d", Int, "$L", 1)

	actual := variableGrouping.String()
	c.Assert(actual, Equals, expected)
}

func (f *VariablesSuite) TestGroupingWithAttachedMixedNoValue(c *C) {
	expected := "const (\n" +
		"\tc int = 1\n" +
		")\n" +
		"\n" +
		"var (\n" +
		"\td int\n" +
		")\n"

	variableGrouping := &VariableGrouping{}
	variableGrouping.Constant("c", Int, "$L", 1)
	variableGrouping.Variable("d", Int, "")

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
			Type: Int,
		},
		Constant: true,
		InGroup:  true,
		Value:    newStatement(0, 0, "$L", 1),
	}
	variableB := &Variable{
		Identifier: Identifier{
			Name: "d",
			Type: Int,
		},
		InGroup: true,
		Value:   newStatement(0, 0, "$L", 1),
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
