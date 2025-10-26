package environment

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/host"
)

func GetDistro() string {
	_, os, version, err := host.PlatformInformation() 
	if err != nil {
		log.Fatal(err.Error())
	}
	
	return fmt.Sprintf("%s %s", os, version)
}
