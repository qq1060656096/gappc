package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
	"path"
	"strings"
)

func CreateApi(projectPath string, stdout io.Writer, apiName string) {
	dirPath := projectPath + "/" + ProjectDirs[ApiDir]
	name := utils.ToCamelCase(path.Base(apiName))
	packagePath := path.Dir(apiName)
	packageName := utils.ToSnakeCase(path.Base(packagePath))
	fileDir := fmt.Sprintf("%s/%s/%s", dirPath, path.Dir(packagePath), packageName)
	os.MkdirAll(fileDir, os.ModePerm)
	fileName := utils.ToSnakeCase(name)
	file := fileDir + fmt.Sprintf("/%s.go", fileName)

	content := strings.Replace(apiTemplate, "{{packageName}}", packageName, -1)
	content = strings.Replace(content, "{{apiName}}", name, -1)
	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}