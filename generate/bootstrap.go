package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
)

func CreateBootstrap(projectPath string, stdout io.Writer) {
	dirPath := projectPath + "/" + ProjectDirs[BootstrapDir]
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + "/application.go"
	utils.WriteToFile(file, BootstrapTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}