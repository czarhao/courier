package actions

import (
	"fmt"
	"runtime"
)

const version = "0.1"

func Version() string {
	v := fmt.Sprintf("Courier version: %v && ", version)
	v += fmt.Sprintf("GO version: %v \n", runtime.Version())
	return v
}
