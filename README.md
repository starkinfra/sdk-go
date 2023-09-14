# Stark Infra Golang SDK

Welcome to the Stark Infra Golang SDK! This tool is made for Golang
developers who want to easily integrate with our API.
This SDK version is compatible with the Stark Infra API v2.

# Introduction

## Index

- [Introduction](#introduction)
    - [Supported Golang versions](#supported-golang-versions)
    - [API documentation](#stark-infra-api-documentation)
    - [Versioning](#versioning)
- [Setup](#setup)
    - [Install our SDK](#1-install-our-sdk)
    - [Create your Private and Public Keys](#2-create-your-private-and-public-keys)
    - [Register your user credentials](#3-register-your-user-credentials)
    - [Setting up the user](#4-setting-up-the-user)
    - [Setting up the error language](#5-setting-up-the-error-language)
- [Resource listing and manual pagination](#resource-listing-and-manual-pagination)
- [Testing in Sandbox](#testing-in-sandbox) 
- [Usage](#usage)
    - [Issuing](#issuing)
        - [Products](#query-issuingproducts): View available sub-issuer card products (a.k.a. card number ranges or BINs)
        - [Holders](#create-issuingholders): Manage card holders
        - [Cards](#create-issuingcards): Create virtual and/or physical cards
        - [Design](#query-issuingdesigns): View your current card or package designs
        - [EmbossingKit](#query-issuingembossingkits): View your current embossing kits
        - [Stock](#query-issuingstocks): View your current stock of a certain IssuingDesign linked to an Embosser on the workspace
        - [Restock](#create-issuingrestocks): Create restock orders of a specific IssuingStock object
        - [EmbossingRequest](#create-issuingembossingrequests): Create embossing requests
        - [Purchases](#process-purchase-authorizations): Authorize and view your past purchases
        - [Invoices](#create-issuinginvoices): Add money to your issuing balance
        - [Withdrawals](#create-issuingwithdrawals): Send money back to your Workspace from your issuing balance
        - [Balance](#get-your-issuingbalance): View your issuing balance
        - [Transactions](#query-issuingtransactions): View the transactions that have affected your issuing balance
        - [Enums](#issuing-enums): Query enums related to the issuing purchases, such as merchant categories, countries and card purchase methods
    - [Pix](#pix)
        - [PixRequests](#create-pixrequests): Create Pix transactions
        - [PixReversals](#create-pixreversals): Reverse Pix transactions
        - [PixBalance](#get-your-pixbalance): View your account balance
        - [PixStatement](#create-a-pixstatement): Request your account statement
        - [PixKey](#create-a-pixkey): Create a Pix Key
        - [PixClaim](#create-a-pixclaim): Claim a Pix Key
        - [PixDirector](#create-a-pixdirector): Create a Pix Director
        - [PixInfraction](#create-pixinfractions): Create Pix Infraction reports
        - [PixChargeback](#create-pixchargebacks): Create Pix Chargeback requests
        - [PixDomain](#query-pixdomains): View registered SPI participants certificates
        - [StaticBrcode](#create-staticbrcodes): Create static Pix BR codes
        - [DynamicBrcode](#create-dynamicbrcodes): Create dynamic Pix BR codes
        - [BrcodePreview](#create-brcodepreviews): Read data from BR Codes before paying them
    - [Lending](#lending)
        - [CreditNote](#create-creditnotes): Create credit notes
        - [CreditPreview](#create-creditpreviews): Create credit previews
        - [CreditHolmes](#create-creditholmes): Create credit holmes debt verification
    - [Identity](#identity)
        - [IndividualIdentity](#create-individualidentities): Create individual identities
        - [IndividualDocument](#create-individualdocuments): Create individual documents
    - [Webhook](#webhook):
        - [Webhook](#create-a-webhook-subscription): Configure your webhook endpoints and subscriptions
        - [WebhookEvents](#process-webhook-events): Manage Webhook events
        - [WebhookEventAttempts](#query-failed-webhook-event-delivery-attempts-information): Query failed webhook event deliveries
- [Handling errors](#handling-errors)
- [Help and Feedback](#help-and-feedback)

## Supported Golang Versions

This library supports the following Golang versions:

* Golang 1.17 or later

## Stark Infra API documentation

Feel free to take a look at our [API docs](https://www.starkinfra.com/docs/api).

## Versioning

This project adheres to the following versioning pattern:

Given a version number MAJOR.MINOR.PATCH, increment:

- MAJOR version when the **API** version is incremented. This may include backwards incompatible changes;
- MINOR version when **breaking changes** are introduced OR **new functionalities** are added in a backwards compatible
  manner;
- PATCH version when backwards compatible bug **fixes** are implemented.

# Setup

## 1. Install our SDK

1.1 In go.mod file, add the path in the required packages

```golang
github.com/starkinfra/sdk-go v0.1.2

```

1.2 You can also explicitly go get the package into a project:

```sh
go get -u github.com/starkinfra/sdk-go
```

## 2. Create your Private and Public Keys

We use ECDSA. That means you need to generate a secp256k1 private
key to sign your requests to our API, and register your public key
with us, so we can validate those requests.

You can use one of following methods:

2.1. Check out the options in our [tutorial](https://starkinfra.com/faq/how-to-create-ecdsa-keys).

2.2. Use our SDK:

```golang
package main

import (
    "github.com/starkinfra/core-go/starkcore/key"
)

func main() {

    privateKey, publicKey := key.Create("")

    // or, to also save .pem files in a specific path
    privateKey, publicKey := key.Create("files/keys/")
}

```

**NOTE**: When you are creating new credentials, it is recommended that you create the
keys inside the infrastructure that will use it, in order to avoid risky internet
transmissions of your **private-key**. Then you can export the **public-key** alone to the
computer where it will be used in the new Project creation.

## 3. Register your user credentials

You can interact directly with our API using two types of users: Projects and Organizations.

- **Projects** are workspace-specific users, that is, they are bound to the workspaces they are created in.
  One workspace can have multiple Projects.
- **Organizations** are general users that control your entire organization.
  They can control all your Workspaces and even create new ones. The Organization is bound to your company's tax ID
  only.
  Since this user is unique in your entire organization, only one credential can be linked to it.

3.1. To create a Project in Sandbox:

3.1.1. Log into [StarkInfra Sandbox](https://web.sandbox.starkinfra.com)

3.1.2. Go to Menu > Integrations

3.1.3. Click on the "New Project" button

3.1.4. Create a Project: Give it a name and upload the public key you created in section 2

3.1.5. After creating the Project, get its Project ID

3.1.6. Use the Project ID and private key to create the struct below:

```golang
package main

import (
    "github.com/starkinfra/core-go/starkcore/user/project"
    "github.com/starkinfra/core-go/starkcore/utils/checks"
)

// Get your private key from an environment variable or an encrypted database.
// This is only an example of a private key content. You should use your own key.

var privateKeyContent = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEILChZrjrrtFnyCLhcxm/hp+9ljWSmG7Wv9HRugf+FnhkoAcGBSuBBAAK\noUQDQgAEpIAM/tMqXEfLeR93rRHiFcpDB9I18MrnCJyTVk0MdD1J9wgEbRfvAZEL\nYcEGhTFYp2X3B7K7c4gDDCr0Pu1L3A==\n-----END EC PRIVATE KEY-----\n"

var project = project.Project{
    Id:          "5656565656565656",
    PrivateKey:  checks.CheckPrivateKey(privateKeyContent),
    Environment: checks.CheckEnvironment("sandbox"),
}

```

3.2. To create Organization credentials in Sandbox:

3.2.1. Log into [StarkInfra Sandbox](https://web.sandbox.starkinfra.com)

3.2.2. Go to Menu > Integrations

3.2.3. Click on the "Organization public key" button

3.2.4. Upload the public key you created in section 2 (only a legal representative of the organization can upload the
public key)

3.2.5. Click on your profile picture and then on the "Organization" menu to get the Organization ID

3.2.6. Use the Organization ID and private key to create the struct below:

```golang
package main

import (
    "github.com/starkinfra/core-go/starkcore/user/organization"
    "github.com/starkinfra/core-go/starkcore/utils/checks"
)

// Get your private key from an environment variable or an encrypted database.
// This is only an example of a private key content. You should use your own key.

var privateKeyContent = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEILChZrjrrtFnyCLhcxm/hp+9ljWSmG7Wv9HRugf+FnhkoAcGBSuBBAAK\noUQDQgAEpIAM/tMqXEfLeR93rRHiFcpDB9I18MrnCJyTVk0MdD1J9wgEbRfvAZEL\nYcEGhTFYp2X3B7K7c4gDDCr0Pu1L3A==\n-----END EC PRIVATE KEY-----\n"

var organization = organization.Organization{
    Id:          "5656565656565656",
    PrivateKey:  checks.CheckPrivateKey(privateKeyContent),
    Environment: checks.CheckEnvironment("sandbox"),
    Workspace:   "", //You only need to set the workspace_id when you are operating a specific WorkspaceId
}

```

NOTE 1: Never hard-code your private key. Get it from an environment variable or an encrypted database.

NOTE 2: We support `'sandbox'` and `'production'` as environments.

NOTE 3: The credentials you registered in `sandbox` do not exist in `production` and vice versa.

## 4. Setting up the user

There are three kinds of users that can access our API: **Organization**, **Project** and **Member**.

- `Project` and `Organization` are designed for integrations and are the ones meant for our SDKs.
- `Member` is the one you use when you log into our webpage with your e-mail.

There are two ways to inform the user to the SDK:

4.1 Passing the user as argument in all functions:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/pixbalance"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    balance := pixbalance.Get(nil) // or organization
    fmt.Println(balance)
}

```

4.2 Set it as a default user in the SDK:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/pixbalance"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    balance := pixbalance.Get(nil) // or organization
    fmt.Println(balance)
}

```

## 5. Setting up the error language

The error language can also be set in the same way as the default user:

```golang
package main

import (
    "github.com/starkinfra/sdk-go/starkinfra"
)

func main() {
    starkinfra.Language = "en-US"
}
```

Language options are "en-US" for english and "pt-BR" for brazilian portuguese. English is default.

# Resource listing and manual pagination

Almost all SDK resources provide a `query` and a `page` function.

- The `query` function provides a straight forward way, through a `channel`, to efficiently iterate through all results
  that match the filters you inform, seamlessly retrieving the next batch of elements from the API only when you reach
  the end of the current batch.
  If you are not worried about data volume or processing time, this is the way to go.

- In this function, in particular, when an error is encountered a `panic()` is raised with the code and error's message

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 200

    requests := PixRequest.Query(params, nil)
    for request := range requests {
        fmt.Println(request)
    }
}
```

- The `page` function gives you full control over the API pagination. With each function call, you receive up to
  100 results and the cursor to retrieve the next batch of elements. This allows you to stop your queries and
  pick up from where you left off whenever it is convenient. When there are no more elements to be retrieved, the
  returned cursor will be `nil`.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 50

    for true {
        requests, cursor, err := PixRequest.Page(params, nil)
        if err.Errors != nil {
            for _, e := range err.Errors {
                panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
            }
        }
        for _, request := range requests {
            fmt.Println(request)
        }
        if cursor == "" {
            break
        }
    }
}
```

To simplify the following SDK examples, we will only use the `query` function, but feel free to use `page` instead.

# Testing in Sandbox

Your initial balance is zero. For many operations in Stark Infra, you'll need funds
in your account, which can be added to your balance by creating a starkbank.Invoice. 

In the Sandbox environment, most of the created starkbank.Invoices will be automatically paid,
so there's nothing else you need to do to add funds to your account. Just create
a few starkbank.Invoice and wait around a bit.

In Production, you (or one of your clients) will need to actually pay this Pix Request
for the value to be credited to your account.

# Usage

Here are a few examples on how to use the SDK. If you have any doubts, check out the function or class docstring to get
more info or go straight to our [API docs].

## Issuing

### Query IssuingProducts

To take a look at the sub-issuer card products available to you, just run the following:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingProduct "github.com/starkinfra/sdk-go/starkinfra/issuingproduct"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 10

    products := IssuingProduct.Query(params, nil)
    for product := range products {
        fmt.Println(product)
    }
}

```

This will tell which card products and card number prefixes you have at your disposal.

### Create IssuingHolders

You can create card holders to which your cards will be bound.
They support spending rules that will apply to all underlying cards.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CardMethod "github.com/starkinfra/sdk-go/starkinfra/cardmethod"
    IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
    IssuingRule "github.com/starkinfra/sdk-go/starkinfra/issuingrule"
    MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
    MerchantCountry "github.com/starkinfra/sdk-go/starkinfra/merchantcountry"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    holders, err := IssuingHolder.Create(
        []IssuingHolder.IssuingHolder{
            {
                Name:       "General USD",
                ExternalId: "1234",
                TaxId:      "012.345.678-90",
                Tags: []string{
                    "Traveler Employee",
                },
                Rules: []IssuingRule.IssuingRule{
                    {
                        Name:         "General USD",
                        Interval:     "day",
                        Amount:       100000,
                        CurrencyCode: "USD",
                        Categories: []MerchantCategory.MerchantCategory{
                            {
                                Type: "services",
                            },
                            {
                                Code: "fastFoodRestaurants",
                            },
                        },
                        Countries: []MerchantCountry.MerchantCountry{
                            {
                                Code: "USA",
                            },
                        },
                        Methods: []CardMethod.CardMethod{
                            {
                                Code: "token",
                            },
                        },
                    },
                },
            },
        }, nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, holder := range holders {
        fmt.Println(holder)
    }
}

```

**Note**: Instead of using IssuingHolder objects, you can also pass each element in dictionary format

### Query IssuingHolders

You can query multiple holders according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 10

    holders := IssuingHolder.Query(params, nil)
    for holder := range holders {
        fmt.Println(holder)
    }
}

```

### Cancel an IssuingHolder

To cancel a single Issuing Holder by its id, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    holder, err := IssuingHolder.Cancel("5705125167366144", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(holder.Id)
}

```

### Get an IssuingHolder

To get a single Issuing Holder by its id, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    holder, err := IssuingHolder.Get("5705125167366144", nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(holder.Id)
}

```

### Query IssuingHolder logs

You can query holder logs to better understand holder life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingholder/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 10

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingHolder log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingholder/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5155165527080960", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create IssuingCards

You can issue cards with specific spending rules.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CardMethod "github.com/starkinfra/sdk-go/starkinfra/cardmethod"
    IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
    IssuingRule "github.com/starkinfra/sdk-go/starkinfra/issuingrule"
    MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
    MerchantCountry "github.com/starkinfra/sdk-go/starkinfra/merchantcountry"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    cards, err := IssuingCard.Create(
        []IssuingCard.IssuingCard{
            {
                HolderName:       "Developers",
                HolderTaxId:      "012.345.678-90",
                HolderExternalId: "1234",
                Rules: []IssuingRule.IssuingRule{
                    {
                        Name:         "General",
                        Interval:     "week",
                        Amount:       50000,
                        CurrencyCode: "USD",
                        Categories: []MerchantCategory.MerchantCategory{
                            {
                                Type: "services",
                                Code: "fastFoodRestaurants",
                            },
                        },
                        Countries: []MerchantCountry.MerchantCountry{
                            {
                                Code: "BRA",
                            },
                        },
                        Methods: []CardMethod.CardMethod{
                            {
                                Code: "token",
                            },
                        },
                    },
                },
            },
        }, nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, card := range cards {
        fmt.Println(card.Id)
    }
}

```

### Query IssuingCards

You can get a list of created cards given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 10

    cards := IssuingCard.Query(params, nil)
    for card := range cards {
        fmt.Println(card.Id)
    }
}

```

### Get an IssuingCard

After its creation, information on a card may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    card, err := IssuingCard.Get("5155165527080960", nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(card.Id)
}

```

### Update an IssuingCard

You can update a specific card by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var patchData = map[string]interface{}{}
    patchData["displayName"] = "ANTHONY EDWARD"

    card, err := IssuingCard.Update("5761721251659776", patchData, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(card.Id)
}

```

### Cancel an IssuingCard

You can also cancel a card by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    card, err := IssuingCard.Cancel("5632230982418432", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(card.Id)
}

```

### Query IssuingCard logs

Logs are pretty important to understand the life cycle of a card.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingcard/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 150

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingCard log

You can get a single log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingcard/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5155165527080960", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Query IssuingDesigns

You can get a list of available designs given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingDesign "github.com/starkinfra/sdk-go/starkinfra/issuingdesign"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    designs := IssuingDesign.Query(params, nil)
    for design := range designs {
        fmt.Println(design.Id)
    }
}

```

### Get an IssuingDesign

Information on a design may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingDesign "github.com/starkinfra/sdk-go/starkinfra/issuingdesign"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    design, err := IssuingDesign.Get("5747368922185728", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(design.Id)
}

```

### Query IssuingEmbossingKits

You can get a list of existing embossing kits given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
	IssuingEmbossingKit "github.com/starkinfra/sdk-go/starkinfra/issuingembossingkit"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    kits := IssuingEmbossingKit.Query(params, nil)
    for kit := range kits {
        fmt.Println(kit.Id)
    }
}

```

### Get an IssuingEmbossingKit

After its creation, information on an embossing kit may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingEmbossingKit "github.com/starkinfra/sdk-go/starkinfra/issuingembossingkit"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    kit, err := IssuingEmbossingKit.Get("5792731695677440", nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(kit.Id)
}

```

### Query IssuingStocks

You can get a list of available stocks given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingStock "github.com/starkinfra/sdk-go/starkinfra/issuingstock"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    stocks := IssuingStock.Query(params, nil)
    for stock := range stocks {
        fmt.Println(stock.Id)
    }
}

```

### Get an IssuingStock

Information on a stock may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingStock "github.com/starkinfra/sdk-go/starkinfra/issuingstock"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    stock, err := IssuingStock.Get("5792731695677440", nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(stock.Id)
}

```

### Query IssuingStock logs

Logs are pretty important to understand the life cycle of a stock.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingstock/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 150

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingStock log

You can get a single log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingstock/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create IssuingRestocks

You can order restocks for a specific IssuingStock.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingRestock "github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    restocks, err := IssuingRestock.Create(
        []IssuingRestock.IssuingRestock{
            {
                Count:   1000,
                StockId: "5152058940325888",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, restock := range restocks {
        fmt.Println(restock.Id)
    }
}

```

### Query IssuingRestocks

You can get a list of created restocks given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingRestock "github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 150

    restocks := IssuingRestock.Query(params, nil)
    for restock := range restocks {
        fmt.Println(restock.Id)
    }
}

```

### Get an IssuingRestock

After its creation, information on a restock may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingRestock "github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    restock, err := IssuingRestock.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(restock.Id)
}

```

### Query IssuingRestock logs

Logs are pretty important to understand the life cycle of a restock.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingrestock/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 1

    restocks := Log.Query(params, nil)
    for restock := range restocks {
        fmt.Println(restock.Id)
    }
}

```

### Get an IssuingRestock log

You can get a single log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingrestock/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5645672351006720", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create IssuingEmbossingRequests

You can create a request to emboss a physical card.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    requests, err := IssuingEmbossingRequest.Create(
        []IssuingEmbossingRequest.IssuingEmbossingRequest{
            {
                CardId:                 "5714424132272128",
                KitId:                  "5648359658356736",
                DisplayName1:           "teste",
                ShippingCity:           "Sao Paulo",
                ShippingCountryCode:    "BRA",
                ShippingDistrict:       "Bela Vista",
                ShippingService:        "loggi",
                ShippingStateCode:      "SP",
                ShippingStreetLine1:    "teste",
                ShippingStreetLine2:    "teste",
                ShippingTrackingNumber: "teste",
                ShippingZipCode:        "12345-678",
                EmbosserId:             "5634161670881280",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, request := range requests {
        fmt.Println(request.Id)
    }
}

```

### Query IssuingEmbossingRequests

You can get a list of created embossing requests given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    requests := IssuingEmbossingRequest.Query(params, nil)
    for request := range requests {
        fmt.Println(request.Id)
    }
}

```

### Get an IssuingEmbossingRequest

After its creation, information on an embossing request may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request, err := IssuingEmbossingRequest.Get("5478251505909760", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(request.Id)
}

```

### Query IssuingEmbossingRequest logs

Logs are pretty important to understand the life cycle of an embossing request.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingEmbossingRequest log

You can get a single log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("6398869424308224", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Process Purchase authorizations

It's easy to process purchase authorizations delivered to your endpoint.
Remember to pass the signature header so the SDK can make sure it's StarkInfra that sent you the event.
If you do not approve or decline the authorization within 2 seconds, the authorization will be denied.

```golang
package main

import (
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    authorization := issuingpurchase.Parse(
        request.Data,
        request.Headers["Digital-Signature"],
        "",
        nil,
    )

    var approved = map[string]interface{}{}
    approved["status"] = "approved"
    approved["reason"] = authorization.amount
    approved["tags"] = []string{"my-purchase-id/123"}

    sendResponse( // you should also implement this method
        issuingpurchase.Response( // this optional method just helps you build the response JSON
            approved,
        ),
    )

    // or

    var denied = map[string]interface{}{}
    denied["status"] = "denied"
    denied["reason"] = "other"
    denied["tags"] = []string{"other-id/456"}

    sendResponse(
        issuingpurchase.Response(
            denied,
        ),
    )
}

```

### Query IssuingPurchases

You can get a list of created purchases given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingPurchase "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    purchases := IssuingPurchase.Query(params, nil)
    for purchase := range purchases {
        fmt.Println(purchase.Id)
    }
}

```

### Get an IssuingPurchase

After its creation, information on a purchase may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingPurchase "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    purchase, err := IssuingPurchase.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(purchase.Id)
}

```

### Query IssuingPurchase logs

Logs are pretty important to understand the life cycle of a purchase.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingPurchase log

You can get a single log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```   

### Create IssuingInvoices

You can create Pix invoices to transfer money from accounts you have in any bank to your Issuing balance,
allowing you to run your issuing operation.

```golang
package main

import (
    "fmt"
    IssuingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
    Utils "infra-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    invoice, err := IssuingInvoice.Create(
        IssuingInvoice.IssuingInvoice{
            Amount: 12345,
            TaxId:  "012.345.678-90",
            Name:   "Jannie Lanister",
            Tags:   []string{"tony", "stark"},
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(invoice.Id)
}

```

**Note**: Instead of using IssuingInvoice objects, you can also pass each element in dictionary format

### Get an IssuingInvoice

After its creation, information on an invoice may be retrieved by its id.
Its status indicates whether it's been paid.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    invoice, err := IssuingInvoice.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(invoice.Id)
}

```

### Query IssuingInvoices

You can get a list of created invoices given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    invoices := IssuingInvoice.Query(params, nil)
    for invoice := range invoices {
        fmt.Println(invoice.Id)
    }
}

```

### Query IssuingInvoice logs

Logs are pretty important to understand the life cycle of an invoice.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IssuingInvoice log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create IssuingWithdrawals

You can create withdrawals to send cash back from your Issuing balance to your Banking balance
by using the Withdrawal resource.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingWithdrawal "github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    withdrawal, err := IssuingWithdrawal.Create(
        IssuingWithdrawal.IssuingWithdrawal{
            Amount:      123456,
            ExternalId:  "my_unique_external_1",
            Description: "testeIssuingWithdrawal",
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(withdrawal.Id)
}

```

**Note**: Instead of using IssuingWithdrawal objects, you can also pass each element in dictionary format

### Get an IssuingWithdrawal

After its creation, information on a withdrawal may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingWithdrawal "github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    withdrawal, err := IssuingWithdrawal.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(withdrawal.Id)
}

```

### Query IssuingWithdrawals

You can get a list of created withdrawals given some filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingWithdrawal "github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    withdrawals := IssuingWithdrawal.Query(params, nil)
    for withdrawal := range withdrawals {
        fmt.Println(withdrawal.Id)
    }
}

```

### Get your IssuingBalance

To know how much money you have available to run authorizations, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingBalance "github.com/starkinfra/sdk-go/starkinfra/issuingbalance"
    "github.com/starkinfra/sdk-go/tests/utils"
    Utils "infra-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    balance := IssuingBalance.Get(nil)
    fmt.Println(balance.Amount)
}

```

### Query IssuingTransactions

To understand your balance changes (issuing statement), you can query
transactions. Note that our system creates transactions for you when
you make purchases, withdrawals, receive issuing invoice payments, for example.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingTransaction "github.com/starkinfra/sdk-go/starkinfra/issuingtransaction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    transactions := IssuingTransaction.Query(params, nil)
    for transaction := range transactions {
        fmt.Println(transaction.Id)
    }
}

```

### Get an IssuingTransaction

You can get a specific transaction by its id:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IssuingTransaction "github.com/starkinfra/sdk-go/starkinfra/issuingtransaction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    transaction, err := IssuingTransaction.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(transaction.Id)
}

```

### Issuing Enums

#### Query MerchantCategories

You can query any merchant categories using this resource.
You may also use MerchantCategories to define specific category filters in IssuingRules.
Either codes (which represents specific MCCs) or types (code groups) will be accepted as filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    categories := MerchantCategory.Query(nil, nil)
    for category := range categories {
        fmt.Println(category.Code)
    }
}

```

#### Query MerchantCountries

You can query any merchant countries using this resource.
You may also use MerchantCountries to define specific country filters in IssuingRules.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    MerchantCountry "github.com/starkinfra/sdk-go/starkinfra/merchantcountry"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    countries := MerchantCountry.Query(nil, nil)
    for country := range countries {
        fmt.Println(country.Code)
    }
}

```

#### Query CardMethods

You can query available card methods using this resource.
You may also use CardMethods to define specific purchase method filters in IssuingRules.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CardMethod "github.com/starkinfra/sdk-go/starkinfra/cardmethod"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    methods := CardMethod.Query(nil, nil)
    for method := range methods {
        fmt.Println(method.Code)
    }
}

```

## Pix

### Create PixRequests

You can create a Pix request to transfer money from one of your users to anyone else:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
	PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
	Utils "github.com/starkinfra/sdk-go/starkinfra/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    requests, err := PixRequest.Create(
        []PixRequest.PixRequest{
            {
                Amount:                12345,
                ExternalId:            "141234121",
                SenderName:            "Edward Stark",
                SenderTaxId:           "20.018.183/0001-80",
                SenderBranchCode:      "1357-9",
                SenderAccountNumber:   "876543-2",
                SenderAccountType:     "checking",
                ReceiverName:          "Edward Stark",
                ReceiverTaxId:         "01234567890",
                ReceiverBankCode:      "20018183",
                ReceiverAccountNumber: "876543-2",
                ReceiverBranchCode:    "1357-9",
                ReceiverAccountType:   "payment",
                EndToEndId:            Utils.EndToEndId("35547753"),
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, request := range requests {
        fmt.Println(request.Id)
    }
}

```

**Note**: Instead of using PixRequest objects, you can also pass each element in dictionary format

### Query PixRequests

You can query multiple Pix requests according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    requests := PixRequest.Query(params, nil)
    for request := range requests {
        fmt.Println(request.Id)
    }
}

```

### Get a PixRequest

After its creation, information on a Pix request may be retrieved by its id. Its status indicates whether it has been
paid.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request, err := PixRequest.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(request.Id)
}

```

### Process inbound PixRequest authorizations

It's easy to process authorization requests that arrived at your endpoint.
Remember to pass the signature header so the SDK can make sure it's StarkInfra that sent you the event.
If you do not approve or decline the authorization within 1 second, the authorization will be denied.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    pixRequest := pixrequest.Parse(
        request.Data,
        request.Headers["Digital-Signature"],
        "",
        nil,
    )

    fmt.Println(pixRequest)

    var approved = map[string]interface{}{}
    approved["status"] = "approved"

    sendResponse( // you should also implement this method
        pixrequest.Response( // this optional method just helps you build the response JSON
            approved,
        ),
    )

    // or

    var denied = map[string]interface{}{}
    denied["status"] = "denied"
    denied["reason"] = "orderRejected"

    sendResponse(
        pixrequest.Response(
            denied,
        ),
    )
}

```

### Query PixRequest logs

You can query Pix request logs to better understand Pix request life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixrequest/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixRequest log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixrequest/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create PixReversals

You can reverse a PixRequest either partially or totally using a PixReversal.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    requests, err := PixRequest.Create(
        []PixRequest.PixRequest{
            {
                Amount:                12345,
                ExternalId:            "my_external_id_unique",
                SenderName:            "Edward Stark",
                SenderTaxId:           "20.018.183/0001-80",
                SenderBranchCode:      "1357-9",
                SenderAccountNumber:   "876543-2",
                SenderAccountType:     "checking",
                ReceiverName:          "Edward Stark",
                ReceiverTaxId:         "01234567890",
                ReceiverBankCode:      "20018183",
                ReceiverAccountNumber: "876543-2",
                ReceiverBranchCode:    "1357-9",
                ReceiverAccountType:   "payment",
                EndToEndId:            EndToEndId("35547753"),
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, request := range requests {
        fmt.Println(request.Id)
    }
}

```

### Query PixReversals

You can query multiple Pix reversals according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixReversal "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    reversals := PixReversal.Query(params, nil)
    for reversal := range reversals {
        fmt.Println(reversal.Id)
    }
}

```

### Get a PixReversal

After its creation, information on a Pix reversal may be retrieved by its id.
Its status indicates whether it has been successfully processed.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixReversal "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    reversal, err := PixReversal.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(reversal.Id)
}

```

### Process inbound PixReversal authorizations

It's easy to process authorization requests that arrived at your endpoint.
Remember to pass the signature header so the SDK can make sure it's StarkInfra that sent you the event.
If you do not approve or decline the authorization within 1 second, the authorization will be denied.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    reversal := pixreversal.Parse(
        request.Data,
        request.Headers["Digital-Signature"],
        "",
        nil,
    )

    fmt.Println(reversal)

    var response := map[string]interface{}{}
    response["status"] = "approved"

    sendResponse( // you should also implement this method
        pixreversal.Response( // this optional method just helps you build the response JSON
            response,
        ),
    )

    // or

    var response := map[string]interface{}{}
    response["status"] = "denied"
    response["reason"] = "orderRejected"

    sendResponse(
        pixreversal.Response(
            response,
        ),
    )
}

```

### Query PixReversal logs

You can query Pix reversal logs to better understand their life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixreversal/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixReversal log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixreversal/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Get your PixBalance

To see how much money you have in your account, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixBalance "github.com/starkinfra/sdk-go/starkinfra/pixbalance"
    "github.com/starkinfra/sdk-go/tests/utils"
    Utils "infra-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    balance := PixBalance.Get(nil)
    fmt.Println(balance.Amount)
}

```

### Create a PixStatement

Statements are generated directly by the Central Bank and are only available for direct participants.
To create a statement of all the transactions that happened on your account during a specific day, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    statement, err := PixStatement.Create(
        PixStatement.PixStatement{
            After:  "2022-05-01T12:00:00:00",
            Before: "2022-07-01T12:00:00:00",
            Types:  "interchange",
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(statement.Id)
}

```

### Query PixStatements

You can query multiple Pix statements according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    statements := PixStatement.Query(params, nil)
    for statement := range statements {
        fmt.Println(statement.Id)
    }
}

```

### Get a PixStatement

Statements are only available for direct participants. To get a Pix statement by its id:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    statement, err := PixStatement.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(statement.Id)
}

```

### Get a PixStatement .csv file

To get the .csv file corresponding to a Pix statement using its id, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
    "github.com/starkinfra/sdk-go/tests/utils"
    "io/ioutil"
)

func main() {

    starkinfra.User = utils.ExampleProject

    csv, err := PixStatement.Csv("5155165527080960", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    filename := fmt.Sprintf("%v%v.csv", "pix-statement", "5155165527080960")
    ioutil.WriteFile(filename, csv, 0666)
}

```

### Create a PixKey

You can create a Pix Key to link a bank account information to a key id:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
    "github.com/starkinfra/sdk-go/tests/utils"
    "time"
)

func main() {

    starkinfra.User = utils.ExampleProject

    accountCreated := time.Date(2022, 02, 01, 0, 0, 0, 0, time.UTC)

    key, err := PixKey.Create(
        PixKey.PixKey{
            AccountCreated: &accountCreated,
            AccountNumber:  "76543",
            AccountType:    "checking",
            BranchCode:     "1234",
            Name:           "Antony",
            TaxId:          "012.345.678-90",
            Id:             "+5511989898989",
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(key.Id)
}

```

### Query PixKeys

You can query multiple Pix keys you own according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    keys := PixKey.Query(params, nil)
    for key := range keys {
        fmt.Println(key.Id)
    }
}

```

### Get a PixKey

Information on a Pix key may be retrieved by its id and the tax ID of the consulting agent.
An endToEndId must be informed so you can link any resulting purchases to this query,
avoiding sweep blocks by the Central Bank.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    key, err := PixKey.Get("5792731695677440", nil, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(key.Id)
}

```

### Update a PixKey

Update the account information linked to a Pix Key.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var patchData = map[string]interface{}{}
    patchData["reason"] = "branchTransfer"
    patchData["accountType"] = "savings"

    key, err := PixKey.Update("+5582764344485", patchData, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(key.Id)
}

```

### Cancel a PixKey

Cancel a specific Pix Key using its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    key, err := PixKey.Cancel("+5517731549401", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(key.Id)
}

```

### Query PixKey logs

You can query Pix key logs to better understand a Pix key life cycle.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixkey/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixKey log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixkey/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create a PixClaim

You can create a Pix claim to request the transfer of a Pix key from another bank to one of your accounts:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    claim, err := PixClaim.Create(
        PixClaim.PixClaim{
            AccountCreated: "2022-01-01",
            AccountNumber:  "76543",
            AccountType:    "checking",
            BranchCode:     "1234",
            Name:           "Jamie Lannister",
            TaxId:          "40.516.230/0001-22",
            KeyId:          "+5511989898989",
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(claim.Id)
}

```

### Query PixClaims

You can query multiple Pix claims according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    claims := PixClaim.Query(params, nil)
    for claim := range claims {
        fmt.Println(claim.Id)
    }
}

```

### Get a PixClaim

After its creation, information on a Pix claim may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    claim, err := PixClaim.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(claim.Id)
}

```

### Update a PixClaim

A Pix Claim can be confirmed or canceled by patching its status.
A received Pix Claim must be confirmed by the donor to be completed.
Ownership Pix Claims can only be canceled by the donor if the reason is "fraud".
A sent Pix Claim can also be canceled.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var patchData = map[string]interface{}{}
    patchData["status"] = "canceled"
    patchData["reason"] = "userRequested"

    claim, err := PixClaim.Update("6608972270272512", patchData, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(claim.Id)
}

```

### Query PixClaim logs

You can query Pix claim logs to better understand Pix claim life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixclaim/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixClaim log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixclaim/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create a PixDirector

To register the Pix director contact information at the Central Bank, run the following:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixDirector "github.com/starkinfra/sdk-go/starkinfra/pixdirector"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    director, err := PixDirector.Create(
        PixDirector.PixDirector{
            Name:       "Edward Stark",
            TaxId:      "03.300.300/0001-00",
            Phone:      "+5511999999999",
            Email:      "ned.stark@company.com",
            Password:   "12345678",
            TeamEmail:  "pix.team@company.com",
            TeamPhones: []string{"+5511989898989"},
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(director.Name)
}

```

### Create PixInfractions

Pix Infraction reports are used to report transactions that raise fraud suspicion, to request a refund or to
reverse a refund. Infraction reports can be created by either participant of a transaction.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    infractions, err := PixInfraction.Create(
        []PixInfraction.PixInfraction{
            {
                ReferenceId: "E35547753202201201450oo8sDGca066",
                Type:        "fraud",
                Description: "testInfractionGolang",
                Tags:        []string{"tony", "stark"},
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, infraction := range infractions {
        fmt.Println(infraction.Id)
    }
}

```

### Query PixInfractions

You can query multiple infraction reports according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    infractions := PixInfraction.Query(params, nil)
    for infraction := range infractions {
        fmt.Println(infraction.Id)
    }
}

```

### Get a PixInfraction

After its creation, information on a Pix Infraction may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    infraction, err := PixInfraction.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(infraction.Id)
}

```

### Update a PixInfraction

A received Pix Infraction can be confirmed or declined by patching its status.
After a Pix Infraction is patched, its status changes to closed.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var patchData = map[string]interface{}{}
    patchData["result"] = "agreed"
    patchData["analysis"] = "Upon investigation fraud was confirmed."

    infraction, err := PixInfraction.Update("5181903216836608", patchData, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(infraction.Id)
}

```

### Cancel a PixInfraction

Cancel a specific Pix Infraction using its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    infraction, err := PixInfraction.Cancel("4772078074986496", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(infraction.Id)
}

```

### Query PixInfraction logs

You can query infraction report logs to better understand their life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixinfraction/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixInfraction log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixinfraction/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create PixChargebacks

A Pix chargeback can be created when fraud is detected on a transaction or a system malfunction
results in an erroneous transaction.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    chargebacks, err := PixChargeback.Create(
        []PixChargeback.PixChargeback{
            {
                Amount:      123456,
                ReferenceId: "E20018183202201201450u34sDGd19lz",
                Reason:      "fraud",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, chargeback := range chargebacks {
        fmt.Println(chargeback.Id)
    }
}

```

### Query PixChargebacks

You can query multiple Pix chargebacks according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    chargebacks := PixChargeback.Query(params, nil)
    for chargeback := range chargebacks {
        fmt.Println(chargeback.Id)
    }
}

```

### Get a PixChargeback

After its creation, information on a Pix Chargeback may be retrieved by its.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    chargeback, err := PixChargeback.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(chargeback.Id)
}

```

### Update a PixChargeback

A received Pix Chargeback can be accepted or rejected by patching its status.
After a Pix Chargeback is patched, its status changes to closed.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var patchData = map[string]interface{}{}
    patchData["result"] = "rejected"
    patchData["rejectionReason"] = "noBalance"

    chargeback, err := PixChargeback.Update("4848592950919168", patchData, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(chargeback.Id)
}

```

### Cancel a PixChargeback

Cancel a specific Pix Chargeback using its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    chargeback, err := PixChargeback.Cancel("4848592950919168", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(chargeback.Id)
}

```

### Query PixChargeback logs

You can query Pix chargeback logs to better understand Pix chargeback life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixchargeback/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a PixChargeback log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/pixchargeback/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Query PixDomains

Here you can list all Pix Domains registered at the Brazilian Central Bank. The Pix Domain object displays the domain
name and the QR Code domain certificates of registered Pix participants able to issue dynamic QR Codes.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixDomain "github.com/starkinfra/sdk-go/starkinfra/pixdomain"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    domains := PixDomain.Query(nil)
    for domain := range domains {
        fmt.Println(domain.Name)
    }
}

```

### Create StaticBrcodes

StaticBrcodes store account information via a BR code or an image (QR code)
that represents a PixKey and a few extra fixed parameters, such as an amount
and a reconciliation ID. They can easily be used to receive Pix transactions.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    StaticBrcode "github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    brcodes, err := StaticBrcode.Create(
        []StaticBrcode.StaticBrcode{
            {
                Name:  "Tony Stark",
                KeyId: "+5511989898989",
                City:  "Rio de Janeiro",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, brcode := range brcodes {
        fmt.Println(brcode.Id)
    }
}

```

### Query StaticBrcodes

You can query multiple StaticBrcodes according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    StaticBrcode "github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    brcodes := StaticBrcode.Query(params, nil)
    for brcode := range brcodes {
        fmt.Println(brcode.Id)
    }
}

```

### Get a StaticBrcodes

After its creation, information on a StaticBrcode may be retrieved by its UUID.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    StaticBrcode "github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    brcode, err := StaticBrcode.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(brcode.Id)
}

```

### Create DynamicBrcodes

BR codes store information represented by Pix QR Codes, which are used to send
or receive Pix transactions in a convenient way.
DynamicBrcodes represent charges with information that can change at any time,
since all data needed for the payment is requested dynamically to an URL stored
in the BR Code. Stark Infra will receive the GET request and forward it to your
registered endpoint with a GET request containing the UUID of the BR code for
identification.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    DynamicBrcode "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    brcodes, err := DynamicBrcode.Create(
        []DynamicBrcode.DynamicBrcode{
            {
                Name:       "Jamie Lannister",
                City:       "Rio de Janeiro",
                ExternalId: "my_unique_id_01",
                Type:       "instant",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, brcode := range brcodes {
        fmt.Println(brcode.Id)
    }
}

```

### Query DynamicBrcodes

You can query multiple DynamicBrcodes according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    DynamicBrcode "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    brcodes := DynamicBrcode.Query(params, nil)
    for brcode := range brcodes {
        fmt.Println(brcode.Id)
    }
}

```

### Get a DynamicBrcode

After its creation, information on a DynamicBrcode may be retrieved by its UUID.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    DynamicBrcode "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    brcode, err := DynamicBrcode.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(brcode.Id)
}

```

### Verify a DynamicBrcode read

When a DynamicBrcode is read by your user, a GET request will be made to the your regitered URL to
retrieve additional information needed to complete the transaction.
Use this method to verify the authenticity of a GET request received at your registered endpoint.
If the provided digital signature does not check out with the StarkInfra public key, a
stark.exception.InvalidSignatureException will be raised.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    uuid := dynamicbrcode.Verify(
        request.Uuid,
        request.Headers["Digital-Signature"],
        "",
        utils.ExampleProject,
    )

    fmt.Println(uuid)
}

```

### Answer to a Due DynamicBrcode read

When a Due DynamicBrcode is read by your user, a GET request containing
the BR code UUID will be made to your registered URL to retrieve additional
information needed to complete the transaction.

The GET request must be answered in the following format within 5 seconds
and with an HTTP status code 200.

```golang
package main

import (
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    uuid := dynamicbrcode.Verify(
        request.Uuid,
        request.Headers["Digital-Signature"],
        "",
        utils.ExampleProject,
    )

    invoice := getMyInvoice(uuid) // you should implement this method to get the information of the BR code from its uuid

    var data = map[string]interface{}{}
    data["version"] = invoice.Version
    data["created"] = invoice.Created
    data["due"] = invoice.Due
    data["keyId"] = invoice.KeyId
    data["status"] = invoice.Status
    data["reconciliationId"] = invoice.ReconciliationId
    data["amount"] = invoice.Amount
    data["senderName"] = invoice.SenderName
    data["receiverName"] = invoice.ReceiverName
    data["receiverStreetLine"] = invoice.ReceiverStreetLine
    data["receiverCity"] = invoice.ReceiverCity
    data["receiverStateCode"] = invoice.ReceiverStateCode
    data["receiverZipCode"] = invoice.ReceiverZipCode

    sendResponse( // you should also implement this method to respond the read request
        dynamicbrcode.ResponseDue(data),
    )
}

```

### Answer to an Instant DynamicBrcode read

When an Instant DynamicBrcode is read by your user, a GET request
containing the BR code UUID will be made to your registered URL to retrieve
additional information needed to complete the transaction.

The get request must be answered in the following format
within 5 seconds and with an HTTP status code 200.

```golang
package main

import (
    "github.com/starkinfra/sdk-go/starkinfra"
    "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    uuid := dynamicbrcode.Verify(
        request.Uuid,
        request.Headers["Digital-Signature"],
        "",
        utils.ExampleProject,
    )

    invoice := getMyInvoice(uuid) // you should implement this method to get the information of the BR code from its uuid

    var data = map[string]interface{}{}
    data["version"] = invoice.Version
    data["created"] = invoice.Created
    data["keyId"] = invoice.KeyId
    data["status"] = invoice.Status
    data["reconciliationId"] = invoice.ReconciliationId
    data["amount"] = invoice.Amount
    data["cashierType"] = invoice.CashierType
    data["cashierBankCode"] = invoice.CashierBankCode
    data["cashAmount"] = invoice.CashAmount

    sendResponse( // you should also implement this method to respond the read request
        dynamicbrcode.ResponseInstant(data),
    )
}

```

## Create BrcodePreviews

You can create BrcodePreviews to preview BR Codes before paying them.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    BrcodePreview "github.com/starkinfra/sdk-go/starkinfra/brcodepreview"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    previews, err := BrcodePreview.Create(
        []BrcodePreview.BrcodePreview{
            {
				Id: "00020126420014br.gov.bcb.pix0120nedstark@hotmail.com52040000530398654075000.005802BR5909Ned Stark6014Rio de Janeiro621605126674869738606304FF71", 
				PayerId: "123.456.78-90",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, preview := range previews {
        fmt.Println(preview.Id)
    }
}

```

## Lending

If you want to establish a lending operation, you can use Stark Infra to
create a CCB contract. This will enable your business to lend money without
requiring a banking license, as long as you use a Credit Fund
or Securitization company.

The required steps to initiate the operation are:

1. Have funds in your Credit Fund or Securitization account
2. Request the creation of an [Identity Check](#create-individualidentities)
   for the credit receiver (make sure you have their documents and express authorization)
3. (Optional) Create a [Credit Simulation](#create-creditpreviews)
   with the desired installment plan to display information for the credit receiver
4. Create a [Credit Note](#create-creditnotes)
   with the desired installment plan

### Create CreditNotes

For lending operations, you can create a CreditNote to generate a CCB contract.

Note that you must have recently created an identity check for that same Tax ID before
being able to create a credit operation for them.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
    Invoice "github.com/starkinfra/sdk-go/starkinfra/creditnote/invoice"
    CreditSigner "github.com/starkinfra/sdk-go/starkinfra/creditsigner"
    "github.com/starkinfra/sdk-go/tests/utils"
    "math/rand"
    "time"
)

func main() {

    starkinfra.User = utils.ExampleProject

    scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

    notes, err := CreditNote.Create(
        []CreditNote.CreditNote{
            {
                TemplateId:    "5706627130851328",
                Name:          "Jamie Lannister",
                TaxId:         "012.345.678-90",
                NominalAmount: rand.Intn(100000 - 1),
                Scheduled:     &scheduled,
                Invoices: []Invoice.Invoice{
                    {
                        Due:    "2023-01-12",
                        Amount: rand.Intn(100000 - 1),
                    },
                },
                Payment: CreditNote.Transfer{
                    BranchCode:    "1234",
                    BankCode:      "00000000",
                    AccountNumber: "129340-1",
                    Name:          "Jaime Lannister",
                    TaxId:         "012.345.678-90",
                },
                Signers: []CreditSigner.CreditSigner{
                    {
                        Name:    "Jaime Lannister",
                        Contact: "jaime.lannister@gmail.com",
                        Method:  "link",
                    },
                },
                ExternalId:  "1234",
                StreetLine1: "Rua ABC",
                StreetLine2: "Ap 123",
                District:    "Jardim Paulista",
                City:        "São Paulo",
                StateCode:   "SP",
                ZipCode:     "01234-567",
                PaymentType: "transfer",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, note := range notes {
        fmt.Println(note.Id)
    }
}

```

**Note**: Instead of using CreditNote objects, you can also pass each element in dictionary format

### Query CreditNotes

You can query multiple credit notes according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    notes := CreditNote.Query(params, nil)
    for note := range notes {
        fmt.Println(note.Id)
    }
}

```

### Get a CreditNote

After its creation, information on a credit note may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    note, err := CreditNote.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(note.Id)
}

```

### Cancel a CreditNote

You can cancel a credit note if it has not been signed yet.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    note, err := CreditNote.Cancel("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(note.Id)
}

```

### Query CreditNote logs

You can query credit note logs to better understand credit note life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/creditnote/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get a CreditNote log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/creditnote/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create CreditPreviews

You can preview a credit operation before creating them (Currently we only have CreditNote / CCB previews):

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditPreview "github.com/starkinfra/sdk-go/starkinfra/creditpreview"
    "github.com/starkinfra/sdk-go/tests/utils"
    "time"
)

func main() {

    starkinfra.User = utils.ExampleProject

    scheduled := time.Date(2023, 01, 10, 0, 0, 0, 0, time.UTC)
    initialDue := time.Date(2023, 01, 20, 0, 0, 0, 0, time.UTC)

    previews, err := CreditPreview.Create(
        []CreditPreview.CreditPreview{
            {
                Credit: CreditPreview.CreditNotePreview{
                    Type:            "sac",
                    NominalAmount:   100000,
                    Scheduled:       &scheduled,
                    TaxId:           "012.345.678-90",
                    InitialDue:      &initialDue,
                    NominalInterest: 10.0,
                    Count:           3,
                    Interval:        "week",
                },
                Type: "credit-note",
            },
            {
                Credit: CreditPreview.CreditNotePreview{
                    Type:            "price",
                    NominalAmount:   100000,
                    TaxId:           "012.345.678-90",
                    Scheduled:       &scheduled,
                    InitialDue:      &initialDue,
                    NominalInterest: 10,
                    Count:           3,
                    Interval:        "year",
                },
                Type: "credit-note",
            },
            {
                Credit: CreditPreview.CreditNotePreview{
                    Type:            "american",
                    NominalAmount:   100000,
                    Scheduled:       &scheduled,
                    TaxId:           "012.345.678-90",
                    InitialDue:      &initialDue,
                    NominalInterest: 10,
                    Count:           3,
                    Interval:        "month",
                },
                Type: "credit-note",
            },
            {
                Credit: CreditPreview.CreditNotePreview{
                    TaxId:           "012.345.678-90",
                    Type:            "bullet",
                    NominalAmount:   100000,
                    Scheduled:       &scheduled,
                    InitialDue:      &initialDue,
                    NominalInterest: 10,
                },
                Type: "credit-note",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, preview := range previews {
        fmt.Println(preview.Type)
    }
}

```

**Note**: Instead of using CreditPreview objects, you can also pass each element in dictionary format

### Create CreditHolmes

Before you request a credit operation, you may want to check previous credit operations
the credit receiver has taken.

For that, open up a CreditHolmes investigation to receive information on all debts and credit
operations registered for that individual or company inside the Central Bank's SCR.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditHolmes "github.com/starkinfra/sdk-go/starkinfra/creditholmes"
    "github.com/starkinfra/sdk-go/tests/utils"
    "time"
)

func main() {

    starkinfra.User = utils.ExampleProject

    notes, err := CreditHolmes.Create(
        []CreditHolmes.CreditHolmes{
            {
                TaxId:      "123.456.789-00",
                Competence: "2023-01-27",
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, note := range notes {
        fmt.Println(note.Id)
    }
}

```

### Query CreditHolmes

You can query multiple credit holmes according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditHolmes "github.com/starkinfra/sdk-go/starkinfra/creditholmes"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    holmes := CreditHolmes.Query(params, nil)
    for sherlock := range holmes {
        fmt.Println(sherlock.Id)
    }
}

```

### Get an CreditHolmes

After its creation, information on a credit holmes may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    CreditHolmes "github.com/starkinfra/sdk-go/starkinfra/creditholmes"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    holmes, err := CreditHolmes.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(holmes.Id)
}

```

### Query CreditHolmes logs

You can query credit holmes logs to better understand their life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/creditholmes/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an CreditHolmes log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/creditholmes/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

## Identity

Several operations, especially credit ones, require that the identity
of a person or business is validated beforehand.

Identities are validated according to the following sequence:

1. The Identity resource is created for a specific Tax ID
2. Documents are attached to the Identity resource
3. The Identity resource is updated to indicate that all documents have been attached
4. The Identity is sent for validation and returns a webhook notification to reflect
   the success or failure of the operation

### Create IndividualIdentities

You can create an IndividualIdentity to validate a document of a natural person

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    identities, err := IndividualIdentity.Create(
        []IndividualIdentity.IndividualIdentity{
            {
                Name:  "Walter White",
                TaxId: "012.345.678-90",
                Tags:  []string{"breaking", "bad"},
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, identity := range identities {
        fmt.Println(identity.Id)
    }
}

```

**Note**: Instead of using IndividualIdentity objects, you can also pass each element in dictionary format

### Query IndividualIdentity

You can query multiple individual identities according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    identities := IndividualIdentity.Query(params, nil)
    for identity := range identities {
        fmt.Println(identity.Id)
    }
}

```

### Get an IndividualIdentity

After its creation, information on an individual identity may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    identity, err := IndividualIdentity.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(identity.Id)
}

```

### Update an IndividualIdentity

You can update a specific identity status to "processing" for send it to validation.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    identity, err := IndividualIdentity.Update("5761721251659776", "processing", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(identity.Id)
}

```

**Note**: Before sending your individual identity to validation by patching its status, you must send all the required
documents using the create method of the CreditDocument resource. Note that you must reference the individual identity
in the create method of the CreditDocument resource by its id.

### Cancel an IndividualIdentity

You can cancel an individual identity before updating its status to processing.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    identity, err := IndividualIdentity.Cancel("5761721251659776", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(identity.Id)
}

```

### Query IndividualIdentity logs

You can query individual identity logs to better understand individual identity life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/individualidentity/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IndividualIdentity log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/individualidentity/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

### Create IndividualDocuments

You can create an individual document to attach images of documents to a specific individual Identity.
You must reference the desired individual identity by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    documents, err := IndividualDocument.Create(
        []IndividualDocument.IndividualDocument{
            {
                Type:       "Walter White",
                Content:    "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD...",
                IdentityId: "012.345.678-90",
                Tags:       []string{"breaking", "bad"},
            },
            {
                Type:       "Walter White",
                Content:    "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD...",
                IdentityId: "012.345.678-90",
                Tags:       []string{"breaking", "bad"},
            },
            {
                Type:       "Walter White",
                Content:    "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD...",
                IdentityId: "012.345.678-90",
                Tags:       []string{"breaking", "bad"},
            },
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    for _, document := range documents {
        fmt.Println(document.Id)
    }
}

```

**Note**: Instead of using IndividualDocument objects, you can also pass each element in dictionary format

### Query IndividualDocument

You can query multiple individual documents according to filters.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    documents := IndividualDocument.Query(params, nil)
    for document := range documents {
        fmt.Println(document.Id)
    }
}

```

### Get an IndividualDocument

After its creation, information on an individual document may be retrieved by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    document, err := IndividualDocument.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(document.Id)
}

```

### Query IndividualDocument logs

You can query individual document logs to better understand individual document life cycles.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/individualdocument/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    logs := Log.Query(params, nil)
    for log := range logs {
        fmt.Println(log.Id)
    }
}

```

### Get an IndividualDocument log

You can also get a specific log by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Log "github.com/starkinfra/sdk-go/starkinfra/individualdocument/log"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    log, err := Log.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(log.Id)
}

```

## Webhook

### Create a webhook subscription

To create a webhook subscription and be notified whenever an event occurs, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    webhook, err := Webhook.Create(
        Webhook.Webhook{
            Url:           "https://webhook.site/dd784f26-1d6a-4ca6-81cb-fda0267761ec",
            Subscriptions: []string{"boleto"},
        }, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(webhook.Id)
}

```

### Query webhooks

To search for registered webhooks, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    webhooks := Webhook.Query(params, nil)
    for webhook := range webhooks {
        fmt.Println(webhook.Id)
    }
}

```

### Get a webhook

You can get a specific webhook by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    webhook, err := Webhook.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(webhook.Id)
}

```

### Delete a webhook

You can also delete a specific webhook by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    webhook, err := Webhook.Delete("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(webhook.Id)
}

```

### Process webhook events

It's easy to process events delivered to your Webhook endpoint.
Remember to pass the signature header so the SDK can make sure it was StarkInfra that sent you the event.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Event "github.com/starkinfra/sdk-go/starkinfra/event"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    request := listen() // this is the method you made to get the events posted to your webhook endpoint

    event := Event.Parse(
        request.Data,
        request.Headers["Digital-Signature"],
        nil,
    )

    if event.Subscription == "pix-request" {
        fmt.Println(event.Log)
    } else if event.Subscription == "pix-reversal" {
        fmt.Println(event.Log)
    } else if event.Subscription == "issuing-card" {
        fmt.Println(event.Log)
    } else if event.Subscription == "issuing-invoice" {
        fmt.Println(event.Log)
    } else if event.Subscription == "issuing-purchase" {
        fmt.Println(event.Log)
    } else if event.Subscription == "credit-note" {
        fmt.Println(event.Log)
    }
}

```

### Query webhook events

To search for webhooks events, run:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Event "github.com/starkinfra/sdk-go/starkinfra/event"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 15

    events := Event.Query(params, nil)
    for event := range events {
        fmt.Println(event.Id)
    }
}

```

### Get a webhook event

You can get a specific webhook event by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Event "github.com/starkinfra/sdk-go/starkinfra/event"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    event, err := Event.Get("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(event.Id)
}

```

### Delete a webhook event

You can also delete a specific webhook event by its id.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Event "github.com/starkinfra/sdk-go/starkinfra/event"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    event, err := Event.Delete("5792731695677440", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(event.Id)
}

```

### Set webhook events as delivered

This can be used in case you've lost events.
With this function, you can manually set events retrieved from the API as
"delivered" to help future event queries with `isDelivered=False`.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Event "github.com/starkinfra/sdk-go/starkinfra/event"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    event, err := Event.Update("4535785248260096", true, nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(event.Id)
}

```

### Query failed webhook event delivery attempts information

You can also get information on failed webhook event delivery attempts.

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Attempt "github.com/starkinfra/sdk-go/starkinfra/event/attempt"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    var params = map[string]interface{}{}
    params["limit"] = 2

    attempts := Attempt.Query(params, nil)
    for attempt := range attempts {
        fmt.Println(attempt.Id)
    }
}

```

### Get a failed webhook event delivery attempt information

To retrieve information on a single attempt, use the following function:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    Attempt "github.com/starkinfra/sdk-go/starkinfra/event/attempt"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    attempt, err := Attempt.Get("5182107395555328", nil)
    if err.Errors != nil {
        for _, e := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
        }
    }

    fmt.Println(attempt.Id)
}

```

# Handling errors

The SDK may return errors as the StarkErrors struct, which contains the "code" and "message" attributes.

It's highly recommended that you handle the errors returned from the functions used to get a feedback of the operation,
as the example below:

__InputError__ will be raised whenever the API detects an error in your request (status code 400). If you catch such an
error, you can get its elements to verify each of the individual errors that were detected in your request by the API.
For example:

```golang
package main

import (
    "fmt"
    "github.com/starkinfra/sdk-go/starkinfra"
    PixReversal "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
    "github.com/starkinfra/sdk-go/tests/utils"
)

func main() {

    starkinfra.User = utils.ExampleProject

    reversals, err := PixReversal.Create(
        []PixReversal.PixReversal{
            {
                Amount:     10000,
                EndToEndId: "E00000000202201060100rzsJzG9PzMg",
                ExternalId: "my_external_id",
                Reason:     "bankError",
            },
        }, nil)
    if err.Errors != nil {
        for _, erro := range err.Errors {
            panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
        }
    }
    
    for _, reversal := range reversals {
        fmt.Println(reversal)
    }
}

```

__InternalServerError__ will be raised if the API runs into an internal error.
If you ever stumble upon this one, rest assured that the development team
is already rushing in to fix the mistake and get you back up to speed.

__UnknownError__ will be raised if a request encounters an error that is
neither __InputErrors__ nor an __InternalServerError__, such as connectivity problems.

__InvalidSignatureError__ will be raised specifically by event.Parse()
when the provided content and signature do not check out with the Stark Bank public
key.

# Help and Feedback

If you have any questions about our SDK, just email us.
We will respond you quickly, pinky promise. We are here to help you integrate with us ASAP.
We also love feedback, so don't be shy about sharing your thoughts with us.

Email: help@starkinfra.com
