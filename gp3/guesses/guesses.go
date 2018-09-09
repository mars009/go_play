package guesses

import (
	"fmt"
	"math/rand"
)

// GuessNumber will try to guess your number, printing out the guess
// or your actual number if it guessed correctly
func GuessNumber(number int) {
	var guess = rand.Intn(100) + 1

	for {
		if guess == number {
			fmt.Printf("Your number is %v\n", guess)
			break
		} else {
			fmt.Printf("Guessed %v\n", guess)
			guess = rand.Intn(100) + 1
		}
	}

}
