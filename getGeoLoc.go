// find geo location via either postal address or via longitude and latitude
//
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	// PlaceData used to hold the JSON parsed data - use pointer instead?
	PlaceData Place
)

// Place structure used from JSON to find place from Long + Lat lookup
type Place struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	}
	Status           string `json:"status"`
	GoogleHttpStatus int    `json:"httpStatus,string"`
}

// To look up a town in UK:
// https://maps.googleapis.com/maps/api/geocode/json?address=Barry&region=uk&key=<KEY_HERE>

// Provide the API Key for geocoding requests
func getAPIKey() (apiGeoKey string, err error) {
	// get contents of environment variable $GAPI
	apiGeoKey = os.Getenv("GAPI")
	if apiGeoKey == "" {
		return "", fmt.Errorf("Google API key not available as environment variable: '$GAPI'")
	}

	// Add the DSAPI key to the weatherSetting structure for storing on exit
	weatherSetting.GoogleAPIKey = apiGeoKey

	return apiGeoKey, nil
}

// request the local town based on lat + long location and return
// https://maps.googleapis.com/maps/api/geocode/json?latlng=51.419212,-3.291481&result_type=locality&key=<KEY_HERE>
func getLocPlace(latLong string) string {

	apiKey, err := getAPIKey()

	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("Set environment variable for Google Places API key as: 'export GAPI=\"api_key_here\"'")
		fmt.Println("or on Windows as:  '$env:GAPI=\"add_api_key_here\"'")
		return "UNKNOWN"
	}
	if debugSwitch {
		fmt.Println("API Key is:", apiKey)
	}

	// check long lat exists
	if latLong == "" {
		fmt.Println("ERROR: missing latitude and longitude for geolocation look up.")
		runtime.Goexit()
	}

	// complete URL ready for call
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s&result_type=locality&key=%s", latLong, apiKey)

	if debugSwitch {
		fmt.Println("Final GeoLoc URL is:", url)
	}

	// configure the web request
	var myGoogleClient = &http.Client{Timeout: 10 * time.Second}

	// make the request to the DarkSky web site
	resp, err := myGoogleClient.Get(url)
	// exit app if web request errors
	if err != nil {
		fmt.Println("\nWARNING HTTP ERROR:\n", err)
		runtime.Goexit()
	}
	defer resp.Body.Close()

	// check the HTTP response code
	if debugSwitch {
		fmt.Println("DEBUG: HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// show all the http response headers
	if debugSwitch {
		fmt.Println("DEBUG: All Google Place API Received HTTP Headers: ")
		for k, v := range resp.Header {
			fmt.Print("\t", k)
			fmt.Print(" : ")
			fmt.Println(v)
		}
		fmt.Println("DEBUG HTTP Headers end\n")
	}

	// unmarshal the JSON data received into the 'Place' structure
	json.NewDecoder(resp.Body).Decode(&PlaceData)

	// save the DarkSky API requests made so far today for info
	//PlaceData.Current.DarkSkyReq = resp.Header.Get("X-Forecast-Api-Calls")
	// save the status code
	PlaceData.GoogleHttpStatus = resp.StatusCode

	// TODO: check for more than one returned value in array:
	//  PlaceData.Results[0].FormattedAddress

	// TODO: check for 'errors' in json returned data - ie not OK

	// direct output of parsed JSON values for debug only
	if debugSwitch {
		fmt.Println("\nDEBUG: Google API JSON parsed data values are:")
		fmt.Println("\tAddress is:", PlaceData.Results[0].FormattedAddress)
		fmt.Println("\tGoogle HTTP Status is:", PlaceData.GoogleHttpStatus)
		fmt.Println("\tJSON Request Status is:", PlaceData.Status)
		fmt.Println("DEBUG: JSON parsed data values end\n")
	}

	// return the place name
	return PlaceData.Results[0].FormattedAddress

}

// request the long + lat location based on town and country and return
// https://maps.googleapis.com/maps/api/geocode/json?address=sandringham+close+barry&region=uk&key=<KEY_HERE>
func getLocLongLat(locTown, locCountry string) string {

	apiKey, err := getAPIKey()

	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("Set environment variable for Google Places API key.")
		fmt.Println("For Linux use:\n\t\t'export GAPI=\"<api_key_here>\"'\n\n")
		runtime.Goexit()
	}
	if debugSwitch {
		fmt.Println("GAPI Key is:", apiKey)
	}
	return ""
}
