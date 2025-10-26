package environment

import (
	"log"

	"github.com/shirou/gopsutil/v4/host"
)

func GetKernel() string {
	kernel, err := host.KernelVersion()
	if err != nil {
		log.Fatal(err.Error())
	}

	if GetOS() == "linux" {
		return "linux " + kernel
	}
	return kernel
}
