package satcomm

import (
	"errors"
)

// LaunchSatellites will launch the satellites from the given launchSites
func LaunchSatellites(totalSats, satsPerSite, timeToLaunch, launchSites int) error {
	if totalSats == 0 || satsPerSite == 0 || launchSites == 0 {
		return errors.New("Invalid input, totalSats, satsPerSite, and " +
			"launchSites can not be 0")
	}
	return nil
}
