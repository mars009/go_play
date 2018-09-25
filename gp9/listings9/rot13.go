package listings9

import "fmt"

// Decipher decodes the given text
func Decipher(text string) {
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13

			// We have to subtract 26 since we have to account for
			// wrapping around
			if c > 'z' {
				c = c - 26
			}
		}

		fmt.Printf("%c", c)
	}
}
