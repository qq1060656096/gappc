package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func MustCheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteToFile(fileName, content string) {
	ok, err := CheckPathExists(fileName)
	if ok {
		str := fmt.Sprintf("fail %s file exist\n%v", fileName, err)
		panic(errors.New(str))
	}

	f, err := os.Create(fileName)
	defer f.Close()
	MustCheckErr(err)
	_, err = f.WriteString(content)
	MustCheckErr(err)
}

func GetProjectInfo() (projectName, projectPath string, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	projectName = filepath.Base(pwd)
	projectPath = filepath.Dir(pwd)
	return
}

func GetCmdPath() (cmdPath string, err error)  {
	cmdPath, err = os.Getwd()
	return
}
