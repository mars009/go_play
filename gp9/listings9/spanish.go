package listings9

import (
	"fmt"
	"unicode/utf8"
)

// SpanishRunes prints out information about the spanish phrase
func SpanishRunes(phrase string) {

	fmt.Println(len(phrase), "bytes")
	fmt.Println(utf8.RuneCountInString(phrase), "runes")

	c, size := utf8.DecodeRuneInString(phrase)
	fmt.Printf("First rune: %c %v bytes", c, size)
}
