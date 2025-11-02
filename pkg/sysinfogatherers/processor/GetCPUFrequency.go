package processor

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUFrequency() string {
	info, err := cpu.Info()
	if err != nil {
		return "??? Ghz"
	}

	return fmt.Sprint(info[0].Mhz / 1000) + " Ghz" 
}