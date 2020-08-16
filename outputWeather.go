//
//
//

package main

import (
	"os"
	"text/template"
)

// The weather forecast output to screen template
var weatherOutTemplate = `
                          WEATHER FORECAST
                    
» Request Status       : {{.Current.HTTPStatus}}
» Weather timezone     : {{.Tz}}
» Latitude & longitude : {{.Lat}} °N, longitude: {{.Long}} °N
» Weather Location     : {{.Place}}
» Forcast Time         : {{.Current.DisplayTime}}

∞∞ Forecast ∞∞

» Weather Currently:
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
	t := template.Must(template.New("weather").Parse(weatherOutTemplate))

	err = t.ExecuteTemplate(os.Stdout, "weather", ParsedData)

	if err != nil {
		return err
	}

	return nil
}
