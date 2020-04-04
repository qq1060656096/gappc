package generate

import (
	"fmt"
	"io"
	"os"
)

func CreateLog(projectPath string, stdout io.Writer) {
	dirPath := projectPath + "/" + ProjectDirs[LogDir]
	os.MkdirAll(dirPath, os.ModePerm)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", dirPath, "\x1b[0m")
}