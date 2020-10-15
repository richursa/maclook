package main

import (
	"flag"
	"fmt"
	"os"

	"./maclook"
	"github.com/tidwall/gjson"
)

func main() {
	macAddress := flag.String("macaddress", "", "macaddress")                    // get macaddress from command line parameter
	detail := flag.Bool("detail", false, "set to true to print detailed report") // check whether user wants to print a detailed report
	apiKey := flag.String("apikey", "", "apikey obtained from macaddress.io")    // get apikey from commandline
	flag.Parse()
	if *apiKey == "" { // if apikey is not given as command line parameter, try to get it from env
		*apiKey = os.Getenv("apikey")
		if *apiKey == "" { // if apikey is not present in env and commandline parameter, raise an error
			fmt.Println("Please provide apikey with -apikey=APIKEY_HERE")
			fmt.Println("Run with -h for help")
			os.Exit(1)
		}
	}
	if *macAddress == "" {
		fmt.Println("Please provide macaddress with -macaddress=MACADDRESS_HERE")
		fmt.Println("Run with -h for help")
		os.Exit(1)
	}
	maclook.SetAPIMacAdressIO(*apiKey) // setup macaddress api and apikey as macaddress.io
	jsonText := maclook.MacLookUp(*macAddress)
	if *detail {
		fmt.Println(string(jsonText))
	} else {
		m, ok := gjson.Parse(string(jsonText)).Value().(map[string]interface{}) // 3rd party library to parse deeply nested json
		if !ok {
			fmt.Println("ERROR: couldnt parse json")
			os.Exit(1)
		}
		fmt.Println("MAC Address  :", *macAddress)
		fmt.Println("Company Name :", m["vendorDetails"].(map[string]interface{})["companyName"])
	}
}
