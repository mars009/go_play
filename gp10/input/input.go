package input

import (
	"fmt"
)

// StringToBool converts strings to booleans
func StringToBool(value string) (result bool, err error) {

	switch value {
	case "true", "yes", "1":
		return true, nil
	case "false", "no", "0":
		return false, nil
	}

	return false, fmt.Errorf("Cannot translate value %v to boolean", value)
}
