package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
	"log"
	"fmt"
)

func GetFreeMem() string {
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%.2f GB", float64(memoryStats.Available) / 1000000000)
}
