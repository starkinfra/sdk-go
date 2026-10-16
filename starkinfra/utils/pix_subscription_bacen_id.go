package utils

import (
	"fmt"
)

func PixSubscriptionBacenId(bankCode string, prefix string) string {

	//	Generates a random Pix subscription bacenId based on your bank code (ISPB)
	//	The date part (yyyyMMdd) is always generated in UTC, regardless of the machine timezone.
	//
	//	Parameters (required):
	//	- bankCode [string]: Your bank code (ISPB). ex: "20018183"
	//	- prefix [string]: Subscription prefix. ex: "RR"
	//
	//	Return:
	//	- Random bacenId based on your bank code.
	return fmt.Sprintf("%v%v", prefix, bacenIdWithLayout(bankCode, "20060102"))
}
