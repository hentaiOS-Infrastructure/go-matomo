package matomo

import (
	"fmt"
	"math/rand"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAllEncoding(t *testing.T) {
	Setup()
	// start with just the recommended parameters
	params := Parameters{
		RecommendedParameters: &RecommendedParameters{
			Rand: Int64Ptr(rand.Int63n(9999999999999999)),
		},
	}

	encoded := params.encode()
	assert.NotNil(t, encoded)
	assert.NotNil(t, encoded["rand"])
	assert.NotZero(t, encoded["rand"])
	assert.Equal(t, "1", encoded["apiv"])
	assert.Equal(t, "0", encoded["send_image"])
	// make sure a few other params don't exist
	assert.Equal(t, 3, len(encoded)) // increment as other auto-populated fields are added
	assert.Empty(t, encoded["_rck"])

	encoded = testAllParams.encode()
	assert.Equal(t, len(encoded), len(encoded)) // this will increase as more fields are supported
}

func TestUserParameterEncoding(t *testing.T) {
	Setup()
	emptyUserParams := &UserParameters{}
	encoded := emptyUserParams.encode()
	// the times should be set to the current server time automatically
	assert.Equal(t, 3, len(encoded))
	// populate all the fields and encode
	encoded = testUserParams.encode()
	assert.Equal(t, len(encoded), len(encoded))

	assert.Equal(t, fmt.Sprintf("%d", *testUserParams.IDTS), encoded["_idts"])
	assert.Equal(t, fmt.Sprintf("%d", *testUserParams.ViewTS), encoded["_viewts"])
	assert.Equal(t, "1", encoded["_idvc"])
	assert.Equal(t, "Keyword Test", encoded["_rck"])
	assert.Equal(t, "Testing", encoded["_rcn"])
	assert.Equal(t, encoded["ua"], encoded["ua"])
	assert.Equal(t, "0", encoded["ag"])
	assert.Equal(t, "1", encoded["cookie"])
	assert.Equal(t, "0", encoded["dir"])
	assert.Equal(t, "0", encoded["fla"])
	assert.Equal(t, "0", encoded["gears"])
	assert.Equal(t, "0", encoded["java"])
	assert.Equal(t, "en-US", encoded["lang"])
	assert.Equal(t, "1", encoded["new_visit"])
	assert.Equal(t, "1", encoded["pdf"])
	assert.Equal(t, "0", encoded["qt"])
	assert.Equal(t, "0", encoded["realp"])
	assert.Equal(t, "1x1", encoded["res"])
	assert.Equal(t, "test-user", encoded["uid"])
	assert.Equal(t, "/users", encoded["urlref"])
	assert.Equal(t, "0", encoded["wma"])

}
func TestEventParameterEncodings(t *testing.T) {
	Setup()
	emptyEventParams := &EventTrackingParameters{}
	encoded := emptyEventParams.encode()
	assert.Equal(t, 0, len(encoded))
	// populate all the fields and encode
	encoded = testEventParams.encode()
	assert.Equal(t, url.QueryEscape(*testEventParams.Category), encoded["e_c"])
	assert.Equal(t, url.QueryEscape(*testEventParams.Action), encoded["e_a"])
	assert.Equal(t, url.QueryEscape(*testEventParams.Name), encoded["e_n"])
	assert.Equal(t, url.QueryEscape(fmt.Sprintf("%v", *testEventParams.Value)), encoded["e_v"])
}

func TestActionParametersEncodings(t *testing.T) {
	Setup()
	emptyActionParams := &ActionParameters{}
	encoded := emptyActionParams.encode()
	assert.Equal(t, 0, len(encoded))
	// populate all the fields and encode
	encoded = testActionParams.encode()
	assert.Equal(t, url.QueryEscape(*testActionParams.ActionName), encoded["action_name"])
	assert.Equal(t, *testActionParams.Url, encoded["url"])
	// assert.Equal(t, *testActionParams.Download, encoded["download"])
	assert.Equal(t, *testActionParams.Link, encoded["link"])
}

func TestContentParameterEncodings(t *testing.T) {
	Setup()
	emptyContentParams := &ContentTrackingParameters{}
	encoded := emptyContentParams.encode()
	assert.Equal(t, 0, len(encoded))
	// populate all the fields and encode
	encoded = testContentParams.encode()
	assert.Equal(t, *testContentParams.Target, encoded["c_t"])
	assert.Equal(t, url.QueryEscape(*testContentParams.Piece), encoded["c_p"])
	assert.Equal(t, url.QueryEscape(*testContentParams.Name), encoded["c_n"])
	assert.Equal(t, url.QueryEscape(*testContentParams.Interaction), encoded["c_i"])
}

var testAllParams = Parameters{
	RecommendedParameters:     &RecommendedParameters{},
	ActionParameters:          testActionParams,
	UserParameters:            testUserParams,
	PagePerformanceParameters: &PagePerformanceParameters{},
	// EventTrackingParameters:   testEventParams,
	ContentTrackingParameters: testContentParams,
	EcommerceParameters:       &EcommerceParameters{},
}

var testActionParams = &ActionParameters{
	ActionName: StringPtr("Hello from Among Us"),
	Url:        StringPtr("https://laboratory.luvvvvuratory.raphielscape.com/"),
	Download:   StringPtr("https://laboratory.luvvvvuratory.raphielscape.com/download.zip"),
	Link:       StringPtr("https://laboratory.luvvvvuratory.raphielscape.com/download.zip"),
}

var testUserParams = &UserParameters{
	URLRef:           StringPtr("/users"),
	CVar:             StringPtr("{\"id\": 1}"),
	IDVC:             Int64Ptr(1),
	ViewTS:           Int64Ptr(time.Now().Add(-1 * time.Minute).Unix()),
	IDTS:             Int64Ptr(time.Now().Add(-10 * time.Minute).Unix()),
	CampaignName:     StringPtr("Testing"),
	CampaignKeyword:  StringPtr("Keyword Test"),
	Resolution:       StringPtr("1x1"),
	CookiesSupported: BoolPtr(true),
	UserAgent:        StringPtr("Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0"),
	Lang:             StringPtr("en-US"),
	UserID:           StringPtr("test-user"),
	NewVisit:         BoolPtr(true),
	UserPlugins: &UserPlugins{
		Flash:       BoolPtr(false),
		Java:        BoolPtr(false),
		Director:    BoolPtr(false),
		Quicktime:   BoolPtr(false),
		RealPlayer:  BoolPtr(false),
		PDF:         BoolPtr(true),
		WMA:         BoolPtr(false),
		Gears:       BoolPtr(false),
		Silverlight: BoolPtr(false),
	},
}

var testEventParams = &EventTrackingParameters{
	Category: StringPtr("Event Category"),
	Action:   StringPtr("Event Action"),
	Name:     StringPtr("Event Name"),
	Value:    Float64Ptr(42.42),
}

var testContentParams = &ContentTrackingParameters{
	Name:        StringPtr("Content Name"),
	Piece:       StringPtr("download.zip"),
	Target:      StringPtr("https://laboratory.luvvvvuratory.raphielscape.com/download.zip"),
	Interaction: StringPtr("Event"),
}
