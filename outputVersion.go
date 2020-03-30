//
//
//

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
)

// versionInfo function collects details of the program being run and
// displays it on stdout
func versionInfo(appname string, appversion string) {
	// define a template for display on screen with place holders for data
	const appInfoTmpl = `
Running '{{.appname}}' version {{.appversion}}.
 - Built with Go Complier '{{.compiler}}' on Golang version '{{.version}}'.
 - Executing on operating system '{{.operSys}}' on CPU architecture '{{.runArch}}'.

Copyright © 2020 Simon Rowe.

For licenses and further information visit:
 - Source code for {{.appname}}: https://github.com/wiremoons/weather/
 - Go language and compiler: https://golang.org/ 
 - DarkSky atrributions: https://darksky.net/attribution/
`
	// build a map with keys set to match the template names used
	// and the data fields to be used in the template as values
	verData := map[string]interface{}{
		"appname":    appname,
		"appversion": appversion,
		"operSys":    strings.Title(runtime.GOOS),
		"runArch":    strings.ToUpper(runtime.GOARCH),
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
