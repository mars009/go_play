package main

import (
	"fmt"
	"go_play/gp2/malacandra"
)

func main() {
	fmt.Printf("Speed needed to get to Malacandra in 28 days is %v km/h", malacandra.SpeedToReachMalacandraIn(28))
}
