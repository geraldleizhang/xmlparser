package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type analysis struct {
	XMLName           xml.Name          `xml:"analysis"`
	Cwsversion        string            `xml:"cwsversion,attr"`
	Time              string            `xml:"time,attr"`
	File              string            `xml:"file,attr"`
	Md5               string            `xml:"md5,attr"`
	Sha1              string            `xml:"sha1,attr"`
	Logpath           string            `xml:"logpath,attr"`
	Calltree          calltree          `xml:"calltree"`
	Processes         processes         `xml:"processes"`
	Running_processes running_processes `xml:"running_processes"`
}

type calltree struct {
	Process_call_list []process_call `xml:"process_call"`
}

type process_call struct {
	Index         string `xml:"index,attr"`
	Pid           string `xml:"pid,attr"`
	Filename      string `xml:"filename,attr"`
	Filename_hash string `xml:"filename_hash,attr"`
	Starttime     string `xml:"starttime,attr"`
	Startreason   string `xml:"startreason,attr"`
	calltree      string `xml:"calltree"`
}

type processes struct {
	Process_list []process `xml:"process"`
}

type process struct {
	Index             string `xml:"index,attr"`
	Pid               string `xml:"pid,attr"`
	Filename          string `xml:"filename,attr"`
	Filename_hash     string `xml:"filename_hash,attr"`
	Filesize          string `xml:"filesize,attr"`
	Md5               string `xml:"md5,attr"`
	Sha1              string `xml:"sha1,attr"`
	Username          string `xml:"username,attr"`
	Parentindex       string `xml:"parentindex,attr"`
	Starttime         string `xml:"starttime,attr"`
	Terminationtime   string `xml:"terminationtime,attr"`
	Startreason       string `xml:"startreason,attr"`
	Terminationreason string `xml:"terminationreason,attr"`
	Executionstatus   string `xml:"executionstatus,attr"`
	Applicationtype   string `xml:"applicationtype,attr"`
}

type running_processes struct {
	Running_processes_list []running_process `xml:"running_process"`
}

type running_process struct {
	Pid                string `xml:"pid,attr"`
	Filename           string `xml:"filename,attr"`
	Filename_hash      string `xml:"filename_hash,attr"`
	Cmdline_parameters string `xml:"cmdline_parameters,attr"`
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

	for _, index := range v.Processes.Process_list {
		fmt.Println(index.Pid)
	}
}
