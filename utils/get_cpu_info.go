package utils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	log "github.com/sirupsen/logrus"
)

func GetCPU() string {
	log.Debug("Retrieving CPU information")

	// Get CPU info
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Error("Failed to get CPU info:", err)
		return "Error fetching CPU information"
	}

	// Get CPU usage percentage
	percentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Error("Failed to get CPU usage:", err)
		return "Error fetching CPU usage"
	}

	if len(cpuInfo) == 0 {
		log.Warn("No CPU info retrieved")
		return "No CPU information available"
	}

	// Format CPU information
	info := fmt.Sprintf("💻 CPU Information:\n"+
		"Model: %s\n"+
		"Cores: %d\n"+
		"MHz: %.2f\n"+
		"Current Usage: %.2f%%",
		cpuInfo[0].ModelName,
		cpuInfo[0].Cores,
		cpuInfo[0].Mhz,
		percentage[0])

	log.WithFields(log.Fields{
		"model": cpuInfo[0].ModelName,
		"cores": cpuInfo[0].Cores,
		"usage": percentage[0],
	}).Debug("CPU information retrieved")

	return info
}
