package python_check_updates

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func fillName(v string, repeatCount int) string {
	return v + strings.Repeat(" ", repeatCount)
}

func fillVersion(v string, repeatCount int) string {
	return strings.Repeat(" ", repeatCount) + v
}

func EchoPrintVersion(pkgList []PackageInfo)  {
	maxNameLen := 0
	maxVersionLen := 0
	marginLeft := 3

	for _, v := range pkgList {
		nameLen := len(v.Name)
		versionLen := len(v.NewVersion)
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
		if versionLen > maxVersionLen {
			maxVersionLen = versionLen
		}
	}

	for _, v := range pkgList {
		nameLen := len(v.Name)
		versionLen := len(v.NewVersion)

		pkgName := v.Name
		pkgVersion := v.NewVersion

		diffNameCount := maxNameLen + marginLeft - nameLen
		diffVersionCount := maxVersionLen - versionLen

		if diffNameCount != 0 {
			pkgName = fillName(pkgName, diffNameCount)
		}

		if diffVersionCount != 0 {
			pkgVersion = fillVersion(pkgVersion, diffVersionCount)
		}

		fmt.Print(" ", pkgName + " →   ")
		color.Yellow(pkgVersion)
	}
}
