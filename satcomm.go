package satcomm

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// LaunchSatellites orchestrates and launches the satellites from the appropriate launch site
func LaunchSatellites(totalSats, satsPerSite, timeToLaunch, launchSites int) error {
	if totalSats == 0 || satsPerSite == 0 || launchSites == 0 {
		return errors.New("Invalid input, totalSats, satsPerSite, and " +
			"launchSites can not be 0")
	}
	currentSat := 1
	currentLS := 1
	var wg sync.WaitGroup
	for currentSat <= totalSats {
		count := 1
		// Loop used to launch the satsPerSite for each launch site
		for count <= satsPerSite && currentSat <= totalSats {
			wg.Add(1)
			go launchSatelite(currentSat, currentLS, timeToLaunch, &wg)
			count++
			currentSat++
		}
		// increment or reset launchSite after workers are done / satellites are launched
		wg.Wait()
		if currentLS == launchSites {
			currentLS = 1
		} else {
			currentLS++
		}
	}
	return nil
}

// launchSatellite launches an individual satellite after timeToLaunch milliseconds
func launchSatelite(satellite, launchSite, timeToLaunch int, wg *sync.WaitGroup) {
	fmt.Printf("Starting launch sequence for SAT%d from LS%d\n", satellite,
		launchSite)
	<-time.NewTimer(time.Duration(timeToLaunch) * time.Millisecond).C
	fmt.Printf("Launched SAT%d from LS%d\n", satellite, launchSite)
	wg.Done()
}
