package malacandra

const distance = 56000000
const hoursPerDay = 24
const speedOfLight = 299792

// SpeedToReachMalacandraIn calculates the speed needed to reach Malacandra in the
// given amount of days
func SpeedToReachMalacandraIn(days int) int {
	return distance / (days * hoursPerDay)
}
