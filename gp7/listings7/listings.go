package listings7

import "fmt"

// WhichTypes prints the types of different primitives
func WhichTypes() {
	fmt.Printf("Type %T for 'hello', %T for 7, %T for 1.3, %T for true\n", "hello", 7, 1.3, true)
}

// ShowBitsOfInt displaus the bits of an integer value
func ShowBitsOfInt(num int) {
	// Instead of repeating the variable twice, you can tell Printf to use
	// the first argument [1] for the second format verb
	fmt.Printf("Bits of %v are : %08[1]b\n", num)
}
