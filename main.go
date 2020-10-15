package main

import (
	"flag"
	"fmt"
	"os"

	"./maclook"
	"github.com/tidwall/gjson"
)

var macAddress string

func main() {
	macAddress := flag.String("macaddress", "", "macaddress")
	detail := flag.Bool("detail", false, "set to true to print detailed report")
	flag.Parse()
	if *macAddress == "" {
		fmt.Println("Please provide macaddress with -macaddress=MACADDRESS_HERE")
		fmt.Println("Run with -h for help")
		os.Exit(1)
	}
	maclook.SetAPIMacAdressIO()
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
