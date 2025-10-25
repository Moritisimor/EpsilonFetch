package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
	"log"
	"fmt"
)

func GetTotalMem() string {
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%.2f GB", float64(memoryStats.Total) / 1000000000)
}
