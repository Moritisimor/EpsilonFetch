package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
	"fmt"
)

func GetFreeMem() string {
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return "?? GB"
	}

	return fmt.Sprintf("%.2f GB", float64(memoryStats.Available) / 1000000000)
}
