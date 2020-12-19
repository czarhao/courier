package container

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadCommandFromPip() ([]string, error) {
	pipe := os.NewFile(uintptr(3), "pipe")
	defer pipe.Close()
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(msg), " "), nil
}
