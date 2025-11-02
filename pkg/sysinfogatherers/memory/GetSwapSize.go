package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
	"fmt"
)

func GetSwapSize() string {
	swapInfo, err :=mem.SwapMemory()
	if err != nil {
		return "0 GB"
	}

	return fmt.Sprintf("%.2f GB", float64(swapInfo.Total) / 1000000000)
}
