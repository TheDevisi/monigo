package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemory() string {
	log.Debug("Retrieving memory statistics")

	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Error("Failed to get memory info: ", err)
		return "Error fetching memory information"
	}

	log.WithFields(log.Fields{
		"total":     memory.Total,
		"available": memory.Available,
		"used":      memory.Used,
		"percent":   memory.UsedPercent,
	}).Debug("Memory statistics retrieved")

	return fmt.Sprintf("All: %v MB, Available: %v MB, Using: %.2f%%",
		memory.Total/1024/1024, memory.Available/1024/1024, memory.UsedPercent)
}
