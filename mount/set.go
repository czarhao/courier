package mount

import (
	"fmt"
	"os"
)

func SetMount() error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get pwd fail, error: %v", err)
	}
	if err := chroot(pwd); err != nil {
		return err
	}
	return mountOther()
}

func ClearMount() error {
	return umountOther()
}