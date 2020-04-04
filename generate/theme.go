package generate

import (
	"fmt"
	"gappc/utils"
	"io"
	"os"
	"strings"
)

func CreateTheme(projectPath string, stdout io.Writer, themeName, subDirName, fileName string) {
	dirPath := projectPath + "/" + ProjectDirs[ThemesDir] + "/" + themeName
	os.MkdirAll(dirPath, os.ModePerm)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", dirPath, "\x1b[0m")
	subDirPath := dirPath + "/" + subDirName
	os.MkdirAll(subDirPath, os.ModePerm)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", subDirPath, "\x1b[0m")
	file := subDirPath + "/" + fileName
	content := strings.Replace(themeTemplate, "{{themeName}}", themeName, -1)
	content = strings.Replace(content, "{{subDirName}}", subDirName, -1)
	content = strings.Replace(content, "{{fileName}}", fileName, -1)

	utils.WriteToFile(file, content)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}