package pix_domain

//	PixDomain object
//	The PixDomain object displays the domain name and the QR Code domain certificate of Pix participants.
//	All certificates must be registered with the Central Bank.
//
//	Attributes (return-only):
//	- certificates [list of pix_domain.Certificate]: certificate information of the Pix participant.
//	- name [string]: current active domain (URL) of the Pix participant.

type PixDomain struct {
	Certificates []Certificate
	Name         string
}

var resource = map[string]string{"class": PixDomain{}, "name": "PixDomain"}

func ParseCertificates() {

}

func Query() {
	//	Retrieve PixDomains
	//	Receive a generator of PixDomain objects.
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixDomain objects with updated attributes
}
