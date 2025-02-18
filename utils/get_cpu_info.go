package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

func Get_cpu_info() {
	cpu, err := cpu.Info()
	fmt.Println(cpu, err)
}
