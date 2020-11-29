package python_check_updates

import (
	"encoding/json"
	"github.com/xjh22222228/python-check-updates/constants"
	"io/ioutil"
	"net/http"
)

type lastVersion struct {
	// name = version
	Version string `json:"name"`
}

func CheckLastVersion() string {
	url := "https://api.github.com/repos/xjh22222228/python-check-updates/releases/latest"

	resp, err := http.Get(url)

	if err != nil {
		return constants.Version
	}

	defer resp.Body.Close()

	var releaseInfo lastVersion
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return constants.Version
	}

	json.Unmarshal(body, &releaseInfo)

	return releaseInfo.Version
}

