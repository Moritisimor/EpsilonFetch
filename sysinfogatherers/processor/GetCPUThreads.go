package processor

import (
	"fmt"
	"log"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUThreads() string {
	physicalCores, err := cpu.Counts(true)
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%d Threads", physicalCores)
}