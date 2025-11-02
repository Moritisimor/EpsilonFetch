package environment

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/shirou/gopsutil/v4/host"
)

func GetDistro() string {
	if GetOS() == "linux" || strings.HasSuffix(strings.ToLower(GetOS()), "bsd") {
		var file *os.File
		var err error

		file, err = os.Open("/etc/os-release")
		if err != nil {
			file, err = os.Open("/usr/lib/os-release")
			if err != nil {
				return "???"
			}
		}
		
		buf := ""
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "NAME") {
				buf += strings.TrimSuffix(strings.TrimPrefix(scanner.Text(), "NAME=\""), "\"") + " "
			}

			if strings.HasPrefix(scanner.Text(), "VERSION") && !strings.HasPrefix(scanner.Text(), "VERSION_ID") {
				buf += strings.TrimSuffix(strings.TrimPrefix(scanner.Text(), "VERSION=\""), "\"")
			}
		}

		return buf
	}

	_, os, version, err := host.PlatformInformation() 
	if err != nil {
		return "???"
	}
	
	return fmt.Sprintf("%s %s", os, version)
}
