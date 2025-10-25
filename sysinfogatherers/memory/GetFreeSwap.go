package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
	"log"
	"fmt"
)

func GetFreeSwap() string {
	swapInfo, err :=mem.SwapMemory()
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%.2f GB", float64(swapInfo.Free) / 1000000000)
}
