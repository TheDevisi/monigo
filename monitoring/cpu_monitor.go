package monitoring

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	log "github.com/sirupsen/logrus"
)

// Work in progress here....

func MonitorCPU() {
	percentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Warn("Cannot get CPU percentage.")
		return
	}
	if len(percentage) > 0 && percentage[0] >= 80.0 {
		log.Warn("CPU usage is high: ", percentage[0], "%")
		fmt.Println(percentage)

	}
}
