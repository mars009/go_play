package tickettomars

import (
	"fmt"
	"math/rand"
)

var spacelines = []string{"SpaceX", "Space Adventures", "Virgin Galactic"}

const secondsPerDay = 86400
const distanceToMars = 62100000

func getRandomAirline() string {
	return spacelines[getNumberInRange(0, 3)]
}

func getNumberInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getSpeedPerHour() int {
	return getNumberInRange(16, 30)
}

func getTripType() string {
	num := getNumberInRange(0, 2)
	if num > 0 {
		return "Round-trip"
	}

	return "One way"
}

func getPrice(speed int) int {

	price := 0

	switch {
	case speed < 20:
		price = getNumberInRange(36, 40)
	case speed < 25:
		price = getNumberInRange(41, 45)
	case speed >= 26:
		price = getNumberInRange(46, 50)
	}

	return price
}

func daysToMars(speed int) int {
	return (distanceToMars / speed) / secondsPerDay
}

// PrintItinerary prints October 13, 2020's itinerary
func PrintItinerary() {

	fmt.Printf("%-16v %v %-11v %v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("======================================\n")

	for i := 0; i < 10; i++ {
		speed := getSpeedPerHour()
		fmt.Printf("%-16v %4v %-11v $%4v\n", getRandomAirline(), daysToMars(speed), getTripType(), getPrice(speed))
	}

}
