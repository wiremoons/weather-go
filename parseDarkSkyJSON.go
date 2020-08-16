//
//
//

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var (
	// ParsedData used to hold the JSON parsed data - use pointer instead?
	ParsedData WeatherMain
)

// WeatherMain defines the main top level values to extract from JSON input
type WeatherMain struct {
	Tz         string     `json:"timezone"`
	Lat        float64    `json:"latitude"`
	Long       float64    `json:"longitude"`
	Place      string     // place via geo lookup
	LonBearing string     // added manually
	LatBearing string     // added manually
	Current    *Currently `json:"currently"`
	Daily      *Daily     `json:"daily"`
}

// Currently define sub values to extract from 'currently' JSON input
type Currently struct {
	Time        *UnixEpoch `json:"time"`
	DisplayTime string     `json:"displayTime,string"`
	Summary     string     `json:"summary"`
	Temp        float64    `json:"temperature"`
	FLTemp      float64    `json:"apparentTemperature"`
	WindSpd     float64    `json:"windSpeed"`
	UV          float64    `json:"uvIndex"`
	HTTPStatus  int        `json:"httpStatus,string"`
	DarkSkyReq  string     `json:"darkSkyApi,string"`
}

// Daily define sub values to extract from 'daily' JSON input
type Daily struct {
	DSummary string `json:"summary"`
}

// UnixEpoch timestamp that unmarshal from a UNIX timestamp.
type UnixEpoch struct {
	time.Time
}

///////////////////////////////////////////////////////////////////////////////
// functions start here
///////////////////////////////////////////////////////////////////////////////
func parseDarkSkyJSON(url string) {

	// configure the web request
	var myClient = &http.Client{Timeout: 10 * time.Second}

	// make the request to the DarkSky web site
	resp, err := myClient.Get(url)
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
		fmt.Println("DEBUG: All Received HTTP Headers: ")
		for k, v := range resp.Header {
			fmt.Print("\t", k)
			fmt.Print(" : ")
			fmt.Println(v)
		}
		fmt.Println("DEBUG HTTP Headers end\n")
	}

	// unmarshal the JSON data received into the 'ParseData' structure
	json.NewDecoder(resp.Body).Decode(&ParsedData)

	// Add additional data to the JSON ParsedData struct:
	//
	// Populate additional struct field with a formated date time
	ParsedData.Current.DisplayTime = ParsedData.Current.Time.Format("Monday 02 Jan 2006 at 15:04 (MST)")
	// save the DarkSky API requests made so far today for info
	ParsedData.Current.DarkSkyReq = resp.Header.Get("X-Forecast-Api-Calls")
	// save the status code
	ParsedData.Current.HTTPStatus = resp.StatusCode
	// use value obtained from place look up
	ParsedData.Place = weatherSetting.GeoLocation
	// add bearing direction for formatting for output
	if ParsedData.Lat < 0 {
		ParsedData.LatBearing = "°S"
	} else {
		ParsedData.LatBearing = "°N"
	}
	if ParsedData.Long < 0 {
		ParsedData.LonBearing = "°W"
	} else {
		ParsedData.LonBearing = "°E"
	}

	// direct output of parsed JSON values for debug only
	if debugSwitch {
		fmt.Println("\nDEBUG: JSON parsed data values are:")
		fmt.Println("\tTimezone is:", ParsedData.Tz)
		fmt.Println("\tLatitude is:", ParsedData.Lat)
		fmt.Println("\tLongitude is:", ParsedData.Long)
		fmt.Println("\tPlace is:", ParsedData.Place)
		fmt.Println("\tTime is:", ParsedData.Current.Time.Format("Monday 02 Jan 2006 at 15:04 (MST)"))
		fmt.Println("\tDisplay Time is:", ParsedData.Current.DisplayTime)
		fmt.Println("\tSummary is:", ParsedData.Current.Summary)
		fmt.Println("\tTemperature is:", ParsedData.Current.Temp)
		fmt.Println("\tFeels Like temperature is:", ParsedData.Current.FLTemp)
		fmt.Println("\tWind speed is:", ParsedData.Current.WindSpd)
		fmt.Println("\tUV Index is:", ParsedData.Current.UV)
		fmt.Println("\tDaily summary is:", ParsedData.Daily.DSummary, "\n")
		fmt.Println("\tHTTP Status code is:", ParsedData.Current.HTTPStatus)
		fmt.Println("\tNo. of DarkSky API requests is:", ParsedData.Current.DarkSkyReq)
		fmt.Println("DEBUG: JSON parsed data values end\n")
	}

	// output all JSON acquired data via function in 'outputWeather.go'
	// source file
	err = OutputForecast(ParsedData)
	if err != nil {
		fmt.Println(err)
	}

}

// UnmarshalJSON is a bespoke JSON unmarshal for Unix EPOCH time (ie 1470788940)
// See example found here: https://github.com/aws/aws-sdk-go/issues/796
func (t *UnixEpoch) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	// convert Epoch to normal time stamp
	t.Time = time.Unix(int64(ts), 0)
	if debugSwitch {
		fmt.Println("DEBUG EPOCH :", t.Format("Monday 02 Jan 2006 at 15:04 (MST)"))
	}
	return nil
}
