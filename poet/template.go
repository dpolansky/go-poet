package poet

import (
	"bytes"
	"fmt"
)

const templatingChar = '$'

// template write go code based on a format string with arguments.
//
// $L replaces with the literal value of the argument (%v).
// $S replaces with the quoted string value of the argument (%q).
// $T argument must be a TypeReference; it replaces with the TypeRef's GetName().
func template(format string, args ...interface{}) string {
	var buffer bytes.Buffer

	currentArg := 0

	for i := 0; i < len(format); i++ {
		if format[i] == templatingChar && i+1 < len(format) {
			if currentArg+1 > len(args) {
				panic(fmt.Sprintf("Not enough arguments for format string ('%s'), got %d", format, len(args)))
			}

			a := args[currentArg]
			switch format[i+1] {
			case 'L':
				buffer.WriteString(fmt.Sprintf("%v", a))
				break
			case 'S':
				buffer.WriteString(fmt.Sprintf("%q", fmt.Sprintf("%v", a)))
				break
			case 'T':
				buffer.WriteString(getQualifiedNameFromArg(a))
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
