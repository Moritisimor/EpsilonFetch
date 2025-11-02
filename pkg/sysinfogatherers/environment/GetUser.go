package environment

import (
	"os/user"
)

func GetUser() string {
	currentUser, userErr := user.Current()
	if userErr != nil {
		return "???"
	}

	return currentUser.Name
}