package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type analysis struct {
	XMLName           xml.Name          `xml:"analysis"`
	Description       string            `xml:",innerxml"`
	Calltree          calltree          `xml:"calltree"`
	Processes         processes         `xml:"processes"`
	Running_processes running_processes `xml:"running_processes"`
}

type calltree struct {
	Process_call string `xml:"process_call"`
	Description  string `xml:",innerxml"`
}

type processes struct {
	Process     string `xml:"process"`
	Description string `xml:",innerxml"`
}

type running_processes struct {
	Running_process string `xml:"running_process"`
	Description     string `xml:",innerxml"`
}

func main() {
	xmlfile, err := os.Open("../tmp/sandbox/12070853.xml")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer xmlfile.Close()

	data, err := ioutil.ReadAll(xmlfile)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	v := analysis{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: ", err)
		return
	}

	//fmt.Println(v)

	fmt.Println(v.Calltree.Description)
}
