//
//
//

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"
	"time"
)

// SET GLOBAL VARIABLES
// set the version of the app and prep var to hold app name
var appversion = "0.1.0"
var appname string

// flag() variables used for command line args
var debugSwitch bool
var helpMe bool
var showVer bool

// used to hold any errors
var err error

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
	Time    *UnixEpoch `json:"time"`
	Summary string     `json:"summary"`
	Temp    float64    `json:"temperature"`
	FLTemp  float64    `json:"apparentTemperature"`
	WindSpd float64    `json:"windSpeed"`
	UV      float64    `json:"uvIndex"`
}

// define sub values to extract from 'daily' JSON input
type Daily struct {
	DSummary string `json:"summary"`
}

//A Time is a time that unmarshals from a UNIX timestamp.
type UnixEpoch struct {
	time.Time
}

// init() always runs before the applications main() function and is
// used here to set-up the flag() variables from the command line
// parameters - which are provided by the user when they run the app.
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

// The weather forecast output to screen template
var weatherOutTmpl = `

∞∞ Forecast ∞∞

» Weather timezone is: '{{.Tz}}' at latitude: '{{.Lat}}' and longitude: '{{.Long}}'
» Time: {{.Current.Time}}

» Weather Currenty:
        Summary:     {{.Current.Summary}}
        Windspeed:   {{.Current.WindSpd}}
        Temperature: {{.Current.Temp}}°C feels like: {{.Current.FLTemp}}°C
        UV Index:    {{.Current.UV}}

» General Outlook:
        '{{.Daily.DSummary}}'

Weather forecast data: Powered by Dark Sky™
Visit: https://darksky.net/poweredby/
`

// main starts here
func main() {

	// Setup function that will run at the end of the program to perform
	// any clean up and esnures 'runtime.GoExit()' calls' work correctly
	defer func() {
		// END OF MAIN()
		fmt.Printf("\nAll is well\n")
		os.Exit(0)
	}()

	// confirm if debug mode is enabled and display other command line
	// flags and their current status
	if debugSwitch {
		fmt.Println("DEBUG: Debug mode enabled")
		fmt.Printf("DEBUG: Number of command line arguments set by user is: %d\n", flag.NFlag())
		fmt.Println("DEBUG: Command line argument settings are:")
		fmt.Println("\t\tDisplay additional debug output when run:", strconv.FormatBool(debugSwitch))
		fmt.Println("\t\tDisplay additional help information:", strconv.FormatBool(helpMe))
		fmt.Println("\t\tShow the applications version:", strconv.FormatBool(showVer))
	}

	// override Go standard flag.Usage() function to get better
	// formating and output by using my own function instead
	flag.Usage = func() {
		if debugSwitch {
			fmt.Println("DEBUG: Running flag.Usage override function")
		}
		myUsage()
	}

	// check if the use just wanted some help?
	if helpMe {
		flag.Usage()
		runtime.Goexit()
	}
	// check if the user just wanted the app version info?
	if showVer {
		versionInfo()
		runtime.Goexit()
	}

	// test JSON parse against weatherdata.json local file
	fileData, err := ioutil.ReadFile("test/weatherdata.json")
	if err != nil {
		fmt.Println("FATAL ERROR: ", err)
		// end here - note any app 'defer' will not happen
		os.Exit(-1)
	}

	var parsedData WeatherMain

	err = json.Unmarshal([]byte(fileData), &parsedData)

	if err != nil {
		fmt.Println(err)
	}

	// direct output of parsed JSON values for debug only
	if debugSwitch {
		fmt.Println("\nDEBUG: JSON parsed data values are:")
		fmt.Println("\tTimezone is:", parsedData.Tz)
		fmt.Println("\tLatitude is:", parsedData.Lat)
		fmt.Println("\tLongitude is:", parsedData.Long)
		fmt.Println("\tTime is:", parsedData.Current.Time.Format("Monday 02 Jan 2006 at 15:04 (MST)"))
		fmt.Println("\tSummary is:", parsedData.Current.Summary)
		fmt.Println("\tTemperature is:", parsedData.Current.Temp)
		fmt.Println("\tFeels Like temperature is:", parsedData.Current.FLTemp)
		fmt.Println("\tWind speed is:", parsedData.Current.WindSpd)
		fmt.Println("\tUV Index is:", parsedData.Current.UV)
		fmt.Println("\tDaily summary is:", parsedData.Daily.DSummary)
	}

	// check and build the template so the data field values are added
	// and the final output is displayed. Check for any error, and
	// abort if one is found.
	t := template.Must(template.New("weather").Parse(weatherOutTmpl))

	err = t.ExecuteTemplate(os.Stdout, "weather", parsedData)
	if err != nil {
		fmt.Println("\nFATAL ERROR: ", err)
		// end here - note any app 'defer' will not happen
		os.Exit(-2)
	}

}

// versionInfo function collects details of the program being run and
// displays it on stdout
func versionInfo() {
	// define a template for display on screen with place holders for data
	const appInfoTmpl = `
Running '{{.appname}}' version {{.appversion}}
 - Built with Go Complier '{{.compiler}}' on Golang version '{{.version}}'
 - Author's web site: http://www.wiremoons.com/
 - Source code for {{.appname}}: https://github.com/wiremoons/amt/
`
	// build a map with keys set to match the template names used
	// and the data fields to be used in the template as values
	verData := map[string]interface{}{
		"appname":    appname,
		"appversion": appversion,
		"compiler":   runtime.Compiler,
		"version":    runtime.Version(),
	}
	// check and build the template so the data field values are added
	// and the final output is displayed. Check for any error, and
	// abort if one is found.
	t := template.Must(template.New("appinf").Parse(appInfoTmpl))

	if err := t.Execute(os.Stdout, verData); err != nil {
		fmt.Printf("FATAL ERROR: in function 'versionInfo()' when building template with err: %v", err)
	}
}

// myUsage function replaces the standard flag.Usage() function from Go. The
// function takes no paramaters, but outputs the command line flags
// that can be used when running the program.
func myUsage() {
	const usageTextTmpl = `
Usage of ./{{.appname}}:
        Flag               Description                                        Default Value
        ¯¯¯¯               ¯¯¯¯¯¯¯¯¯¯¯                                        ¯¯¯¯¯¯¯¯¯¯¯¯¯
        -d                 show debug output                                  false
        -h                 display help for this program                      false
        -v                 display program version                            false
 `
	// build a map with keys set to match the template names used
	// and the data fields to be used in the template as values
	useData := map[string]interface{}{
		"appname": appname,
	}
	// check and build the template so the data field values are added
	// and the final output is displayed. Check for any error, and
	// abort if one is found.
	t := template.Must(template.New("useinf").Parse(usageTextTmpl))

	if err := t.Execute(os.Stdout, useData); err != nil {
		fmt.Printf("FATAL ERROR: in function 'myUsage()' when building template with err: %v", err)
	}

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
		fmt.Println("DEBUG :", t.Format("Monday 02 Jan 2006 at 15:04 (MST)"))
	}
	return nil
}

// change time format to our prefered format
// const longForm = "Monday 02 Jan 2006 at 15:04 (MST)"
// t.Time, err = time.Parse(t.Time, longForm)
