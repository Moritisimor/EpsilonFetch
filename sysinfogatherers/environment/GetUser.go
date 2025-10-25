package environment

import (
	"fmt"
	"os"
	"os/user"
)

func GetUser() *user.User {
	currentUser, userErr := user.Current()
	if userErr != nil {
		fmt.Printf("Could not read current user!\nError: %e\n", userErr)
		os.Exit(1)
	}

	return currentUser
}