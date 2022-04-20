package parser

import (
	"fmt"
	"strings"
)

//nolint:deadcode,unused
var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

//nolint:deadcode,unused
func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

//nolint:deadcode,unused
func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

//nolint:deadcode,unused
func incIdent() { traceLevel = traceLevel + 1 }
//nolint:deadcode,unused
func decIdent() { traceLevel = traceLevel - 1 }

//nolint:deadcode,unused
func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

//nolint:deadcode,unused
func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
