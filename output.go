//
//
//

package main

import (
	"fmt"
	"os"
	"runtime"
	"text/template"
)

// The weather forecast output to screen template
var weatherOutTmpl = `
                          WEATHER FORECAST
                    
» Request Status: {{.Current.HttpStatus}}
» Weather timezone is: '{{.Tz}}' at latitude: '{{.Lat}}' and longitude: '{{.Long}}'
» Weather Location is: 'TODO'

∞∞ Forecast ∞∞

» Time: {{.Current.DisplayTime}}

» Weather Currenty:
        Summary:     {{.Current.Summary}}
        Windspeed:   {{.Current.WindSpd}}
        Temperature: {{.Current.Temp}}°C feels like: {{.Current.FLTemp}}°C
        UV Index:    {{.Current.UV}}

» General Outlook:
        '{{.Daily.DSummary}}'

» Alerts:
        'TODO'

Weather forecast data: Powered by Dark Sky™
Visit: https://darksky.net/poweredby/
DarkSky API requests made: {{.Current.DarkSkyReq}}
`

// function to output the weather forecast data
//
func OutputForecast(ParsedData WeatherMain) (err error) {
	// check and build the template so the data field values are added
	// and the final output is displayed. Check for any error, and
	// abort if one is found.
	t := template.Must(template.New("weather").Parse(weatherOutTmpl))

	err = t.ExecuteTemplate(os.Stdout, "weather", ParsedData)

	if err != nil {
		return err
	}

	return nil
}

// versionInfo function collects details of the program being run and
// displays it on stdout
func versionInfo(appname string, appversion string) {
	// define a template for display on screen with place holders for data
	const appInfoTmpl = `
Running '{{.appname}}' version {{.appversion}}
 - Built with Go Complier '{{.compiler}}' on Golang version '{{.version}}'
 - Author's web site: http://www.wiremoons.com/
 - Source code for {{.appname}}: https://github.com/wiremoons/weather/
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
func myUsage(appname string) {
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
