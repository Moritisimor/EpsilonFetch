package environment

import (
	"fmt"
	"os"
)

func GetHost() string {
	currentHost, hostErr := os.Hostname()
	if hostErr != nil {
		fmt.Printf("Could not read OS Hostname!\nError: %e\n", hostErr)
		os.Exit(1)
	}
	
	return currentHost
}