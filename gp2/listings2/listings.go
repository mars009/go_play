package listings

import (
	"fmt"
	"math/rand"
)

const marsDays = 687
const weightInMars = .3783
const lightSpeed = 299792 //km/s
const hoursPerDay = 23

// PrintAgeAndWeightInMars prints the weight and age of a person based on given params
func PrintAgeAndWeightInMars(age int, weight float32) {
	fmt.Printf("%-23v %4v\n", "Your weight in Mars is", weight*weightInMars)
	fmt.Printf("%-23v %4v\n", "Your age in Mars is", age*365/67)
}

// TimeToGetToMars returns the time to get to Mars in @lightspeed based on the given distance
func TimeToGetToMars(speed int, distance int) int {
	//fmt.Printf("Time to get to Mars @ distance %v is %v", distance, distance/lightSpeed)

	if speed == -1 {
		speed = lightSpeed
	}

	return distance / speed
}

// DaysToArriveToMars returns the numbers of days to arrive to mars based on given speed and distance
// Use -1 to use speed of light  (299792 km/s)
func DaysToArriveToMars(speed int, distance int) int {
	return TimeToGetToMars(speed, distance) / hoursPerDay
}

// GetRandomDistance returns a random distance between 56 million km and 401 million km
func GetRandomDistance() int {
	// Adding 56 million makes sure our values have the desired range
	return rand.Intn(345000001) + 56000000
}
