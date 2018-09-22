package piggybankint

import (
	"fmt"
	"math/rand"
)

var coins = []int{5, 10, 25}
var total = 0

func getRandomIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

// AddToPiggyBank adds nickels, dimes and pennies to the piggy bank
// until it reaches $20
func AddToPiggyBank() {
	for {
		if float32(total)/100 > 20 {
			break
		}

		total += coins[getRandomIntInRange(0, 3)]
		fmt.Printf("Total in piggy bank: $%.2f\n", float32(total)/100)
	}
}
