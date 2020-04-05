package generate

import (
"fmt"
"gappc/utils"
"io"
"os"
	"path"
	"strings"
)

func CreateRouters(projectPath string, stdout io.Writer)  {
	projectName := path.Base(projectPath)
	dirPath := projectPath + "/" + ProjectDirs[RoutersDir]
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + "/api_router.go"
	content := strings.Replace(apiRoutersTemplate, "{{projectName}}", projectName, -1)
	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/app_router.go"
	content = strings.Replace(appRoutersTemplate, "{{projectName}}", projectName, -1)
	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/router.go"
	utils.WriteToFile(file, routersTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}