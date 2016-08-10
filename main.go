package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/dpolansky/go-poet/poet"
)

func main() {
	add := poet.NewFuncSpec("add").
		Parameter("a", poet.TypeReferenceFromInstance(int(0))).
		Parameter("b", poet.TypeReferenceFromInstance(int(0))).
		ResultParameter("", poet.TypeReferenceFromInstance(int(0))).
		Statement("return a + b")

	foo := poet.NewStructSpec("foo").
		Field("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{})).
		AttachFunction("f", add)

	inter := poet.NewInterfaceSpec("AddWriter").
		EmbedInterface(poet.TypeReferenceFromInstance((*io.Writer)(nil))).
		Method(add)

	fooFunc := poet.NewFuncSpec("foo").Parameter("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{}))

	file := poet.NewFileSpec("main").
		CodeBlock(add).
		CodeBlock(inter).
		CodeBlock(foo).
		CodeBlock(fooFunc)

	file.VariableGrouping().
		Variable("a", poet.TypeReferenceFromInstance(0), "$L", 7).
		Variable("b", poet.TypeReferenceFromInstance(0), "$L", 2).
		Constant("c", poet.TypeReferenceFromInstance(0), "$L", 3).
		Constant("d", poet.TypeReferenceFromInstance(0), "$L", 43)

	fmt.Println(file)
}
