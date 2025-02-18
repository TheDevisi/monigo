package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemory() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Info(err)
	}
	fmt.Println(memory)
	return fmt.Sprintf("All: %v MB, Avalible: %v MB, Using: %.2f%%",
		memory.Total/1024/1024, memory.Available/1024/1024, memory.UsedPercent)
}
