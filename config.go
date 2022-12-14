package matomo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	Domain    string
	AuthToken string
	SiteID    string // if not provided, will be required in the call
	Rec       string // currently must always be set to 1
}

var config *Configuration

func Setup() {
	if config != nil {
		return
	}
	config = &Configuration{}
	config.Domain = strings.TrimSuffix(envHelper("MATOMO_DOMAIN", ""), "/")
	if config.Domain == "" {
		// TODO: convert to logger
		log.Printf("ERROR: MATOMO_DOMAIN was not set, so events will not be tracked")
		fmt.Fprintf(os.Stderr, "ERROR: MATOMO_DOMAIN was not set, so events will not be tracked\n")
	}
	// make sure they didn't put the matomo.php at the end
	config.Domain = strings.TrimSuffix(config.Domain, "matomo.php")
	config.SiteID = envHelper("MATOMO_SITE_ID", "")
	config.AuthToken = envHelper("MATOMO_AUTH_TOKEN", "")

	config.Rec = "1"

}

func envHelper(key, defaultValue string) string {
	found := os.Getenv(key)
	if found == "" {
		found = defaultValue
	}
	return found
}

func init() {
	Setup()
}
