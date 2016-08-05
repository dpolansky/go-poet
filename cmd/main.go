package main

import (
	"bytes"
	"fmt"

	"gitlab.com/kdeenanauth/gopoet"
)

func main() {
	f := gopoet.NewFuncSpec("main")
	fmtImport := gopoet.Import{
		Package: "fmt",
		Name:    "Println",
	}

	f.Parameter("myStr", gopoet.ImportString())
	f.Parameter("someBuffer", gopoet.ImportFromInstance(bytes.Buffer{}))
	f.ResultParameter("result", gopoet.ImportInt())
	f.ResultParameter("_", gopoet.ImportString())
	f.BlockStart("for (i := 0; i < 5; i++)")
	f.Statement("$T($S)", fmtImport, "I Make Gains")
	f.BlockEnd()
	f.Statement("$T($S)", fmtImport, "Hello Kevin")

	fileSpec := gopoet.NewFileSpec("blah")
	fileSpec.CodeBlock(f)

	structSpec := gopoet.NewStructSpec("foo")
	structSpec.StructComment("this is a comment")
	structSpec.FieldWithTag("test", gopoet.ImportString(), "test")
	structSpec.Field("blah", gopoet.ImportInt())
	structSpec.Field("buf", gopoet.ImportFromInstance(bytes.Buffer{}))

	testInitPackage := gopoet.Import{
		Package: "os",
		Alias:   "_",
	}

	interfSpec := gopoet.NewInterfaceSpec("baz")
	interfSpec.Method(*f)

	fileSpec.CodeBlock(interfSpec)

	fileSpec.InitializationPackage(testInitPackage)
	fileSpec.CodeBlock(structSpec)
	fmt.Println(fileSpec.String())
}

// myFmt := ImportDef.Get("fmt", nil)
// blankImport := ImportDef.GetInitialization("fmt")
// aliasImport := ImportDef.GetAlias("fmt", "tmf")
// aliasImport := ImportDef.GetAliasedType("fmt", "tmf", "PrintLn")
// unqualifiedFmt := ImportDef.GetUnqualified("fmt", "PrintLn")
// myFmtPrintLn := ImportDef.Get("fmt", "PrintLn")

// inlineStruct := StructSpec{}

// mainSpec := FuncSpec { Name="main"}
// mainSpec.AddStatement("$T.PrintLn($S)", ImportDef.Get("fmt"), "Hello World")
// mainSpec.AddStatement("$T.($S)", myFmtPrintLn, "Hello World")
// mainSpec.AddStatement("$T.helloStatic()", ImportDef.Get(A{}))
// mainSpec.AddStatement("a := &$T{}", ImportDef.Get(A{}))
// mainSpec.AddStatement("a.hello()")

// fileSpec := FileSpec { Package="main" }
// fileSpec.AddCodeBlock(mainSpec)
// fileSpec.String()

// intSpec := InterfaceSpec
// intSpec.addMethod("Hello")
