//
//
//

package main

import (
	"fmt"
	"os"
	"text/template"
)

// myUsage function replaces the standard flag.Usage() function from Go. The
// function takes no paramaters, but outputs the command line flags
// that can be used when running the program.
func myUsage(appname string) {
	const usageTextTmpl = `
Purpose
¯¯¯¯¯¯¯
 Use '{{.appname}}' application to find the current weather forecast 
 information for the geographical planet earth location you provide.
 
Usage
¯¯¯¯¯
Run ./{{.appname}} with:

    Flag      Description                          Default Value
    ¯¯¯¯      ¯¯¯¯¯¯¯¯¯¯¯                          ¯¯¯¯¯¯¯¯¯¯¯¯¯
    -d        show debug output                    false
    -h        display help for this program        false
    -v        display program version              false
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
