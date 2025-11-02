package environment

import (
	"os"
)

func GetHost() string {
	currentHost, hostErr := os.Hostname()
	if hostErr != nil {
		return "???"
	}
	
	return currentHost
}