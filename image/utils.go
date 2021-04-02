package image

import (
	"io/ioutil"
	"os"
	"path"
)

func fileExist(dir, file string) (bool, string) {
	p := path.Join(dir, file)
	_, err := os.Stat(p)
	return err == nil, p
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