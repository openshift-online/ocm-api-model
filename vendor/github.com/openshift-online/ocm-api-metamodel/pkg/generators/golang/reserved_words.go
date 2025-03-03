/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package golang

// AvoidReservedWord checks if the given string is one of the Go reserved words. If it is it will
// add a trailing underscore, so that the result will never be a Go reserved word.
func AvoidReservedWord(word string) string {
	_, isReserved := reservedWords[word]
	if isReserved {
		word = word + "_"
	}
	return word
}

// reservedWords is the list of Go reserved words, which should be avoided in generated names. Note
// that we use a map as a simple implementation of a set, as all what we need is to check if strings
// are reserved words.
var reservedWords = map[string]interface{}{
	"break":       nil,
	"case":        nil,
	"chan":        nil,
	"const":       nil,
	"continue":    nil,
	"default":     nil,
	"defer":       nil,
	"else":        nil,
	"fallthrough": nil,
	"for":         nil,
	"func":        nil,
	"go":          nil,
	"goto":        nil,
	"if":          nil,
	"import":      nil,
	"interface":   nil,
	"map":         nil,
	"package":     nil,
	"range":       nil,
	"return":      nil,
	"select":      nil,
	"struct":      nil,
	"switch":      nil,
	"type":        nil,
	"var":         nil,
}
