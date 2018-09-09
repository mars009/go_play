package main

import (
	"fmt"
	"go_play/gp3/guesses"
	"go_play/gp3/listings3"
)

func main() {
	fmt.Println(listings3.AppleIsGreterThanBanana())

	fmt.Println(listings3.FollowDirection("go north"))

	fmt.Println(listings3.GetRoomDescritpion("Cave"))

	fmt.Println(listings3.ShouldYouLeap(2100))

	fmt.Println(listings3.FollowCommandSwitch("lake"))

	listings3.CountDown(5)

	guesses.GuessNumber(54)
}
