package rootfs

import (
	"os"
)

func dumpFile(filename string, date []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(date)
	return err
}