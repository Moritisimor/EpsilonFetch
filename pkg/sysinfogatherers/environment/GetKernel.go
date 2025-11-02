package environment

import "github.com/shirou/gopsutil/v4/host"

func GetKernel() string {
	kernel, err := host.KernelVersion()
	if err != nil {
		return "???"
	}

	if GetOS() == "linux" {
		return "linux " + kernel
	}
	return kernel
}
