package piggybank

import (
	"fmt"
	"math"
	"math/rand"
)

var coins = []float64{.05, .10, .25}
var total = 0.0

func getRandomNumberInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getCoin() float64 {
	return coins[getRandomNumberInRange(0, 3)]
}

// AddToPiggyBank adds nickels, dimes and quarters to the piggy bank
// until the total amount reaches $20
func AddToPiggyBank() {
	for {
		if math.Abs(total) > 20 {
			break
		}
		total += getCoin()
		fmt.Printf("Total in bank: $%.2f\n", total)
	}
}
