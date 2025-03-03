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

package names

import (
	"strings"
)

// Name represents a name formed of multiple words. It is inteded to simplify the use of different
// strategies for representing names as strings, like using different separators or using camel
// case. The workds that form the name are stored separated, so there is no need to parse the name
// each time that the workds are needed.
type Name struct {
	text  string
	words []*Word
}

// NewName creates a new name with the given words.
func NewName(words ...*Word) *Name {
	name := new(Name)
	name.words = make([]*Word, len(words))
	copy(name.words, words)
	return name
}

// Text returns the text that was originally used to create this name.
func (n *Name) Text() string {
	return n.text
}

// SetText sets the text that was originally used to create this name.
func (n *Name) SetText(text string) {
	n.text = text
}

// Words returns a slice containing the words of the name. The returned slice is a copy of the
// internal representation, to it is safe to modify after calling this function.
func (n *Name) Words() []*Word {
	words := make([]*Word, len(n.words))
	copy(words, n.words)
	return words
}

// Append creates a new name that has the words of this name, plus the given additional word added
// at the end.
func (n *Name) Append(words ...*Word) *Name {
	words = append(n.words, words...)
	return NewName(words...)
}

// Insert creates a new name that has the given words, plus the words of this name added at the end.
func (n *Name) Insert(words ...*Word) *Name {
	words = append(words, n.words...)
	return NewName(words...)
}

// Equals check if this name is equal to the given name.
func (n *Name) Equals(name *Name) bool {
	if n == nil && name == nil {
		return true
	}
	if n == nil || name == nil {
		return false
	}
	if len(n.words) != len(name.words) {
		return false
	}
	for i := 0; i < len(n.words); i++ {
		if !n.words[i].Equals(name.words[i]) {
			return false
		}
	}
	return true
}

// LowerJoined converts all the words of the name to lower case and joins them using the given
// separator.
func (n *Name) LowerJoined(separator string) string {
	chunks := make([]string, len(n.words))
	for i, word := range n.words {
		chunks[i] = strings.ToLower(word.String())
	}
	return strings.Join(chunks, separator)
}

// UpperJoined converts all the words of the name to upper case and joins them using the given
// separator.
func (n *Name) UpperJoined(separator string) string {
	chunks := make([]string, len(n.words))
	for i, word := range n.words {
		chunks[i] = strings.ToUpper(word.String())
	}
	return strings.Join(chunks, separator)
}

// CapitalizedJoined capitalizes all the words of the name to lower case and joins them using the
// given separator.
func (n *Name) CapitalizedJoined(separator string) string {
	chunks := make([]string, len(n.words))
	for i, word := range n.words {
		chunks[i] = word.Capitalize()
	}
	return strings.Join(chunks, separator)
}

// String generates a string representing this name, consisting on the list of words of the name
// separated by an underscore.
func (n *Name) String() string {
	texts := make([]string, len(n.words))
	for i, word := range n.words {
		texts[i] = word.String()
	}
	return strings.Join(texts, "_")
}

// Snake converts this name to snake case.
func (n *Name) Snake() string {
	return n.LowerJoined("_")
}

// Camel conerts this name to camel case.
func (n *Name) Camel() string {
	return n.CapitalizedJoined("")
}

// Compare compares two names lexicographically.
func Compare(a, b *Name) int {
	var result int
	aLen := len(a.words)
	bLen := len(b.words)
	minLen := aLen
	if bLen < minLen {
		minLen = bLen
	}
	for i := 0; i < minLen; i++ {
		aText := a.words[i].String()
		bText := b.words[i].String()
		result = strings.Compare(aText, bText)
		if result != 0 {
			break
		}
	}
	if result == 0 {
		if aLen < bLen {
			result = -1
		} else if aLen > bLen {
			result = +1
		}
	}
	return result
}

// Cat returns a new name created concatenating the given names.
func Cat(names ...*Name) *Name {
	result := new(Name)
	length := 0
	for _, name := range names {
		length = length + len(name.words)
	}
	result.words = make([]*Word, length)
	i := 0
	for _, name := range names {
		for _, word := range name.words {
			result.words[i] = word
			i++
		}
	}
	return result
}

// Names is an slice of names, intended to simplify sorting.
type Names []*Name

// Len implements the Len method of sort.Interface, so that slices of names can be easily sorted.
func (n Names) Len() int {
	return len(n)
}

// Swap implements the Swap method of sort.Interface, so that slices of names can be easily sorted.
func (n Names) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// Less implements the Less method of sort.Interface, so that slices of names can be easily sorted.
func (n Names) Less(i, j int) bool {
	return Compare(n[i], n[j]) == -1
}
