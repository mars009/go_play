package listings8

import (
	"fmt"
	"math/big"
)

var lightSpeedBig = big.NewInt(299792)
var secondsPerDayBig = big.NewInt(86400)

// AndromedaGalaxy prints out the distance in days to the star @ light speed
func AndromedaGalaxy() {
	/*
	 * For extremely big numbers you need to first create the
	 * big.Int by using 'new()' and then use the 'SetString()'
	 * method of big.Int
	 */
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeedBig)

	days := new(big.Int)
	days.Div(seconds, secondsPerDayBig)

	fmt.Println("That is", days, "days of travel at light speed")
}
