//
//
package main

import (
	"fmt"
)

// Provide the API Key for geocoding requests
func getAPIKey() (apiGeoKey string) {
	return "AIzaSyCbPkXThwYoHgOgcJ_ZV89PYKyt213VpAw"
}

func getLocation() {

	apiKey := getAPIKey()
	fmt.Println("API Key is:", apiKey)
}
