package environment

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
)

func GetUptime() string {
	upTime, err := host.Uptime()
	if err != nil {
		return "?? Days, ?? Hours, ?? Minutes"
	}

	days := upTime / 86400
	hours := upTime / 3600
	minutes := (upTime % 3600) / 60

	return fmt.Sprintf("%d Days, %d Hours, %d Minutes", days, hours, minutes)
}
