package processor

import (
	"fmt"
	"log"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUCores() string {
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%d Cores", physicalCores)
}
