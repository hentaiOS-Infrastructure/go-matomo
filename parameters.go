package matomo

// field descriptions are all from the Matomo docs as of 20210609: https://developer.matomo.org/api-reference/tracking-api

import (
	"fmt"
	"math/rand"
	"time"
)

// Parameters are the content that gets sent to the API. If the field is nil, it is skipped. If it isn't nil, it will be
// automatically encoded and added to the body of the request. sendImage will be set to false. We have a matomo tag, but
// aren't currently using it and just using hard coded values. Keep in mind that many of these fields are included for
// completeness sake and will not likely be known or relevant in a server-side context (eg: the user's resolution).
type Parameters struct {
	RecommendedParameters     *RecommendedParameters
	UserParameters            *UserParameters
	ActionParameters          *ActionParameters
	PagePerformanceParameters *PagePerformanceParameters
	EventTrackingParameters   *EventTrackingParameters
	ContentTrackingParameters *ContentTrackingParameters
	EcommerceParameters       *EcommerceParameters
	OtherParameters           *OtherParameters
}

// RecommendedParameters are the recommended parameters that really should be provided on each call if available
type RecommendedParameters struct {
	// Meant to hold a random value that is generated before each request. Using it helps avoid the tracking request being cached by the browser or a proxy. If not set, the SDK will set it for you.
	Rand *int64 `json:"rand" matomo:"rand"` // generated at call time if not provided
	// The parameter &apiv=1 defines the api version to use (currently always set to 1). The SDK sets this for you.
	APIV *int64 `json:"apiv" matomo:"apiv"` // always set to 1
	// Matomo will respond with a HTTP 204 response code instead of a GIF image.
	SendImage *int64 `json:"send_image" matomo:"send_image"`
}

type ActionParameters struct {
	// The title of the action being tracked. It is possible to use slashes / to set one or several categories for this action. For example, Help / Feedback will create the Action Feedback in the category Help.
	ActionName *string `json:"action_name" matomo:"action_name"`
	// The full URL for the current action.
	Url *string `json:"url" matomo:"url"`
	// The unique visitor ID, must be a 16 characters hexadecimal string. Every unique visitor must be assigned a different ID and this ID must not change after it is assigned. If this value is not set Matomo (formerly Piwik) will still track visits, but the unique visitors metric might be less accurate.
	VisitorID *string `json:"_id" matomo:"_id"`
	// URL of a file the user has downloaded. Used for tracking downloads.
	Download *string `json:"download" matomo:"download"`
	// An external URL the user has opened. Used for tracking outlink clicks.
	Link *string `json:"link" matomo:"link"`
}

// UserParameters are user specific parameters for the event
type UserParameters struct {
	// The full HTTP Referrer URL. This value is used to determine how someone got to your website (ie, through a website, search engine or campaign).
	URLRef *string `json:"urlref" matomo:"urlref"`
	//  Visit scope custom variables. This is a JSON encoded string of the custom variable array.
	CVar *string `json:"_cvar" matomo:"_cvar"`
	// The current count of visits for this visitor. To set this value correctly, it would be required to store the value for each visitor in your application (using sessions or persisting in a database). Then you would manually increment the counts by one on each new visit or "session", depending on how you choose to define a visit. This value is used to populate the report Visitors > Engagement > Visits by visit number.
	IDVC *int64 `json:"_idvc" matomo:"_idvc"`
	// The UNIX timestamp of this visitor's previous visit. This parameter is used to populate the report Visitors > Engagement > Visits by days since last visit.
	ViewTS *int64 `json:"_viewts" matomo:"_viewts"`
	// The UNIX timestamp of this visitor's first visit. This could be set to the date where the user first started using your software/app, or when he/she created an account. This parameter is used to populate the Goals > Days to Conversion report.
	IDTS *int64 `json:"_idts" matomo:"_idts"`
	// The Campaign name. Used to populate the Referrers > Campaigns report. Note: this parameter will only be used for the first pageview of a visit.
	CampaignName *string `json:"_rcn" matomo:"_rcn"`
	// The Campaign Keyword. Used to populate the Referrers > Campaigns report (clicking on a campaign loads all keywords for this campaign). Note: this parameter will only be used for the first pageview of a visit.
	CampaignKeyword *string `json:"_rck" matomo:"_rck"`
	//  The resolution of the device the visitor is using, eg 1280x1024.
	Resolution *string `json:"res" matomo:"res"`
	// The current hour (local time). The SDK will automatically set this if you don't.
	CurrentHour *string `json:"h" matomo:"h"`
	// The current minute (local time). The SDK will automatically set this if you don't.
	CurrentMinute *string `json:"m" matomo:"m"`
	// The current second (local time). The SDK will automatically set this if you don't.
	CurrentSecond *string `json:"s" matomo:"s"`
	// Various user plugins that the server likely won't know about.
	UserPlugins *UserPlugins `json:"plugins" matomo:"-"`
	// When set to 1, the visitor's client is known to support cookies.
	CookiesSupported *bool `json:"cookie" matomo:"cookie"`
	// An override value for the User-Agent HTTP header field. The user agent is used to detect the operating system and browser used.
	UserAgent *string `json:"ua" matomo:"ua"`
	// An override value for the Accept-Language HTTP header field. This value is used to detect the visitor's country if GeoIP is not enabled.
	Lang *string `json:"lang" matomo:"lang"`
	// Defines the User ID for this request. User ID is any non-empty unique string identifying the user (such as an email address or an username). To access this value, users must be logged-in in your system so you can fetch this user ID from your system, and pass it to Matomo. The User ID appears in the visits log, the Visitor profile, and you can Segment reports for one or several User ID (userId segment). When specified, the User ID will be "enforced". This means that if there is no recent visit with this User ID, a new one will be created. If a visit is found in the last 30 minutes with your specified User ID, then the new action will be recorded to this existing visit.
	UserID *string `json:"uid" matomo:"uid"`
	// If set to 1, will force a new visit to be created for this action. T
	NewVisit *bool `json:"new_visit" matomo:"new_visit"`
}

