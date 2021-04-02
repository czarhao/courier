package image

import (
	"courier/configs"
	"io/ioutil"
	"os"
	"path"
)

func fileExist(dir, file string) (bool, string) {
	p := path.Join(dir, file)
	info, err := os.Stat(p)
	if err != nil {
		return false, p
	}
	if info.IsDir() {
		if len(DirFileName(p)) == 0 {
			_ = os.Remove(p)
			return false, p
		}
	}
	return true, p
}

func dumpFile(filename string, date []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(date)
	return err
}

func DirFileName(dir string) []string {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	fileNames := make([]string, 0, len(fileInfos))
	for _, file := range fileInfos {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

func ExpectContainerDir(config *configs.ContainerConfig) string {
	return path.Join(config.Image.ContainerDir, config.Other.Name)
}