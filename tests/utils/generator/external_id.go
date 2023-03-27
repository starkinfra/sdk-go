package tax_id_generator

import (
	"fmt"
	"math/rand"
	"time"
)

func ExternalId() string {
	var randomString string

	for i := 0; i < 11; i++ {
		random := fmt.Sprintf("%v", Chars[rand.Intn(len(Chars))])
		randomString += random
	}
	return fmt.Sprintf("%v%v", time.Now().Format("20060102150405.999999999Z07001504"), randomString)
}