// UserPlugins is a sub-struct of capabilities for a user
type UserPlugins struct {
	Flash       *bool `json:"fla" matomo:"fla"`
	Java        *bool `json:"java" matomo:"java"`
	Director    *bool `json:"dir" matomo:"dir"`
	Quicktime   *bool `json:"qt" matomo:"qt"`
	RealPlayer  *bool `json:"realp" matomo:"realp"`
	PDF         *bool `json:"pdf" matomo:"pdf"`
	WMA         *bool `json:"wma" matomo:"wma"`
	Gears       *bool `json:"gears" matomo:"gears"`
	Silverlight *bool `json:"ag" matomo:"ag"`
}

type PagePerformanceParameters struct {
}

// EventTrackParameters add context to a user's actions on your platform.
type EventTrackingParameters struct {
	// The event category. Must not be empty. (eg. Videos, Music, Games...)
	Category *string `json:"e_c" matomo:"e_c"`
	// The event action. Must not be empty. (eg. Play, Pause, Duration, Add Playlist, Downloaded, Clicked...)
	Action *string `json:"e_a" matomo:"e_a"`
	// The event name. (eg. a Movie name, or Song name, or File name...)
	Name *string `json:"e_n" matomo:"e_n"`
	// The event value. Must be a float or integer value (numeric), not a string.
	Value *float64 `json:"e_v" matomo:"e_v"`
}

type ContentTrackingParameters struct {
	// The name of the content. For instance 'Ad Foo Bar'
	Name *string `json:"c_n" matomo:"c_n"`
	// The actual content piece. For instance the path to an image, video, audio, any text
	Piece *string `json:"c_p" matomo:"c_p"`
	// The target of the content. For instance the URL of a landing page
	Target *string `json:"c_t" matomo:"c_t"`
	// The name of the interaction with the content. For instance a 'click'
	Interaction *string `json:"c_i" matomo:"c_i"`
}

type EcommerceParameters struct {
}

type OtherParameters struct {
	// Override value for the visitor IP (both IPv4 and IPv6 notations supported).
	CIP *string `json:"cip" matomo:"cip"`
}

// StringPtr converts a static string to a pointer for use in the api
func StringPtr(input string) *string {
	return &input
}

// Int64Ptr converts a static int64 to a pointer for use in the api
func Int64Ptr(input int64) *int64 {
	return &input
}

// BoolPtr converts a static bool to a pointer for use in the api
func BoolPtr(input bool) *bool {
	return &input
}

// Float64Ptr converts a static float to a pointer for use in the api
func Float64Ptr(input float64) *float64 {
	return &input
}

//
// below, we set up all of the encoders for the structs to convert them into
// map[string]string for embedding in the URL

func (params *Parameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	// loop through the fields
	if params.RecommendedParameters != nil {
		subRet := params.RecommendedParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}
	if params.UserParameters != nil {
		subRet := params.UserParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}
	if params.EventTrackingParameters != nil {
		subRet := params.EventTrackingParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}
	if params.ContentTrackingParameters != nil {
		subRet := params.ContentTrackingParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}
	if params.ActionParameters != nil {
		subRet := params.ActionParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}
	if params.OtherParameters != nil {
		subRet := params.OtherParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}

	return ret
}

func (params *RecommendedParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	// set the required constants
	params.APIV = Int64Ptr(1)
	params.SendImage = Int64Ptr(0)
	if params.Rand == nil {
		params.Rand = Int64Ptr(rand.Int63n(99999999999999999))
	}
	// loop through the fields
	ret["apiv"] = fmt.Sprintf("%v", *params.APIV)
	ret["send_image"] = fmt.Sprintf("%v", *params.SendImage)
	ret["rand"] = fmt.Sprintf("%v", *params.Rand)

	return ret
}

