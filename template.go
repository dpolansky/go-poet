package gopoet

import (
	"bytes"
	"fmt"
	"strings"
)

const TEMPLATING_CHAR = '$'

// $T.Println($S + "$$asdf")
func Template(format string, args ...interface{}) string {
	var buffer bytes.Buffer

	currentArg := 0

	for i := 0; i < len(format); i++ {
		if format[i] == TEMPLATING_CHAR && i+1 < len(format) {
			if currentArg+1 > len(args) {
				panic(fmt.Sprintf("Not enough arguments for format string ('%s'), got %d", format, len(args)))
			}

			switch format[i+1] {
			case 'T':
				buffer.WriteString(getQualifiedNameFromArg(args[currentArg]))
				break
			case 'S':
				buffer.WriteString("\"")
				buffer.WriteString(strings.Replace(args[currentArg].(string), "\"", "\\\"", -1))
				buffer.WriteString("\"")
				break
			default:
				panic(fmt.Sprintf("Unrecognized templating character in format string ('%s')", format))
			}

			currentArg++
			i++
		} else {
			buffer.WriteByte(format[i])
		}
	}

	return buffer.String()
}

func getQualifiedNameFromArg(obj interface{}) (result string) {
	importSpec, ok := obj.(ImportSpec)
	if !ok {
		panic(fmt.Sprintf("$T must take an instance of ImportSpec, got %+v", obj))
	}

	if importSpec.NeedsQualifier() {
		if importSpec.GetPackageAlias() != "" {
			result += importSpec.GetPackageAlias()
		} else {
			result += importSpec.GetPackage()
		}
	}

	if importSpec.GetName() != "" {
		result += "."
		result += importSpec.GetName()
	}

	return result
}
