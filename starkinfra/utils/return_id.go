package utils

import (
	"fmt"
)

//	Generates a random return-id based on your bank code (ISPB)
//
//	Parameters (required):
//	- bankCode [string]: Your bank code (ISPB). ex: "20018183"
//
//	Return:
//	- Random returnId based on your bank code.

func ReturnId(bankCode string) string {
	return fmt.Sprintf("D%v", BankCode(bankCode))
}
