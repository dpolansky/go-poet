package poet

import (
	"fmt"
)

// MethodSpec represents a method, with a receiver name and type.
type MethodSpec struct {
	FuncSpec
	ReceiverName string
	Receiver     TypeReference
}

var _ CodeBlock = (*MethodSpec)(nil)

// NewMethodSpec creates a new method with the given method name, receiverName, and receiver type.
func NewMethodSpec(name, receiverName string, receiver TypeReference) *MethodSpec {
	return &MethodSpec{
		FuncSpec: FuncSpec{
			Name: name,
		},
		ReceiverName: receiverName,
		Receiver:     receiver,
	}
}

func (m *MethodSpec) String() string {
	writer := newCodeWriter()

	signature, args := m.Signature()
	format := fmt.Sprintf("func ($L $T) %s {", signature)
	args = append([]interface{}{m.ReceiverName, m.Receiver}, args...)
	writer.WriteStatement(newStatement(0, 1, format, args...))

	for _, st := range m.Statements {
		writer.WriteStatement(st)
	}

	writer.WriteStatement(newStatement(-1, 0, "}"))

	return writer.String()
}
