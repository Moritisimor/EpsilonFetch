package processor

import (
	"log"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUModel() string {
	info, err := cpu.Info()
	if err != nil {
		log.Fatal(err.Error())
	}

	return info[0].ModelName
}