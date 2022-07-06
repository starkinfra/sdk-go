package pix_domain

//	pix_domain.Certificate object
//	The Certificate object displays the certificate information from a specific domain.
//
//	Attributes (return-only):
//	- content [string]: certificate of the Pix participant in PEM format.

type Certificate struct {
	Content string `json:"content"`
}

var resource = map[string]string{"class": Certificate{}, "name": "Certificate"}
