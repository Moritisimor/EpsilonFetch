package environment

import "runtime"

func GetOS() string {
	return runtime.GOOS
}