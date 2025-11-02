package processor

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUUsage() string {
	usage, err := cpu.Percent(time.Second / 2, false)
	if err != nil {
		return "??% used"
	}

	return fmt.Sprintf("%.2f", usage[0]) + "% used"
}