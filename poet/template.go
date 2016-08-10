package poet

import (
	"bytes"
	"fmt"
	"strings"
)

const templatingChar = '$'

func template(format string, args ...interface{}) string {
	var buffer bytes.Buffer

	currentArg := 0

	for i := 0; i < len(format); i++ {
		if format[i] == templatingChar && i+1 < len(format) {
			if currentArg+1 > len(args) {
				panic(fmt.Sprintf("Not enough arguments for format string ('%s'), got %d", format, len(args)))
			}

			switch format[i+1] {
			case 'L':
				buffer.WriteString(fmt.Sprintf("%v", args[currentArg]))
				break
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

func getQualifiedNameFromArg(obj interface{}) string {
	typeRef, ok := obj.(TypeReference)
	if !ok {
		panic(fmt.Sprintf("$T must implement TypeReference, got type=%T %#v", obj, obj))
	}

	return typeRef.GetName()
}
