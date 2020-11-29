package python_check_updates

import (
	"bufio"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type PackageInfo struct {
	Name string
	NewVersion string
	OldVersion string
}

// ReadRequirements("requirements.txt")
func ReadRequirements(fileName string) []PackageInfo {
	cwd, err := os.Getwd()
	splitRegex := regexp.MustCompile("==|>=|<=|~=|>|<")
	trimRegex := regexp.MustCompile(`\s+`)

	if err != nil {
		panic(err)
	}

	absPath := filepath.Join(cwd, fileName)
	file, err := os.Open(absPath)

	color.Green("Checking " + absPath + "\n")

	if err != nil {
		log.Panicln("File open failed ")
	}

	defer file.Close()

	inputReader := bufio.NewReader(file)
	var pkgList []PackageInfo

	for {
		inputStr, readerErr := inputReader.ReadString('\n')
		inputStr = trimRegex.ReplaceAllString(inputStr, "")
		pkgNameSplit := splitRegex.Split(inputStr, -1)

		if len(pkgNameSplit) == 0 {
			continue
		}

		pkgName := pkgNameSplit[0]
		if pkgName != "" {
			pkgList = append(pkgList, PackageInfo{
				Name: pkgName,
				OldVersion: strings.Replace(inputStr, pkgName, "", -1),
			})
		}

		if readerErr == io.EOF {
			break
		}
	}

	return pkgList
}

