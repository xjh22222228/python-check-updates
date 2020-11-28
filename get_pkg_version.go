package python_check_updates

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb/v3"
	"net/http"
	"regexp"
	"sync"
)

var wg sync.WaitGroup

func reqGetVersion(pkgInfoChan chan PackageInfo, pkgName string) {
	defer wg.Done()

	version := "0.0.0"
	url := "https://pypi.org/project/" + pkgName
	resp, err := http.Get(url)
	versionRegex := regexp.MustCompile(`\s(\d+\.\d+\.*\d*)`)

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

	name := doc.Find(".package-header__name").Text()
	subMatch := versionRegex.FindStringSubmatch(name)

	if len(subMatch) > 0 {
		version = subMatch[0]
	}

	pkgInfoChan <- PackageInfo{
		Name: pkgName,
		NewVersion: version,
	}
}

func GetNewVersion(pkgList []PackageInfo) []PackageInfo {
	newPkgList := make([]PackageInfo, 0)
	pkgInfoChan := make(chan PackageInfo)
	bar := pb.StartNew(len(pkgList))
	bar.SetMaxWidth(50)

	for _, v := range pkgList {
		wg.Add(1)
		go func(pkgName string) {
			reqGetVersion(pkgInfoChan, pkgName)
			bar.Increment()
		}(v.Name)
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
