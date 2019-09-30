package watcher

import (
	"os"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

var (
	fieldEnvelope = false
)

func init() {
	envelope, ok := os.LookupEnv("GOGOPROTO_FIELD_ENVELOPE")
	if ok && strings.TrimSpace(envelope) == "1" {
		fieldEnvelope = true
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
