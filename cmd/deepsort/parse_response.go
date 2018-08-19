package main

import (
	"strings"

	"github.com/savaki/jq"
)

var replaceSpace = strings.NewReplacer(" ", "_")
var replaceComma = strings.NewReplacer(",", "")
var replaceDoubleQuote = strings.NewReplacer("\"", "")

func parseResponse(rawResponse []byte) string {
	op, _ := jq.Parse(".body.predictions.[0].classes.[0].cat")
	value, _ := op.Apply(rawResponse)
	class := strings.Split(string(value), " ")
	class = class[1:len(class)]
	result := strings.Join(class, " ")
	result = replaceSpace.Replace(result)
	result = replaceComma.Replace(result)
	result = replaceDoubleQuote.Replace(result)
	return (result)
}
