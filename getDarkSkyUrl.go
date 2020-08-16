//
//

package main

import (
	"fmt"
	"os"
	"runtime"
)

const BASEURL = "https://api.darksky.net/forecast/"

// Provide the API Key for DarkSky API requests
func getDSAPIKey() (apiDSKey string, err error) {
	// get contents of environment variable $DSAPI
	apiDSKey = os.Getenv("DSAPI")
	if apiDSKey == "" {
		return "", fmt.Errorf("DarkSky API key not available as environment variable: '$GAPI'")
	}
	return apiDSKey, nil
}

// Example format for DarkSky API call:
// https://api.darksky.net/forecast/<API_KEY>/51.419212,-3.291481?exclude=minutely,hourly?units=uk2
//
func GetURL(coord string) (url string, err error) {

	// get the DarkSky API key from environment variable 'DSAPI'
	key, err := getDSAPIKey()
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("Set environment variable for DarkSky API key.")
		fmt.Println("For Linux use:\n\t\t'export DSAPI=\"<api_key_here>\"'\n\n")
		runtime.Goexit()
	}
	if debugSwitch {
		fmt.Println("DSAPI Key is:", key)
	}

	// Add the DSAPI key to the weatherSetting structure for storing on exit
	weatherSetting.DarkSkyAPIKey = key

	if coord == "" {
		coord = "51.419212,-3.291481"
	}

	// construct final url to obtain forecast
	url = BASEURL + key + "/" + coord + "?units=uk2&exclude=minutely,hourly"

	if debugSwitch {
		fmt.Println("DEBUG: final url is: ", url)
	}

	if url == "" {
		return "", fmt.Errorf("Unable to create DarkSky URL")
	}

	return url, nil
}
