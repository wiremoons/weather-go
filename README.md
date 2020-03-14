
# Weather Forecast Command Line Tool

## Purpose

A command line tool to obatin the current weather forecast from a command prompt
or shell. 

Written in Go (Golong) so should compile and run on Windows, macOS, Linux, 
FreeBSD, etc. Program is being written and developed on a Raspberry Pi 4B from
an iPad Pro via Blink Shell.

Project has been created for personal use - so no big budget development team
is involved, so progress will be sporadic I expect... 😳

## Project Status

Early development. 

Manual entry of longitude and latitide cordintates will be required in 
the `getURL.go` file if you wish to play with the program.

## Build

Download (clone) the repo to your computer, and run: `make` or `go build` should 
work as well.

Dependancies to build include Go, and Make if prefered. Pre built binary versions 
 will be made available at a later date.

Further help to be added here as development progresses.


## Features

Current output (as of version 0.2.0 is:
```

                          WEATHER FORECAST
                    
» Request Status: 200
» Weather timezone is: 'Europe/London' at latitude: '51.419212' and longitude: '-3.291481'
» Weather Location is: 'TODO'

∞∞ Forecast ∞∞

» Time: Saturday 14 Mar 2020 at 22:42 (UTC)

» Weather Currenty:
        Summary:     Overcast
        Windspeed:   16.74
        Temperature: 9.47°C feels like: 6.12°C
        UV Index:    0

» General Outlook:
        'Light rain today through Friday.'

» Alerts:
        'TODO'

Weather forecast data: Powered by Dark Sky™
Visit: https://darksky.net/poweredby/
DarkSky API requests made: 44

All is well
```

Also output are available with command line flags: `-h` for *help* and 
`-v` to dispaly *version* information.


## License

Open source **MIT License** - see license file for detials,

