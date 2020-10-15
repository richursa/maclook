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

// SetAPIMacAdressIO : set api for mac address lookup from macadrress.io
func SetAPIMacAdressIO(apiKey string) {
	defaultAPIProvider.url = "https://api.macaddress.io/v1"
	defaultAPIProvider.apiKey = apiKey
	defaultAPIProvider.output = "json"
}

// MacLookUp : Look up the macaddress from the given url and return result
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
