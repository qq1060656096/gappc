package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
	"strings"
)

func CreateMain(projectPath string, stdout io.Writer, projectName string) {
	dirPath := projectPath
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + "/main.go"
	content := strings.Replace(mainTemplate, "{{projectName}}", projectName, -1)
	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}