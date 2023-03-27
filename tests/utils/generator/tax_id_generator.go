package tax_id_generator

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func Cpf() string {
	var cpf string
	i := 0
	rand.Seed(time.Now().UTC().UnixNano())
	numbers := rand.Perm(9)
	numbers = append(numbers, calculateSpecialDigit(numbers, len(numbers)))
	numbers = append(numbers, calculateSpecialDigit(numbers, len(numbers)))

	for _, c := range numbers {
		cpf += strconv.Itoa(c)
		i += 1
	}

	format := regexp.MustCompile(`(?P<1>\d{3})(?P<2>\d{3})(?P<3>\d{3})(?P<4>\d{2})`)

	names := format.SubexpNames()
	formatted := fmt.Sprintf("${%s}.${%s}.${%s}-${%s}", names[1], names[2], names[3], names[4])

	return format.ReplaceAllString(cpf, formatted)
}

func Cnpj() string {
	var cnpj string
	i := 0
	rand.Seed(time.Now().UTC().UnixNano())
	numbers := rand.Perm(11)
	numbers = append(numbers, calculateSpecialDigit(numbers, len(numbers)))
	numbers = append(numbers, calculateSpecialDigit(numbers, len(numbers)))

	for _, c := range numbers {
		cnpj += strconv.Itoa(c)
		i += 1
	}

	format := regexp.MustCompile(`(?P<1>\d{2})(?P<2>\d{3})(?P<3>\d{3})(?P<4>\d{4})(?P<5>\d{2})`)

	names := format.SubexpNames()
	formatted := fmt.Sprintf("${%s}.${%s}.${%s}/${%s}-${%s}", names[1], names[2], names[3], names[4], names[5])

	return format.ReplaceAllString(cnpj, formatted)
}

func calculateSpecialDigit(data []int, n int) int {
	var total int

	for i := 0; i < n; i++ {
		total += data[i] * (n + 1 - i)
	}

	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}
