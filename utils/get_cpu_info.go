package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPU() string {
	cpu, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("CPU: %v", cpu)

}
