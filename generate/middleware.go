package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
	"strings"
)

func CreateMiddleware(projectPath string, stdout io.Writer, middlewareName string) {
	dirPath := projectPath + "/" + ProjectDirs[MiddlewareDir]
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + fmt.Sprintf("/%s.go", utils.ToSnakeCase(middlewareName))
	utils.WriteToFile(file, strings.Replace(MiddlewareTemplate, "{{middlewareName}}", utils.ToCamelCase(middlewareName), -1))
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}