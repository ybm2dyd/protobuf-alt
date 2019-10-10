package watcher

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

var (
	fieldEnvelope = false
	vizLevel      = false
)

func init() {
	envelope, ok := os.LookupEnv("GOGOPROTO_FIELD_ENVELOPE")
	if ok && strings.TrimSpace(envelope) == "1" {
		fieldEnvelope = true
	}
	level, ok := os.LookupEnv("GOGOPROTO_VIZ_LEVEL")
	if ok && strings.TrimSpace(level) == "1" {
		vizLevel = true
	}
}

func PrintFieldEnvelope(g *generator.Generator, messageName, fieldName string) func() {
	if !fieldEnvelope {
		return func() {}
	}

	g.P(`{ // FE - `, messageName, ":", fieldName)
	g.In()
	return func() {
		g.P(`}`)
		g.Out()
	}
}

func vizLevelFuncParam() string {
	if !vizLevel {
		return ""
	} else {
		return "vizLevel int8"
	}
}

func vizLevelFuncArgument() string {
	if !vizLevel {
		return ""
	} else {
		return "vizLevel"
	}
}

func PrintFuncSignatureOfSize(g *generator.Generator, ccTypeName, sizeName string) {
	if !vizLevel {
		g.P(`func (m *`, ccTypeName, `) `, sizeName, `() (n int) {`)
	} else {
		g.P(`func (m *`, ccTypeName, `) `, sizeName, `() (n int) {`)
		g.In()
		g.P(`return m.`, sizeName, `WithVizLevel(0)`)
		g.Out()
		g.P("}")
		g.P(`func (m *`, ccTypeName, `) `, sizeName, `WithVizLevel(`, vizLevelFuncParam(), `) (n int) {`)
	}
}

func FuncInvocationOfSize(sizeName string) string {
	if !vizLevel {
		return sizeName + "()"
	} else {
		return fmt.Sprintf("%sWithVizLevel(%s)", sizeName, vizLevelFuncArgument())
	}
}

func PrintFuncSignatureOfMarshal(g *generator.Generator, ccTypeName string) {
	if !vizLevel {
		g.P(`func (m *`, ccTypeName, `) Marshal() (dAtA []byte, err error) {`)
	} else {
		g.P(`func (m *`, ccTypeName, `) Marshal() (dAtA []byte, err error) {`)
		g.In()
		g.P(`return m.MarshalWithVizLevel(0)`)
		g.Out()
		g.P("}")
		g.P(`func (m *`, ccTypeName, `) MarshalWithVizLevel(`, vizLevelFuncParam(), `) (dAtA []byte, err error) {`)
	}
}

func PrintFuncSignatureOfMarshalTo(g *generator.Generator, ccTypeName, methodName string) {
	if !vizLevel {
		g.P(`func (m *`, ccTypeName, `) `, methodName, `(dAtA []byte) (int, error) {`)
	} else {
		g.P(`func (m *`, ccTypeName, `) `, methodName, `(dAtA []byte) (int, error) {`)
		g.In()
		g.P(`return m.`, methodName, `WithVizLevel(dAtA, 0)`)
		g.Out()
		g.P("}")
		g.P(`func (m *`, ccTypeName, `) `, methodName, `WithVizLevel(dAtA []byte, `, vizLevelFuncParam(), `) (int, error) {`)
	}
}

func FuncInvocationOfMarshalTo(methodName, dAtA string) string {
	if !vizLevel {
		return fmt.Sprintf("%s(%s)", methodName, dAtA)
	} else {
		return fmt.Sprintf("%sWithVizLevel(%s, %s)", methodName, dAtA, vizLevelFuncArgument())
	}
}
