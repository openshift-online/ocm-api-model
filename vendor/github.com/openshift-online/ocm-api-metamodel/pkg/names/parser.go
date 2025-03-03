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
	"bytes"
	"strings"
	"unicode"
)

// ParseUsingSeparator separates the given text into words, using the given separator character, and
// creates a new name containing those words. For example, to convert the text 'my_favorite_fruit'
// into a name the method can be used as follows:
//
//	name := names.ParseUsingSeparator("my_favorite_fruit", "_")
func ParseUsingSeparator(text string, separator string) *Name {
	chunks := strings.Split(text, separator)
	words := make([]*Word, len(chunks))
	for i, chunk := range chunks {
		words[i] = NewWord(chunk)
	}
	name := NewName(words...)
	name.SetText(text)
	return name
}

// ParseUsingCase separates the given text into words, using the case transitions as separators, and
// creates a new name containing those words.
func ParseUsingCase(text string) *Name {
	// Convert the text to an array of runes so that we can easily access the previous, current
	// and next runes:
	var runes []rune
	for _, r := range text {
		runes = append(runes, r)
	}

	// Split the text at the points where there are case transitions:
	var chunks []string
	size := len(runes)
	if size > 1 {
		buffer := new(bytes.Buffer)
		for i := 0; i < size-1; i++ {
			current := runes[i]
			next := runes[i+1]
			buffer.WriteRune(current)
			if unicode.IsLower(current) != unicode.IsLower(next) {
				chunk := buffer.String()
				chunks = append(chunks, chunk)
				buffer.Reset()
			}
		}
		buffer.WriteRune(runes[size-1])
		chunk := buffer.String()
		chunks = append(chunks, chunk)
	} else {
		chunks = []string{text}
	}

	var words []*Word
	size = len(chunks)
	if size > 1 {
		// Process the chunks:
		i := 0
		for i < size-1 {
			current := chunks[i]
			next := chunks[i+1]

			// A single upper case character followed by a lower case chunk is one
			// single word, like 'Name`:
			if len(current) == 1 && isUpper(current) && isLower(next) {
				word := NewWord(strings.ToLower(current) + next)
				words = append(words, word)
				i = i + 2
				continue
			}

			// A chunk with two or more upper case characters followed by a single `s`
			// is the plural form of an initialism, like `CPUs` or `IDs`:
			if len(current) >= 2 && isUpper(current) && next == "s" {
				word := NewInitialism(current + "s")
				words = append(words, word)
				i = i + 2
				continue
			}

			// An upper case chunk followed by a lower case chunk is two words and the
			// last character of the first chunk is the first character of the second
			// word, like in `CPUList` which corresponds to `CPU` and `List`:
			if isUpper(current) && isLower(next) {
				first := NewInitialism(current[0 : len(current)-1])
				second := NewWord(strings.ToLower(current[len(current)-1:]) + next)
				words = append(words, first)
				words = append(words, second)
				i = i + 2
				continue
			}

			// Anything else is a word by itself:
			word := parseWord(text)
			words = append(words, word)
			i++
		}

		// If there is still a chunk to process then it is a word by itself:
		if i < size {
			word := parseWord(chunks[i])
			words = append(words, word)
		}
	} else {
		word := parseWord(text)
		words = []*Word{word}
	}

	// Create the name from the stored words:
	name := NewName(words...)
	name.SetText(text)
	return name
}

// parseWord converts the given text into a Word object assuming that it is a single word.
func parseWord(text string) *Word {
	if len(text) >= 2 && isUpper(text) {
		return NewInitialism(text)
	}
	return NewWord(strings.ToLower(text))
}
