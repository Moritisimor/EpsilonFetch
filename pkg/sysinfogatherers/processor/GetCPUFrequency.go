package processor

import (
	"fmt"
	"log"
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUFrequency() string {
	info, err := cpu.Info()
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprint(info[0].Mhz / 1000) + " Ghz" 
}