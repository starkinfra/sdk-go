package pixdomain

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	PixDomain struct
//
//	The PixDomain struct displays the domain name and the QR Code domain certificate of Pix participants.
//	All certificates must be registered with the Central Bank.
//
//	Attributes (return-only):
//	- Certificates [slice of PixDomain.Certificate struct]: Certificate information of the Pix participant.
//	- Name [string]: Current active domain (URL) of the Pix participant.

type PixDomain struct {
	Certificates []Certificate `json:",omitempty"`
	Name         string        `json:",omitempty"`
}

var resource = map[string]string{"name": "PixDomain"}

func Query(user user.User) (chan PixDomain, chan Error.StarkErrors) {
	//	Retrieve PixDomain structs
	//
	//	Receive a channel of PixDomain structs of Pix participants able to issue BR Codes
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel  of PixDomain structs with updated attributes
	var pixDomain PixDomain
	domains := make(chan PixDomain)
	domainsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, nil, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixDomain)
			if err != nil {
				domainsError <- Error.UnknownError(err.Error())
				continue
			}
			domains <- pixDomain
		}
		for err := range errorChannel {
			domainsError <- err
		}
		close(domains)
		close(domainsError)
	}()
	return domains, domainsError
}
