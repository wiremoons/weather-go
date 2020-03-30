// Weather forecast retrieval application using DarkSky API
//
//

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	appversion = "0.3.0"
	appname    string
	// flag() variables CLI args
	debugSwitch bool
	helpMe      bool
	showVer     bool
	err         error
)

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
	// any clean up and ensures 'runtime.GoExit()' calls' work correctly
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

	// TODO: Add switch below as would be neater for runtime options?

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

	// TODO: default switch would run app from here
	//  if settings exists (ie not first run) then run
	//  otherwise run 'setup' to create settings

	// Obtain URL from function in 'GetURL.go' source file
	// TODO: get coords from settings first
	url, err := GetURL("51.419212,-3.291481")
	// exit app if url request errors
	if err != nil {
		fmt.Println("\nWARNING HTTP ERROR:\n", err)
		runtime.Goexit()
	}

	// get place name from: getGeoLoc.go
	myPlace := getLocPlace("51.419212,-3.291481")

	fmt.Println("PLACE IS:", myPlace)

	// call function in parseJSON.go to obtain weather data and output
	parseDarkSkyJSON(url)

	// END PROGRAM
}
