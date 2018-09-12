package listings6

import "fmt"

// PrintFloatFormatter examples of printing floats with formatter
func PrintFloatFormatter() {
	third := 1.0 / 3

	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	// '4' represents the width (min number of characters to display)
	// '2' represents the precision (specifies number of digits after decimal point)
	fmt.Printf("%4.2f\n", third)
	// To left pad with zeroes, prefix the width with a '0'
	fmt.Printf("%05.2f\n", third)
}

// ElevenDimes prints the result of adding 11 dimes
func ElevenDimes() {

	total := 0.0

	for i := 0; i < 11; i++ {
		total += .10
	}

	fmt.Printf("Eleven Dimes are %v\n", total)
}
