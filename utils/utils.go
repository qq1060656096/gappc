package utils

import (
	"os"
	"path/filepath"
)

func MustCheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteToFile(fileName, content string) {
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