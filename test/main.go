//resp, err := http.PostForm("http://example.com/form", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct")
	//resp, err := client.Head("http://fuel.truckmove.ru")
	//resp, err := client.Post("http://fuel.truckmove.ru")
	resp, err := client.PostForm("http://google.com", url.Values{"login_login": {"123"}})
	if err != nil {
		log.Panic("Responce:", resp, "\nError:", err)
	}
	defer resp.Body.Close()
	bytes := make([]byte, 100)
	for {
		bytes = bytes[:cap(bytes)]
		n, err := resp.Body.Read(bytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		bytes = bytes[:n]
		fmt.Println(string(bytes))
	}
}
