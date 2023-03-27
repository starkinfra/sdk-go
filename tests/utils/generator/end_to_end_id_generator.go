package tax_id_generator

import "fmt"

func EndToEndId(bankCode string) string {

	// Generates a random end-to-end-id based on your bank code (ISPB)
	//
	// Parameters (required):
	// - bankCode [string]: Your bank code (ISPB). ex: "20018183"
	//
	// Return:
	// - Random endToEndId based on your bank code.
	return fmt.Sprintf("E%v", BacenId(bankCode))
}