func (params *UserParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	now := time.Now()
	// convert the fields
	if params.URLRef != nil {
		ret["urlref"] = *params.URLRef
	}
	if params.CVar != nil {
		ret["_cvar"] = *params.CVar
	}
	if params.IDVC != nil {
		ret["_idvc"] = fmt.Sprintf("%v", *params.IDVC)
	}
	if params.ViewTS != nil {
		ret["_viewts"] = fmt.Sprintf("%v", *params.ViewTS)
	}
	if params.IDTS != nil {
		ret["_idts"] = fmt.Sprintf("%v", *params.IDTS)
	}
	if params.CampaignName != nil {
		ret["_rcn"] = *params.CampaignName
	}
	if params.CampaignKeyword != nil {
		ret["_rck"] = *params.CampaignKeyword
	}
	if params.Resolution != nil {
		ret["res"] = *params.Resolution
	}
	if params.UserAgent != nil {
		ret["ua"] = *params.UserAgent
	}
	if params.CurrentHour != nil {
		ret["h"] = *params.CurrentHour
	} else {
		ret["h"] = now.Format("15")
	}
	if params.CurrentMinute != nil {
		ret["m"] = *params.CurrentMinute
	} else {
		ret["m"] = now.Format("04")
	}
	if params.CurrentSecond != nil {
		ret["s"] = *params.CurrentSecond
	} else {
		ret["s"] = now.Format("05")
	}
	if params.CookiesSupported != nil {
		if *params.CookiesSupported {
			ret["cookie"] = "1"
		} else {
			ret["cookie"] = "0"
		}
	}
	if params.Lang != nil {
		ret["lang"] = *params.Lang
	}
	if params.UserID != nil {
		ret["uid"] = *params.UserID
	}
	if params.NewVisit != nil {
		ret["new_visit"] = "1"
	}

	// now plugins
	plugins := params.UserPlugins.encode()
	for k, v := range plugins {
		ret[k] = v
	}

	return ret
}

func (params *UserPlugins) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	if params.Flash != nil {
		if *params.Flash {
			ret["fla"] = "1"
		} else {
			ret["fla"] = "0"
		}
	}
	if params.Java != nil {
		if *params.Java {
			ret["java"] = "1"
		} else {
			ret["java"] = "0"
		}
	}
	if params.Director != nil {
		if *params.Director {
			ret["dir"] = "1"
		} else {
			ret["dir"] = "0"
		}
	}
	if params.Quicktime != nil {
		if *params.Quicktime {
			ret["qt"] = "1"
		} else {
			ret["qt"] = "0"
		}
	}
	if params.RealPlayer != nil {
		if *params.RealPlayer {
			ret["realp"] = "1"
		} else {
			ret["realp"] = "0"
		}
	}
	if params.PDF != nil {
		if *params.PDF {
			ret["pdf"] = "1"
		} else {
			ret["pdf"] = "0"
		}
	}
	if params.WMA != nil {
		if *params.WMA {
			ret["wma"] = "1"
		} else {
			ret["wma"] = "0"
		}
	}
	if params.Gears != nil {
		if *params.Gears {
			ret["gears"] = "1"
		} else {
			ret["gears"] = "0"
		}
	}
	if params.Silverlight != nil {
		if *params.Silverlight {
			ret["ag"] = "1"
		} else {
			ret["ag"] = "0"
		}
	}

	return ret
}

func (params *EventTrackingParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	// both Action and Category are required
	if params.Action != nil && params.Category != nil {
		ret["e_c"] = *params.Category
		ret["e_a"] = *params.Action
	}
	if params.Name != nil {
		ret["e_n"] = *params.Name
	}
	if params.Value != nil {
		ret["e_v"] = fmt.Sprintf("%v", *params.Value)
	}

	return ret
}

func (params *ContentTrackingParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}

	if params.Name != nil {
		ret["c_n"] = *params.Name
	}

	if params.Piece != nil {
		ret["c_p"] = *params.Piece
	}

	if params.Interaction != nil {
		ret["c_i"] = *params.Interaction
	}

	if params.Target != nil {
		ret["c_t"] = *params.Target
	}

	return ret

}

func (params *ActionParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}

	if params.ActionName != nil {
		ret["action_name"] = *params.ActionName
	}
	if params.VisitorID != nil {
		ret["_id"] = *params.VisitorID
	}
	if params.Url != nil {
		ret["url"] = *params.Url
	}

	if params.Download != nil {
		ret["download"] = *params.Download
	}

	if params.Link != nil {
		ret["link"] = *params.Link
	}

	return ret
}

func (params *OtherParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	if config.AuthToken == "" {
		return ret
	}
	if params.CIP != nil {
		ret["cip"] = *params.CIP
	}

	return ret
}
