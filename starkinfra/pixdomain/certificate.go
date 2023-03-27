package pixdomain

//	PixDomain.Certificate struct
//
//	The Certificate struct displays the certificate information from a specific domain.
//
//	Attributes (return-only):
//	- Content [string]: Certificate of the Pix participant in PEM format.

type Certificate struct {
	Content string `json:",omitempty"`
}
