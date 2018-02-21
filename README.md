#satcomm

##Overview
satcomm simulates launch sites simultaneously launching satellites by using Go
concurrency constructs. The command line interface allows the user to define
flags to control the number of satellites launched, the time it takes for 
each to launch, the number of satellites per launch site that can be
launched at once, and the number of launch sites.

##Install
1. Clone repository: 
 - `git clone https://github.com/itsHabib/satcomm`
2. Navigate into command line interface directory:
 - `cd satcomm/cmd/satcomm`
3. Build command line application:
 - `go build`

##Usage
Use --help to see what flags are available and what the defaults are for each flag:

- `satcomm --help`

Ex. changing the number of satellites to launch to 100:

- `satcomm -sats=100`
