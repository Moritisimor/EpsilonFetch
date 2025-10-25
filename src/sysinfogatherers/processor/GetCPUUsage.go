package processor

import (
	"fmt"
	"log"
	"time"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUUsage() string {
	usage, err := cpu.Percent(time.Second / 2, false)
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%.2f", usage[0]) + "% used"
}