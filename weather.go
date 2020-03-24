//
//
//

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	appversion = "0.2.0"
	appname    string
	// flag() variables CLI args
	debugSwitch bool
	helpMe      bool
	showVer     bool
	err         error
	// used to hold the JSON parsed data - use pointer instead?
	ParsedData WeatherMain
)

// define main top level values to extract from JSON input
type WeatherMain struct {
	Tz      string     `json:"timezone"`
	Lat     float64    `json:"latitude"`
	Long    float64    `json:"longitude"`
	Current *Currently `json:"currently"`
	Daily   *Daily     `json:"daily"`
}

// define sub values to extract from 'currently' JSON input
type Currently struct {
	Time        *UnixEpoch `json:"time"`
	DisplayTime string     `json:"displayTime,string"`
	Summary     string     `json:"summary"`
	Temp        float64    `json:"temperature"`
	FLTemp      float64    `json:"apparentTemperature"`
	WindSpd     float64    `json:"windSpeed"`
	UV          float64    `json:"uvIndex"`
	HttpStatus  int        `json:"httpStatus,string"`
	DarkSkyReq  string     `json:"darkSkyApi,string"`
}

// define sub values to extract from 'daily' JSON input
type Daily struct {
	DSummary string `json:"summary"`
}

// A Time is a time that unmarshals from a UNIX timestamp.
type UnixEpoch struct {
	time.Time
}

// init() always runs before the applications main() function
func init() {
	// flag types available are: IntVar; StringVar; BoolVar
	// flag parameters are: variable; cmd line flag; initial value; description.
	// 'description' is used by flag.Usage() on error or for help output
	flag.BoolVar(&debugSwitch, "d", false, "\tshow debug output")
	flag.BoolVar(&helpMe, "h", false, "\tdisplay help for this program")
	flag.BoolVar(&showVer, "v", false, "\tdisplay program version")
	// get the command line args passed to the program
	flag.Parse()
	// get the name of the application as called from the command line
	appname = filepath.Base(os.Args[0])
}

///////////////////////////////////////////////////////////////////////////////
// main starts here
///////////////////////////////////////////////////////////////////////////////
func main() {

	// Setup function that will run at the end of the program to perform
	// any clean up and esnures 'runtime.GoExit()' calls' work correctly
	defer func() {
		// END OF MAIN()
		fmt.Printf("\nAll is well\n")
		os.Exit(0)
	}()

	// if debug mode is enabled then display other command line
	// flag settings and their current status
	if debugSwitch {
		fmt.Println("DEBUG: Debug mode enabled")
		fmt.Printf("DEBUG: Total CLI arguments: %d\n", flag.NFlag())
		fmt.Println("DEBUG: CLI arguments set:")
		fmt.Println("\t\tShow 'debug' output:", strconv.FormatBool(debugSwitch))
		fmt.Println("\t\tShow help:", strconv.FormatBool(helpMe))
		fmt.Println("\t\tShow version:", strconv.FormatBool(showVer))
	}

	// override Go standard flag.Usage() function
	flag.Usage = func() {
		if debugSwitch {
			fmt.Println("DEBUG: Running flag.Usage override function")
		}
		myUsage(appname)
	}

	// check if the user requested help?
	if helpMe {
		flag.Usage()
		runtime.Goexit()
	}
	// check if the user requested app version?
	if showVer {
		versionInfo(appname, appversion)
		runtime.Goexit()
	}

	// Obtain URL from function in 'GetURL.go' source file
	// TODO: get coords from settings first
	url, err := GetURL("51.419212,-3.291481")
	// exit app if url request errors
	if err != nil {
		fmt.Println("\nWARNING HTTP ERROR:\n", err)
		runtime.Goexit()
	}

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

	// check the HTTP reponse code
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

	// unmarshall the JSON data received into the 'ParseData' structs
	json.NewDecoder(resp.Body).Decode(&ParsedData)

	// Add addtional data to the JSON ParsedData struct:
	//
	// Populate additional struct field with a formated date time
	ParsedData.Current.DisplayTime = ParsedData.Current.Time.Format("Monday 02 Jan 2006 at 15:04 (MST)")
	// save the DarkSky API requests made so far today for info
	ParsedData.Current.DarkSkyReq = resp.Header.Get("X-Forecast-Api-Calls")
	// save the status code
	ParsedData.Current.HttpStatus = resp.StatusCode

	// direct output of parsed JSON values for debug only
	if debugSwitch {
		fmt.Println("\nDEBUG: JSON parsed data values are:")
		fmt.Println("\tTimezone is:", ParsedData.Tz)
		fmt.Println("\tLatitude is:", ParsedData.Lat)
		fmt.Println("\tLongitude is:", ParsedData.Long)
		fmt.Println("\tTime is:", ParsedData.Current.Time.Format("Monday 02 Jan 2006 at 15:04 (MST)"))
		fmt.Println("\tDisplay Time is:", ParsedData.Current.DisplayTime)
		fmt.Println("\tSummary is:", ParsedData.Current.Summary)
		fmt.Println("\tTemperature is:", ParsedData.Current.Temp)
		fmt.Println("\tFeels Like temperature is:", ParsedData.Current.FLTemp)
		fmt.Println("\tWind speed is:", ParsedData.Current.WindSpd)
		fmt.Println("\tUV Index is:", ParsedData.Current.UV)
		fmt.Println("\tDaily summary is:", ParsedData.Daily.DSummary)
		fmt.Println("DEBUG: JSON parsed data values end\n")
	}

	// output all JSON aquired data via function in 'output.go' source file
	err = OutputForecast(ParsedData)
	if err != nil {
		fmt.Println(err)
	}

	// END PROGRAM
}

// Bespoke JSON unmarshall for Unix EPOCH time (ie 1470788940)
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
