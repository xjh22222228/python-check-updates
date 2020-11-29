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
	maxNewVersionLen := 0
	maxOldVersionLen := 0
	marginLeft := 3

	for _, v := range pkgList {
		nameLen := len(v.Name)
		newVersionLen := len(v.NewVersion)
		oldVersionLen := len(v.OldVersion)

		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
		if newVersionLen > maxNewVersionLen {
			maxNewVersionLen = newVersionLen
		}
		if oldVersionLen > maxOldVersionLen {
			maxOldVersionLen = oldVersionLen
		}
	}

	for _, v := range pkgList {
		pkgName := v.Name
		pkgNewVersion := v.NewVersion
		pkgOldVersion := v.OldVersion

		nameLen := len(pkgName)
		newVersionLen := len(pkgNewVersion)
		oldVersionLen := len(pkgOldVersion)

		diffNameCount := maxNameLen + marginLeft - nameLen
		diffNewVersionCount := maxNewVersionLen - newVersionLen
		diffOldVersionCount := maxOldVersionLen - oldVersionLen

		if diffNameCount != 0 {
			pkgName = fillName(pkgName, diffNameCount)
		}

		if diffNewVersionCount != 0 {
			pkgNewVersion = fillVersion(pkgNewVersion, diffNewVersionCount)
		}

		if diffOldVersionCount != 0 {
			pkgOldVersion = fillVersion(pkgOldVersion, diffOldVersionCount)
		}

		fmt.Printf(
			" %v    %v    →  %v\n",
			pkgName,
			pkgOldVersion,
			color.YellowString(pkgNewVersion),
		)
	}
}
