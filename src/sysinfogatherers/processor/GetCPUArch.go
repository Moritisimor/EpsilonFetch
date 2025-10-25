package processor

import "runtime"

func GetArch() string {
	return runtime.GOARCH
}
