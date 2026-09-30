package sdk

import (
	"regexp"
	"testing"
	"time"

	Utils "github.com/starkinfra/sdk-go/starkinfra/utils"
	"github.com/stretchr/testify/assert"
)

func TestPixSubscriptionBacenId(t *testing.T) {

	bacenId := Utils.PixSubscriptionBacenId("32160637", "RR")

	assert.Len(t, bacenId, 29)
	assert.Regexp(t, regexp.MustCompile(`^RR32160637\d{8}[a-zA-Z0-9]{11}$`), bacenId)
	assert.Equal(t, time.Now().Format("20060102"), bacenId[10:18])
}

func TestPixSubscriptionBacenIdRandomPartDiffers(t *testing.T) {

	first := Utils.PixSubscriptionBacenId("32160637", "RR")
	second := Utils.PixSubscriptionBacenId("32160637", "RR")

	assert.NotEqual(t, first, second)
}

func TestEndToEndIdAndReturnIdKeepMinutePrecision(t *testing.T) {

	endToEndId := Utils.EndToEndId("32160637")
	returnId := Utils.ReturnId("32160637")

	assert.Len(t, endToEndId, 32)
	assert.Len(t, returnId, 32)
	assert.Regexp(t, regexp.MustCompile(`^E32160637\d{12}[a-zA-Z0-9]{11}$`), endToEndId)
	assert.Regexp(t, regexp.MustCompile(`^D32160637\d{12}[a-zA-Z0-9]{11}$`), returnId)
}
