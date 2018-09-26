package main

import (
	"fmt"
	"go_play/gp10/input"
)

func main() {
	fmt.Println(input.StringToBool("no"))
	value, err := input.StringToBool("nein")

	if err == nil {
		fmt.Printf("%T\n", value)
	} else {
		fmt.Println(err)
	}

}
