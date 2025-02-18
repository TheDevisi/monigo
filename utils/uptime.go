package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func Uptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		return "Error while fetching uptime" // Returning an error message if an error occurred
	}

	seconds := uptime % 60
	minutes := (uptime / 60) % 60
	hours := uptime / 3600

	return fmt.Sprintf("Uptime: %d hour(s) %d minute(s) %d second(s)", hours, minutes, seconds) // Formatting and returning the uptime
}
