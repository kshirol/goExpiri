package goExpiri

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type expiriconfig struct {
	ServName string `yaml:"servname"`
	ServFQDN string `yaml:"servfqdn"`
	ServPort int64  `yaml:"servport"`
}

// ConfigReader This function reads the YAML file for the following fields
// Server Name : Common Name for the HTTP Server
// Server FQDN : Fully Qualified Domain Name for the HTTP Server
// Server Port : Port on which this server serves:w
func ConfigReader() {
	expiriconf := expiriconfig{}

	ExpiriYAMLFile, err := ioutil.ReadFile("ExpiriConf.yml")
	if err != nil {
		fmt.Println("Failed to open ExpiriConf.yml File", err)
		return
	}

	err = yaml.Unmarshal([]byte(ExpiriYAMLFile), &expiriconf)
	if err != nil {
		fmt.Println("Failed to unmarshal the yaml File", err)
		return
	}

	fmt.Println(expiriconf.ServName, expiriconf.ServFQDN, expiriconf.ServPort)
}
