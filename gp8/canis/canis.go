package canis

import "fmt"

const distanceToCanis = 236000000000000000
const lightSpeedConst = 299792
const secsPerDayConst = 86400

// DistanceToCanis prints distance of this galaxy in light years
func DistanceToCanis() {

	fmt.Println("Canis Galaxy is", distanceToCanis, "km away")

	lightDays := distanceToCanis / lightSpeedConst / secsPerDayConst
	lightYears := lightDays / 365

	fmt.Println("Canis Galaxy is", lightYears, "light years away")

}
