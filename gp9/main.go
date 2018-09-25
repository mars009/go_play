package main

import (
	"fmt"
	"go_play/gp9/caesar"
	"go_play/gp9/spanish"
)

func main() {

	// listings9.PrintCharacters("shalom")

	// fmt.Printf("%c\n", 'g'-'a'+'A')

	// listings9.Decipher("uv vagreangvbany fcnpr fgngvba")

	// listings9.SpanishRunes("Que onda")

	caesar.DesipherCaesar()

	spanish.CipherSpanish()
	fmt.Println()
	spanish.DecipherSpanish()

}
