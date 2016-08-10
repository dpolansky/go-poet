# go-poet

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/dpolansky/go-poet/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/dpolansky/go-poet?status.svg)](http://godoc.org/github.com/dpolansky/go-poet/poet)
[![Build Status](https://travis-ci.org/dpolansky/go-poet.svg?branch=master)](https://travis-ci.org/dpolansky/go-poet)
[![Go Report Card](https://goreportcard.com/badge/github.com/dpolansky/go-poet)](https://goreportcard.com/report/github.com/dpolansky/go-poet)
[![Codecov Coverage Report](https://codecov.io/github/dpolansky/go-poet/coverage.svg?branch=master)](https://codecov.io/gh/dpolansky/go-poet)

go-poet is a Go package for generating Go code, inspired by [javapoet](https://github.com/square/javapoet).

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
```
package main

import (
        "fmt"
)

func main() {
        fmt.Println("Hello world!")
}
```
and the go-poet code to generate it
```
main := poet.NewFuncSpec("main").
		Statement("$T($S)", poet.TypeReferenceFromInstance(fmt.Println), "Hello world!")

file := poet.NewFileSpec("main").
		CodeBlock(main)
```

## Getting Started
To get started, import `"github.com/dpolansky/go-poet/poet"`

The end goal of go-poet is to create a compilable file. To construct a new file with package `main`:

```
file := poet.NewFileSpec("main")
```  

Files contain CodeBlocks which can be global variables, functions, structs, interfaces. Go-poet will handle any imports for you via TypeReferences.
The types that you create or reference can be used in code via Templates.

## Code Blocks
### Functions
Functions can be attached to a File:
```
add := poet.NewFuncSpec("add").
	Parameter("a", poet.TypeReferenceFromInstance(int(0))).
	Parameter("b", poet.TypeReferenceFromInstance(int(0))).
	ResultParameter("", poet.TypeReferenceFromInstance(int(0))).
        Statement("return a + b")
    
file.CodeBlock(add)
```

This will produce the function:
```
func add(a int, b int) int {
        return a + b
}
```
### Interfaces
```
inter := poet.NewInterfaceSpec("AddWriter").
	EmbedInterface(poet.TypeReferenceFromInstance((*io.Writer)(nil))).
	Method(add)
```
produces
```
type AddWriter interface {
        io.Writer
        add(a int, b int) int
}
```

### Structs
Structs can have fields, directly attached methods, and a comment.
```
foo := poet.NewStructSpec("foo").
	Field("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{})).
	AttachFunction("f", add)
```
produces
```
type foo struct {
        buf *bytes.Buffer
}

func (f *foo) add(a int, b int) int {
        return a + b
}
```

### Globals
Global variables and constants can be added directly to a file, either standalone or in groups.
```
file.GlobalVariable("a", poet.TypeReferenceFromInstance(""), "$S", "hello")
file.GlobalConstant("b", poet.TypeReferenceFromInstance(0), "$L", 1)
```
```
var a string = "hello"

const b int = 1
```

or if you want to group them
```
file.VariableGrouping().
		Variable("a", poet.TypeReferenceFromInstance(0), "$L", 7).
		Variable("b", poet.TypeReferenceFromInstance(0), "$L", 2).
		Constant("c", poet.TypeReferenceFromInstance(0), "$L", 3).
		Constant("d", poet.TypeReferenceFromInstance(0), "$L", 43)
```
```
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
```
poet.NewFuncSpec("foo").Parameter("buf", poet.TypeReferenceFromInstance(&bytes.Buffer{}))
```
produces
```
func foo(buf *bytes.Buffer) {
}
```
The `poet.TypeReferenceFromInstance` function takes an instance of a variable or a function and uses reflection to determine it's type and package.

### Package Aliases
To use an aliased package's name from a TypeReference, use `poet.TypeReferenceFromInstanceWithAlias`.
```
TypeReferenceFromInstanceWithAlias(&bytes.Buffer{}, "myAlias")
```
produces a TypeReference with type
```
*myAlias.Buffer
```

### Custom Names
For type aliases, such as `byte` which aliases `uint8`, you may want to reference the aliased name instead of the underlying type.

To do this, use `poet.TypeReferenceFromInstanceWithCustomName`
```
poet.TypeReferenceFromInstanceWithCustomName(byte('A'), "byte")
```
### Unqualified Types
If you want a type to be unqualified, create a type alias with the prefix `_unqualified` followed by the name
```
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