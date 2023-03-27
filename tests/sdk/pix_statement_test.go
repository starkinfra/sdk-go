package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestPixStatementPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	statement, err := PixStatement.Create(Example.PixStatement(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, statement.Id)
	fmt.Println(statement.Id)
}

func TestPixStatementQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	statements := PixStatement.Query(params, nil)
	for statement := range statements {
		assert.NotNil(t, statement.Id)
		fmt.Println(statement.Id)
	}
}

func TestPixStatementPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	statements, cursor, err := PixStatement.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, statement := range statements {
		assert.NotNil(t, statement.Id)
		fmt.Println(statement.Id)
	}
	fmt.Println(cursor)
}

func TestPixStatementInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var statementList []PixStatement.PixStatement
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	statements := PixStatement.Query(paramsQuery, nil)
	for statement := range statements {
		statementList = append(statementList, statement)
	}

	statement, err := PixStatement.Get(statementList[rand.Intn(len(statementList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, statement.Id)
	fmt.Println(statement.Id)
}

func TestPixStatementCsv(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var statementList []PixStatement.PixStatement
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	statements := PixStatement.Query(paramsQuery, nil)
	for statement := range statements {
		statementList = append(statementList, statement)
	}

	csv, err := PixStatement.Csv(statementList[rand.Intn(len(statementList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.csv", "statement", statementList[(rand.Intn(len(statementList)))].Id)
	errFile := ioutil.WriteFile(filename, csv, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, csv)
}
