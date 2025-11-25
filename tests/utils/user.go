package utils

import (
	"github.com/starkinfra/core-go/starkcore/user/organization"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
	"os"
)

var PrivateKey, _ = checks.CheckPrivateKey(os.Getenv("PRIVATE_KEY"))  // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
var Environment, _ = checks.CheckEnvironment("sandbox")               // Environment: "sandbox"

var ExampleProject = project.Project{
	Id:          os.Getenv("PROJECT_ID"),                             // Id: "8888888888888888"
	PrivateKey:  PrivateKey,
	Environment: Environment,
}

var ExampleOrganization = organization.Organization{
	Id:          os.Getenv("ORGANIZATION_ID"),                        // Id: "8888888888888888"
	PrivateKey:  PrivateKey,
	Environment: Environment,
}
