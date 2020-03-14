//
//

package main

// URL format:  "https://api.darksky.net/forecast/APIKEY/LATITUDE,LONGITUDE,TIME?units=ca&lang=en"
const BASEURL = "https://api.darksky.net/forecast/"

//https://api.darksky.net/forecast/66fd639c6914180e12c355899c5ec267/51.419212,-3.291481?units=uk2&exclude=minutely,hourly
//https://api.darksky.net/forecast/66fd639c6914180e12c355899c5ec267/51.419212,-3.291481?exclude=minutely,hourly?units=uk2

func GetURL() (URL string, err error) {

	key := "66fd639c6914180e12c355899c5ec267"
	coord := "51.419212,-3.291481"

	// construct final url to obtain forecast
	URL = BASEURL + key + "/" + coord + "?units=uk2&exclude=minutely,hourly"

	// if debugSwith {
	// fmt.Println("DEBUG: final url is: ", url)
	// }

	return URL, nil
}
