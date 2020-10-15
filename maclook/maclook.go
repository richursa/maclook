package maclook

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type apiProvider struct {
	url    string
	apiKey string
	output string
}

var defaultAPIProvider apiProvider

func SetAPIMacAdressIO() {
	defaultAPIProvider.url = "https://api.macaddress.io/v1"
	defaultAPIProvider.apiKey = "at_bWCRgYLtmZxNwdiqNFMU9NbIZuZrs"
	defaultAPIProvider.output = "json"
}

func MacLookUp(macAddress string) []byte {
	resp, err := http.Get(defaultAPIProvider.url + "?" + "apiKey=" + defaultAPIProvider.apiKey + "&" + "output=" + defaultAPIProvider.output + "&" + "search=" + macAddress)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	return body
}
