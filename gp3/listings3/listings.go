package listings3

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// AppleIsGreterThanBanana returns whether an apple is greater than a banana
func AppleIsGreterThanBanana() bool {
	return "apple" > "banana"
}

// FollowDirection retunrs a string containing where to go based on given command
func FollowDirection(command string) string {
	command = strings.ToLower(command)
	if command == "go east" {
		return "You head further into the mountain."
	} else if command == "go inside" {
		return "You enter a cave where you live the rest of your life"
	} else {
		return "??"
	}
}

// GetRoomDescritpion returns the description based on the given room
func GetRoomDescritpion(room string) string {

	room = strings.ToLower(room)

	if room == "cave" {
		return "Cold dark place."
	} else if room == "entrance" {
		return "Beginning of long journey."
	} else if room == "mountain" {
		return "Nature all around."
	} else {
		return "Unknown room"
	}
}

// IsLeapYear returns whether given year is leap or not
// A year is leap if is evenly divisible by 4, but not evenly divisible by 100
// or any year that is evenly divisible by 400
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// ShouldYouLeap returns a string specyfing whether you should leap or not for the given year
func ShouldYouLeap(year int) string {
	if IsLeapYear(year) {
		return "Look before you leap"
	}
	return "Keep your feet on the ground"
}

// FollowCommandSwitch example that uses a switch to set the result
func FollowCommandSwitch(room string) string {

	room = strings.ToLower(room)
	result := ""

	switch room {
	case "cave":
		result += "You find yourself in a dimly lit cave.\n"
	case "lake":
		result += "The ice seems solid enough.\n"
		fallthrough
	case "underwater":
		result += "The water is freaking cold."
	}

	return result
}

// CountDown counts backawrds from given start time
func CountDown(start int) {

	var noProblems = true

	for start > 0 {
		fmt.Println(start)
		time.Sleep(time.Second)

		if rand.Intn(100) == 0 {
			fmt.Println("rand")
			noProblems = false
			break
		}

		start--
	}

	if noProblems {
		fmt.Println("Go!")
	} else {
		fmt.Println("Boom! Oh noes!!")
	}

}
