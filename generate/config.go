package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
)

func CreateConfig(projectPath string, stdout io.Writer)  {
	dirPath := projectPath + "/" + ProjectDirs[ConfigDir]
	os.MkdirAll(dirPath, os.ModePerm)
	file := dirPath + "/.app.env"
	utils.WriteToFile(file, appConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/.cache.env"
	utils.WriteToFile(file, cacheConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
	file = dirPath + "/.db.env"
	utils.WriteToFile(file, dbConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}