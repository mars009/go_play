package caesar

import "fmt"

// DesipherCaesar desiphers a message by removing 3 from every character within [a-]
func DesipherCaesar() {
	message := "L fdph, L vdz, L frqtxhuhg"

	for _, c := range message {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			fmt.Printf("%c", c-3)
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
