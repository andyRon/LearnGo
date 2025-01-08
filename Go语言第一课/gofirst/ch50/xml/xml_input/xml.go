package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type recurlyservers struct {
	XMLName     xml.Name `xml_input:"servers"`
	Version     string   `xml_input:"version,attr"`
	Svs         []server `xml_input:"server"`
	Description string   `xml_input:",innerxml"`
}
type server struct {
	XMLName    xml.Name `xml_input:"server"`
	ServerName string   `xml_input:"serverName"`
	ServerIP   string   `xml_input:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml_input")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
