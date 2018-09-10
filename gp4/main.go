package main

import (
	"go_play/gp4/listings4"
	"go_play/gp4/randomDate"
)

func main() {

	listings4.PrintRandomDate()
	for i := 0; i < 10; i++ {
		randomDate.RandomDate()
	}

}
