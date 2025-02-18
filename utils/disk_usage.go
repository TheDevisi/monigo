package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

func GetDiskUsage() string {
	diskInfo, err := disk.Usage("/") // ("/" - корневой диск в Linux/macOS, "C:\" для Windows)
	if err != nil {
		return "Error while fetching disk usage"
	}

	// Форматируем данные
	return fmt.Sprintf("Disk: entire disk %.2f GB, avalible %.2f GB, using %.2f%%",
		float64(diskInfo.Total)/1024/1024/1024,
		float64(diskInfo.Free)/1024/1024/1024,
		diskInfo.UsedPercent)
}
