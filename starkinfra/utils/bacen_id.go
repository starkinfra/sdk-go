package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var Chars = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", 
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2","3", "4", "5", "6", "7", "8", "9",
}

func BacenId(bankCode string) string {
	return bacenIdWithLayout(bankCode, "200601021504")
}

func bacenIdWithLayout(bankCode string, dateLayout string) string {
	var randomString string

	for i := 0; i < 11; i++ {
		random := fmt.Sprintf("%v", Chars[rand.Intn(len(Chars))])
		randomString += random
	}
	return fmt.Sprintf("%v%v%v", bankCode, time.Now().Format(dateLayout), randomString)
}

func BankCode(BankCode string) string {
	var randomString string
	randomSource := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2",
		"3", "4", "5", "6", "7", "8", "9",
	}

	for i := 1; i <= 11; i++ {
		randomString += randomSource[rand.Intn(len(randomSource))]
	}

	return fmt.Sprintf("%v%v%v",
		BankCode,
		time.Now().Format("200601021504"),
		randomString,
	)
}
