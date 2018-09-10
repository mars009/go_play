package listings4

import (
	"fmt"
	"math/rand"
)

const era = "AD"

// PrintRandomDate prints a random date
func PrintRandomDate() {

	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)

}
