package date

import (
	"math/rand"
	"time"
)

func FutureDate(days int) string {

	duration := 24 * time.Duration(days) * time.Hour

	future := time.Now().Add(duration)
	date := future.Format("2006-01-02")
	return date
}

func RandomFutureDate(days int) string {
	random := rand.Intn(days - 1)
	return FutureDate(random)
}

func FutureDateTime(days int) string {

	duration := 24 * time.Duration(days) * time.Hour

	future := time.Now().Add(duration)
	date := future.Format(time.RFC3339Nano)
	return date
}

func PastDate(days int) string {

	duration := -24 * time.Duration(days) * time.Hour

	future := time.Now().Add(duration)
	date := future.Format("2006-01-02")
	return date
}

func RandomPastDate(days int) string {

	random := rand.Intn(days - 1)
	return PastDate(random)
}

func PastDateTime(days int) string {

	duration := -24 * time.Duration(days) * time.Hour

	future := time.Now().Add(duration)
	date := future.Format(time.RFC3339Nano)
	return date
}

func RandomPastDateTime(days int) string {

	random := rand.Intn(days - 1)
	return PastDateTime(random)
}

func RandomFutureDateTime(days int) string {

	random := rand.Intn(days - 1)
	return FutureDateTime(random)
}
