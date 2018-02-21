package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/itsHabib/satcomm"
)

func main() {
	satsToLaunch := flag.Int("sats", 500, "number of satellites to launch")
	timeToLaunch := flag.Int("ttl", 300, "number of milliseconds it takes to launch")
	launchSites := flag.Int("sites", 2, "number of launch sites")
	satsPerSite := flag.Int("persite", 4, "number of satellites each sight can launch at once")
	flag.Parse()

	err := satcomm.LaunchSatellites(*satsToLaunch, *satsPerSite, *timeToLaunch, *launchSites)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println("All satellites succesfully launched")
}
