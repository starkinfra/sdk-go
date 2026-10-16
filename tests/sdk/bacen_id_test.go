package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func assertUtcMinute(t *testing.T, stamp string) {
	generated, err := time.Parse("200601021504", stamp)
	assert.Nil(t, err)
	assert.LessOrEqual(t, time.Since(generated).Abs(), 2*time.Minute)
}

func TestEndToEndIdUsesUtc(t *testing.T) {
	endToEndId := utils.EndToEndId("32160637")
	assert.Equal(t, 32, len(endToEndId))
	assertUtcMinute(t, endToEndId[9:21])
}

func TestReturnIdUsesUtc(t *testing.T) {
	returnId := utils.ReturnId("32160637")
	assert.Equal(t, 32, len(returnId))
	assertUtcMinute(t, returnId[9:21])
}
