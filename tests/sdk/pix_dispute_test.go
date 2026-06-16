package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixDispute "github.com/starkinfra/sdk-go/starkinfra/pixdispute"
	PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func deriveSentReferenceId(t *testing.T) (string, bool) {
	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "success"

	requests, errorChannel := PixRequest.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				return "", false
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			if request.Flow == "out" && request.EndToEndId != "" {
				return request.EndToEndId, true
			}
		}
	}

	return "", false
}

func TestPixDisputePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	referenceId, ok := deriveSentReferenceId(t)
	if !ok {
		t.Skip("no successful PixRequest in sandbox to derive a reference_id; skipping create happy-path")
	}

	disputes, err := PixDispute.Create([]PixDispute.PixDispute{
		{
			ReferenceId:   referenceId,
			Method:        "scam",
			OperatorEmail: "ned.stark@company.com",
			OperatorPhone: "+5511999999999",
			Tags:          []string{"tony", "stark"},
		},
	}, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidDispute" || e.Code == "invalidReferenceId" || e.Code == "repeatedReferenceId" {
				t.Skipf("no fresh eligible transaction in sandbox (%s); skipping create happy-path", e.Code)
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	for _, dispute := range disputes {
		assert.NotNil(t, dispute.Id)
	}
}

func TestPixDisputePostMethodOther(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	referenceId, ok := deriveSentReferenceId(t)
	if !ok {
		t.Skip("no successful PixRequest in sandbox to derive a reference_id; skipping create-other happy-path")
	}

	disputes, err := PixDispute.Create([]PixDispute.PixDispute{
		{
			ReferenceId:   referenceId,
			Method:        "other",
			OperatorEmail: "ned.stark@company.com",
			OperatorPhone: "+5511999999999",
			Description:   "testDisputeGolang",
			Tags:          []string{"tony", "stark"},
		},
	}, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidDispute" || e.Code == "invalidReferenceId" || e.Code == "repeatedReferenceId" {
				t.Skipf("no fresh eligible transaction in sandbox (%s); skipping create-other happy-path", e.Code)
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	for _, dispute := range disputes {
		assert.NotNil(t, dispute.Id)
		assert.Equal(t, "other", dispute.Method)
		assert.NotEmpty(t, dispute.Description)
	}
}

func TestPixDisputeQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	disputes, errorChannel := PixDispute.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case dispute, ok := <-disputes:
			if !ok {
				break loop
			}
			assert.NotNil(t, dispute.Id)
		}
	}
}

func TestPixDisputePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	disputes, cursor, err := PixDispute.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, dispute := range disputes {
		assert.NotNil(t, dispute.Id)
	}

	assert.NotNil(t, cursor)
}

func TestPixDisputeInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var disputeList []PixDispute.PixDispute

	disputes, errorChannel := PixDispute.Query(paramsQuery, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case dispute, ok := <-disputes:
			if !ok {
				break loop
			}
			disputeList = append(disputeList, dispute)
		}
	}

	for _, dispute := range disputeList {
		getDispute, err := PixDispute.Get(dispute.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getDispute.Id)
		assert.NotEmpty(t, getDispute.Status)
		assert.NotEmpty(t, getDispute.Flow)
		assert.NotNil(t, getDispute.Created)
	}

	assert.LessOrEqual(t, len(disputeList), limit)
}

func TestPixDisputeCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var disputeList []PixDispute.PixDispute

	disputes, errorChannel := PixDispute.Query(paramsQuery, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case dispute, ok := <-disputes:
			if !ok {
				break loop
			}
			disputeList = append(disputeList, dispute)
		}
	}

	if len(disputeList) == 0 {
		t.Skip("no PixDispute in sandbox; skipping cancel happy-path")
	}

	canceled, err := PixDispute.Cancel(disputeList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidCancellationStatus" {
				t.Skip("no cancelable PixDispute in sandbox (entity in non-cancelable status); skipping cancel happy-path")
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	assert.NotNil(t, canceled.Id)
	assert.Equal(t, "canceled", canceled.Status)
}

func TestPixDisputeTransactionFields(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var disputeList []PixDispute.PixDispute

	disputes, errorChannel := PixDispute.Query(paramsQuery, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case dispute, ok := <-disputes:
			if !ok {
				break loop
			}
			disputeList = append(disputeList, dispute)
		}
	}

	for _, dispute := range disputeList {
		for _, transaction := range dispute.Transactions {
			assert.NotNil(t, transaction.EndToEndId)
		}
	}
}
