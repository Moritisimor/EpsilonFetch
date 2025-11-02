package processor

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUCores() string {
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		return "? Cores"
	}

	return fmt.Sprintf("%d Cores", physicalCores)
}
