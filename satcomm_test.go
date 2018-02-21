package satcomm

import (
	"strings"
	"testing"
)

func TestLaunchSatellites(t *testing.T) {
	t.Run("Testing invalid input", func(t *testing.T) {
		// Input is invalid when either the total satellites, the satellites
		// launched per site, or the number of launch sites is 0
		invalidCases := []error{
			LaunchSatellites(0, 4, 2, 2),
			LaunchSatellites(500, 0, 2, 2),
			LaunchSatellites(500, 4, 3, 0),
		}
		for _, invalidCase := range invalidCases {
			if invalidCase == nil {
				t.Errorf("Error should not be nil on an invalid case")
			}
			if invalidCase != nil && !strings.Contains(invalidCase.Error(), "Invalid input") {
				t.Errorf("Error message should contain: Invalid input")
			}
		}
	})
}
