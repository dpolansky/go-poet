# go-poet

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/dpolansky/go-poet/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/dpolansky/go-poet?status.svg)](http://godoc.org/github.com/dpolansky/go-poet/poet)
[![Build Status](https://travis-ci.org/dpolansky/go-poet.svg?branch=master)](https://travis-ci.org/dpolansky/go-poet)
[![Go Report Card](https://goreportcard.com/badge/github.com/dpolansky/go-poet)](https://goreportcard.com/report/github.com/dpolansky/go-poet)
[![Codecov Coverage Report](https://codecov.io/github/dpolansky/go-poet/coverage.svg?branch=master)](https://codecov.io/gh/dpolansky/go-poet)

go-poet is a Go package for generating Go code, inspired by [javapoet](https://github.com/square/javapoet).

Typically, code generation uses text templating which can be error prone and hard to maintain. This project aims to fix these issues by giving you an API to generate Go constructs in a typesafe way, while handling imports and formatting for you.

  - [Installation](#installation)
  - [Example](#example)
  - [Getting Started](#getting-started)
  - [Code Blocks](#code-blocks)
    - [Functions](#functions)
    - [Interfaces](#interfaces)
    - [Structs](#structs)
    - [Globals](#globals)
  - [Type References](#type-references)
    - [Package Aliases](#package-aliases)
    - [Custom Names](#custom-names)
    - [Unqualified Types](#unqualified-types)
  - [Templating](#templating)

## Installation
```
$ go get github.com/dpolansky/go-poet/poet
```

## Example
Here's a Hello World Go file
```go
package main

import (
        "fmt"
)

func main() {
        fmt.Println("Hello world!")
}
```
and the go-poet code to generate it
```go
main := poet.NewFuncSpec("main").
		Statement("$T($S)", poet.TypeReferenceFromInstance(fmt.Println), "Hello world!")

file := poet.NewFileSpec("main").
		CodeBlock(main)
```

## Getting Started
To get started, import `"github.com/dpolansky/go-poet/poet"`

The end goal of go-poet is to create a compilable file. To construct a new file with package `main`:

```go
file := poet.NewFileSpec("main")
```  

Files contain CodeBlocks which can be global variables, functions, structs, interfaces. go-poet will handle any imports for you via TypeReferences.
The types that you create or reference can be used in code via Templates.

## Code Blocks
### Functions
Functions can be attached to a File:
```go
add := poet.NewFuncSpec("add").
	Parameter("a", poet.Int).
	Parameter("b", poet.Int).
	ResultParameter("", poet.Int).
    Statement("return a + b")
    
file.CodeBlock(add)
```
This will produce the function:
```go
func add(a int, b int) int {
        return a + b
}
```

To add control flow statements, use BlockStart and BlockEnd. Indentation will be handled for you.
```go
loop := poet.NewFuncSpec("loop").
	BlockStart("for i := 0; i < 5; i++").
	Statement("$T($L)", poet.TypeReferenceFromInstance(fmt.Println), "i").
	BlockEnd()
```
produces
```go
func loop() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```
### Interfaces
```go
inter := poet.NewInterfaceSpec("AddWriter").
	EmbedInterface(poet.TypeReferenceFromInstance((*io.Writer)(nil))).
	Method(add)
```
produces
```go
type AddWriter interface {
        io.Writer
        add(a int, b int) int
}
```

### Structs
Structs can have fields, directly attached methods, and a comment.
```go
foo := poet.NewStructSpec("foo").
	Field("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{})).
	AttachFunction("f", add)
```
produces
```go
type foo struct {
        buf *bytes.Buffer
}

func (f *foo) add(a int, b int) int {
        return a + b
}
```

### Globals
Global variables and constants can be added directly to a file, either standalone or in groups.
```go
file.GlobalVariable("a", poet.String, "$S", "hello")
file.GlobalConstant("b", poet.Int, "$L", 1)
```
```go
var a string = "hello"

const b int = 1
```

or if you want to group them
```go
file.VariableGrouping().
		Variable("a", poet.Int, "$L", 7).
		Variable("b", poet.Int, "$L", 2).
		Constant("c", poet.Int, "$L", 3).
		Constant("d", poet.Int, "$L", 43)
```
```go
const (
        c int = 3
        d int = 43
)

var (
        a int = 7
        b int = 2
)

```
## Type References
To ensure type safe code and handle a generated file's imports, use TypeReferences.

For example, to use `bytes.Buffer` as a parameter to a function
```go
poet.NewFuncSpec("foo").Parameter("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{}))
```
produces
```go
func foo(buf *bytes.Buffer) {
}
```
The `poet.TypeReferenceFromInstance` function takes an instance of a variable or a function and uses reflection to determine it's type and package.

### Package Aliases
To use an aliased package's name from a TypeReference, use `poet.TypeReferenceFromInstanceWithAlias`.
```go
TypeReferenceFromInstanceWithAlias(&bytes.Buffer{}, "myAlias")
```
produces a TypeReference with type
```go
*myAlias.Buffer
```

### Custom Names
For type aliases, you may want to reference the aliased name instead of the underlying type.

To do this, use `poet.TypeReferenceFromInstanceWithCustomName`
```go
poet.TypeReferenceFromInstanceWithCustomName(uint8(0), "byte")
```
### Unqualified Types
If you want a type to be unqualified, create a type alias with the prefix `_unqualified` followed by the name
```go
type _unqualifiedBuffer bytes.Buffer
typeRef := TypeReferenceFromInstance(_unqualifiedBuffer{})
```
produces the type `Buffer`

## Templating
Format strings are used to construct statements in functions or values for variables.

We currently support these format specifiers:
* **Strings** `$S` Takes a `string` as input, surrounding it with quotes and escaping quotes within the input
* **Literals** `$L` Takes any value as input, and uses Go's `Sprintf` `%v` formatting to write the input
* **Types** `$T` Takes a TypeReference as input, and writes its qualified/aliased name

## Authors
[Dave Polansky](http://github.com/dpolansky)

[Kevin Deenanauth](http://github.com/kdeenanauth)