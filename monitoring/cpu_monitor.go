package monitoring

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	log "github.com/sirupsen/logrus"
)

// Work in progress here....

func MonitorCPU() bool {
	percentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Warn("Cannot get CPU percentage.")
		return false
	}

	if len(percentage) > 0 && percentage[0] >= 1.0 {
		log.Warn("CPU usage is high: ", percentage[0], "%")
		return true
	}
	return false
}
