//
//

package main

import (
	"fmt"
)

// URL format:  "https://api.darksky.net/forecast/APIKEY/LATITUDE,LONGITUDE,TIME?units=ca&lang=en"
// https://api.darksky.net/forecast/66fd639c6914180e12c355899c5ec267/51.419212,-3.291481?exclude=minutely,hourly?units=uk2

const BASEURL = "https://api.darksky.net/forecast/"

func GetURL(coord string) (url string, err error) {

	key := "66fd639c6914180e12c355899c5ec267"

	if coord == "" {
		coord = "51.419212,-3.291481"
	}

	// construct final url to obtain forecast
	url = BASEURL + key + "/" + coord + "?units=uk2&exclude=minutely,hourly"

	if debugSwitch {
		fmt.Println("DEBUG: final url is: ", url)
	}

	if url == "" {
		return "", fmt.Errorf("Unable to create DarkSky URL")
	}

	return url, nil
}
