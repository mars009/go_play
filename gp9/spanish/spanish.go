package spanish

import "fmt"

func CipherSpanish() {
	message := "Hola Estación Espacial Internacional"

	for _, char := range message {
		fmt.Printf("%c", char+3)
	}

}

func DecipherSpanish() {
	message := "Krod#Hvwdflr̄q#Hvsdfldo#Lqwhuqdflrqdo"

	for _, char := range message {
		fmt.Printf("%c", char-3)
	}
}
