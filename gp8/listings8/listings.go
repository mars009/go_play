package listings8

import "fmt"

const lightSpeed = 299792
const secondsPerDay = 86400

const distance int64 = 41.3e12

// AlphaCentauri prints out the distance in days to the star @ light speed
func AlphaCentauri() {

	fmt.Println("Alpha Centauri is", distance, "km away")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed")

}
