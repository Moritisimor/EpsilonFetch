package processor

import "github.com/shirou/gopsutil/v4/cpu"

func GetCPUModel() string {
	info, err := cpu.Info()
	if err != nil {
		return "???"
	}

	return info[0].ModelName
}