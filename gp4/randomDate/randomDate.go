package randomDate

import (
	"fmt"
	"go_play/gp3/listings3"
	"math/rand"
)

const era = "AD"

func RandomDate() {

	// Range between 2000 & 2020
	year := 2000 + rand.Intn(21)
	month := rand.Intn(12) + 1
	daysInMonth := 31
	isLeap := listings3.IsLeapYear(year)

	switch month {
	case 2:
		if isLeap {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	fmt.Println(era, year, month, daysInMonth)
}
