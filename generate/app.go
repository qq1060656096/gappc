package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
	"path"
	"strings"
)

func CreateApp(projectPath string, stdout io.Writer, appName string) {
	dirPath := projectPath + "/" + ProjectDirs[AppDir]
	name := utils.ToCamelCase(path.Base(appName))
	packagePath := path.Dir(appName)
	packageName := utils.ToSnakeCase(path.Base(packagePath))
	fileDir := fmt.Sprintf("%s/%s/%s", dirPath, path.Dir(packagePath), packageName)
	os.MkdirAll(fileDir, os.ModePerm)
	fileName := utils.ToSnakeCase(name)
	file := fileDir + fmt.Sprintf("/%s.go", fileName)

	content := strings.Replace(appTemplate, "{{packageName}}", packageName, -1)
	content = strings.Replace(content, "{{apiName}}", name, -1)
	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}