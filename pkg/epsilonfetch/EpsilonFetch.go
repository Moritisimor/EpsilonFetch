package epsilonfetch

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/EpsilonFetch/pkg/sysinfogatherers/environment"
	"github.com/Moritisimor/EpsilonFetch/pkg/sysinfogatherers/memory"
	"github.com/Moritisimor/EpsilonFetch/pkg/sysinfogatherers/processor"
)

// EpsilonFetch is the heart function, printing prettily formatted and colored text about the PC's Hard- and Software.
func EpsilonFetch() {
	if environment.GetOS() == "android" {
		color.PrintRedln("EpsilonFetch does not support Android.")
		return
	}

	color.PrintRainbowln("[Epsilon Fetch]")

	color.PrintRedln("[Runtime]")
	color.PrintRed("-> Operating System: ")
	color.PrintMagentaln(environment.GetOS())
	color.PrintRed("-> OS Version: ")
	color.PrintMagentaln(environment.GetDistro())
	color.PrintRed("-> Kernel Version: ")
	color.PrintMagentaln(environment.GetKernel())
	color.PrintRed("-> Hostname: ")
	color.PrintMagentaln(environment.GetHost())
	color.PrintRed("-> Current User: ")
	color.PrintMagentaln(environment.GetUser().Name)
	color.PrintRed("-> Uptime: ")
	color.PrintMagentaln(environment.GetUptime())

	color.PrintGreenln("[CPU]")
	color.PrintGreen("-> Model: ")
	color.PrintMagentaln(processor.GetCPUModel())
	color.PrintGreen("-> Cores: ")
	color.PrintMagentaln(processor.GetCPUCores())
	color.PrintGreen("-> Threads: ")
	color.PrintMagentaln(processor.GetCPUThreads())
	color.PrintGreen("-> Frequency: ")
	color.PrintMagentaln(processor.GetCPUFrequency())
	color.PrintGreen("-> Architecture: ")
	color.PrintMagentaln(processor.GetArch())
	color.PrintGreen("-> Current Usage: ")
	color.PrintMagentaln(processor.GetCPUUsage())

	color.PrintBlueln("[Memory]")
	color.PrintBlue("-> Total Physical Installed: ")
	color.PrintMagentaln(memory.GetTotalMem())
	color.PrintBlue("-> Total Physical Available: ")
	color.PrintMagentaln(memory.GetFreeMem())
	color.PrintBlue("-> Total Swap: ")
	color.PrintMagentaln(memory.GetSwapSize())
	color.PrintBlue("-> Total Swap Available: ")
	color.PrintMagentaln(memory.GetFreeSwap())
	
	color.PrintRainbowln("[Epsilon Fetch]\n")
}

