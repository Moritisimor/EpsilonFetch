package processor

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUThreads() string {
	physicalCores, err := cpu.Counts(true)
	if err != nil {
		return "?? Threads"
	}

	return fmt.Sprintf("%d Threads", physicalCores)
}