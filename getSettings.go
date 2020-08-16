//
//
//  Get settings to run the app or if none exist offer to create them!

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var (
	// path and filename for weather configuration settings
	weatherConfig  string = ""
	weatherSetting WeatherSettings
)

// WeatherSettings structure to hold the configuration and settings for the weather app
// stored and loaded from JSON setting file in function below
type WeatherSettings struct {
	GoogleAPIKey  string  `json:"googleapikey"`
	DarkSkyAPIKey string  `json:"darkskyapikey"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	LatLong       string  `json:"latlong"`
	GeoLocation   string  `json:"geolocation"`
}

// obtain the setting to run
func getSettings() (err error) {

	// get the configuration file and path based on OS
	weatherConfig = getConfigFile()

	// either load the settings or get user to provide them
	if settingsExist(weatherConfig) {
		// load the weather setting into the struct 'WeatherSettings'
		err = loadWeatherSetting(weatherConfig)
		if err != nil {
			fmt.Println("ERROR loading settings: ", err)
			return err
		}
		// done - setting loaded to struct
		return nil
	}

	// no existing settings file - so create some default settings...
	requestUserSettingInput()
	return nil

}

// obtain the setting to run
func saveSettings() (err error) {

	// get the configuration file and path based on OS
	weatherConfig = getConfigFile()
	// create the config file directory if needed
	chkPath := filepath.Dir(weatherConfig)
	os.MkdirAll(chkPath, os.ModePerm)

	// save the weather setting into the struct 'WeatherSettings'
	err = saveWeatherSetting(weatherConfig)
	if err != nil {
		fmt.Println("ERROR saving settings: ", err)
		return err
	}
	return nil

}

// ask the user to input some settings to create an initial config
func requestUserSettingInput() {
	fmt.Println("Initial settings required for weather forecast area")
	// set some defaults in the config file - as none exist yet...
	weatherSetting.GoogleAPIKey = ""
	weatherSetting.DarkSkyAPIKey = ""
	weatherSetting.Latitude = 51.419212
	weatherSetting.Longitude = -3.291481
	weatherSetting.LatLong = "51.419212,-3.291481"
	weatherSetting.GeoLocation = "Barry. Wales"
	return
}

// depending on operating system set the configuration file location
func getConfigFile() (weatherConfig string) {

	if runtime.GOOS == "windows" {
		weatherConfig = os.Getenv("APPDATA") + "/wm-weather/config.json"
	} else {
		weatherConfig = os.Getenv("HOME") + "/.config/wm-weather/config.json"
	}

	if debugSwitch {
		fmt.Printf("DEBUG: baseline weather config file assumed as: '%s'\n", weatherConfig)
	}

	return weatherConfig
}

// check if configuration settings file already exists?
func settingsExist(weatherConfig string) bool {
	if _, err := os.Stat(weatherConfig); os.IsNotExist(err) {
		return false
	}
	return true
}

// load the weather setting into the struct 'WeatherSettings'
func loadWeatherSetting(weatherConfig string) (err error) {
	jsonConfig, err := ioutil.ReadFile(weatherConfig)
	if err != nil {
		fmt.Println("ERROR when reading settings JSON config file :", err)
		return err
	}
	err = json.Unmarshal(jsonConfig, &weatherSetting)
	if err != nil {
		fmt.Println("ERROR when parsing JSON config file :", err)
		return err
	}
	return nil
}

// save the weather setting from the struct 'WeatherSettings'
func saveWeatherSetting(weatherConfig string) (err error) {

	jsonConfig, err := json.Marshal(weatherSetting)
	if err != nil {
		fmt.Println("ERROR when marshaling settings to JSON :", err)
		return err
	}

	err = ioutil.WriteFile(weatherConfig, []byte(jsonConfig), 0664)
	if err != nil {
		fmt.Println("ERROR saving the JSON data to the settings file :", err)
		return err
	}
	return nil
}
