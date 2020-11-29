package python_check_updates

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb/v3"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func reqGetVersion(
	pkgInfoChan chan PackageInfo,
	pkgName string,
	oldVersion string,
) {
	defer wg.Done()

	version := "0.0.0"
	url := "https://pypi.org/project/" + pkgName
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	nameVer := doc.Find(".package-header__name").Text()
	nameVer = strings.TrimSpace(nameVer)
	subMatch := strings.Split(nameVer, " ")

	if len(subMatch) >= 2 {
		version = subMatch[1]
	}

	pkgInfoChan <- PackageInfo{
		Name: pkgName,
		NewVersion: version,
		OldVersion: oldVersion,
	}
}

func GetNewVersion(pkgList []PackageInfo) []PackageInfo {
	newPkgList := make([]PackageInfo, 0)
	pkgInfoChan := make(chan PackageInfo)
	bar := pb.StartNew(len(pkgList))
	bar.SetMaxWidth(50)

	for _, v := range pkgList {
		wg.Add(1)
		go func(pkgName string, oldVersion string) {
			reqGetVersion(pkgInfoChan, pkgName, oldVersion)
			bar.Increment()
		}(v.Name, v.OldVersion)
	}

	go func() {
		wg.Wait()
		close(pkgInfoChan)
	}()

	for {
		v, ok := <-pkgInfoChan
		if !ok {
			break
		}
		newPkgList = append(newPkgList, v)
	}

	bar.Finish()
	return newPkgList
}
