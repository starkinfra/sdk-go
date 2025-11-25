package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixStatement "github.com/starkinfra/sdk-go/starkinfra/pixstatement"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPixStatementPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	statement, err := PixStatement.Create(Example.PixStatement(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, statement.Id)
}

func TestPixStatementQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	statements, errorChannel := PixStatement.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case statement, ok := <-statements:
			if !ok {
				break loop
			}
			assert.NotNil(t, statement.Id)
		}
	}
}

func TestPixStatementPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	statements, cursor, err := PixStatement.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, statement := range statements {
		assert.NotNil(t, statement.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixStatementInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var statementList []PixStatement.PixStatement

	statements, errorChannel := PixStatement.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case statement, ok := <-statements:
			if !ok {
				break loop
			}
			statementList = append(statementList, statement)
		}
	}

	for _, statement := range statementList {
		getStatement, err := PixStatement.Get(statement.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getStatement.Id)
	}

	assert.Equal(t, limit, len(statementList))
}

func TestPixStatementCsv(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var statementList []PixStatement.PixStatement

	statements, errorChannel := PixStatement.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case statement, ok := <-statements:
			if !ok {
				break loop
			}
			statementList = append(statementList, statement)
		}
	}

	csv, err := PixStatement.Csv(statementList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.csv", "statement", statementList[0].Id)
	errFile := os.WriteFile(filename, csv, 0666)
	if errFile != nil {
		t.Errorf("error: %s", errFile.Error())
	}
	assert.NotNil(t, csv)
}
