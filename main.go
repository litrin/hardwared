/*
	Copyright (C) 2016 Litrin J.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/litrin/hardwared/v1/host/format/hw_info"
)

var sURL = flag.String("base_url", "/v1/Host", "The base URL")
var sPort = flag.String("port", "8080", "Listening port")
var sDisableCache = flag.Bool("no-cache", true, "Enable cache info, for hotplug system please disable it!")

// configurations Only for server mode
var sServer = flag.Bool("server", false, "server mode")
var sFile = flag.String("config", "hosts.conf", "host name in each row enabled in server mode")

// define process cache for client
var HW_INFO = hw_info.All()

func hw_detail(w http.ResponseWriter, r *http.Request) {

	if *sDisableCache {
		HW_INFO = hw_info.All()
	}
	b, err := json.Marshal(HW_INFO)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Server", "Hardware Infomation Service")

	if err == nil {
		fmt.Fprintf(w, "%s", b)
	} else {
		fmt.Fprintf(w, "%s", err)
	}

}
func read_host_list(filename string) []string {
	file_handle, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer file_handle.Close()
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 0x20)
	content := ""
	for {
		n, err := file_handle.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		content = content + string(buff[:n])
	}

	return strings.Split(content, "\n")
}

func get_api_data(hostname string) hw_info.Page_framework {

	client := http.DefaultClient
	var oContent hw_info.Page_framework
	println("Fetching from: " + hostname)
	content, err := client.Get(hostname)
	if err != nil {
		panic(err)
		return oContent
	}

	buff := make([]byte, content.ContentLength)
	content.Body.Read(buff)

	json.Unmarshal(buff, &oContent)

	return oContent

}

func get_page() *hw_info.Page_framework {
	host_list := read_host_list(*sFile)
	channel := make(chan hw_info.Page_framework, len(host_list))
	//defer close(channel)

	length := 0
	for _, url := range host_list {
		if len(url) < 7 || url[:4] != "http" {

			continue

		}
		length++

		go func(c chan hw_info.Page_framework, api string) {
			body := get_api_data(api)

			c <- body
			//close(c)
		}(channel, url)
	}

	return hw_info.Combind(channel, length)

}

func cluster_detail(w http.ResponseWriter, r *http.Request) {

	cluster_info := get_page()

	b, err := json.Marshal(cluster_info)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Server", "Hardware Infomation Service")

	if err == nil {
		fmt.Fprintf(w, "%s", b)
	} else {
		fmt.Fprintf(w, "%s", err)
	}

}

func main() {
	flag.Parse()

	if *sServer {
		http.HandleFunc(*sURL, cluster_detail)
	} else {
		http.HandleFunc(*sURL, hw_detail)
	}
	listening_port := ":" + *sPort

	log.Fatal(http.ListenAndServe(listening_port, nil))
}
