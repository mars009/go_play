package listings9

import "fmt"

// PrintCharacters prints the characters of the given word
func PrintCharacters(word string) {
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c", word[i])
	}

	fmt.Println()
}
