package poet

import "bytes"

type MethodSpec struct {
	// CodeBlock

	FuncSpec
	ReceiverName string
	Receiver     TypeReference
}

var _ CodeBlock = (*MethodSpec)(nil)

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

	writer.WriteStatement(m.createSignature())

	for _, st := range m.Statements {
		writer.WriteStatement(st)
	}

	writer.WriteStatement(statement{
		BeforeIndent: -1,
		Format:       "}",
	})

	return writer.String()
}

func (m *MethodSpec) createSignature() statement {
	formatStr := bytes.Buffer{}
	signature, args := m.Signature()

	formatStr.WriteString("func ")
	formatStr.WriteString("(")
	formatStr.WriteString(m.ReceiverName)
	formatStr.WriteString(" ")
	formatStr.WriteString(m.Receiver.GetName())
	formatStr.WriteString(") ")
	formatStr.WriteString(signature)
	formatStr.WriteString(" {")

	return statement{
		AfterIndent: 1,
		Format:      formatStr.String(),
		Arguments:   args,
	}
}
