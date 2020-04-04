package generate

import (
"fmt"
"gappc/utils"
"io"
"os"
)

func CreateRouters(projectPath string, stdout io.Writer)  {
	dirPath := projectPath + "/" + ProjectDirs[RoutersDir]
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + "/api_router.go"
	utils.WriteToFile(file, apiRoutersTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/app_router.go"
	utils.WriteToFile(file, appRoutersTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/router.go"
	utils.WriteToFile(file, routersTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}