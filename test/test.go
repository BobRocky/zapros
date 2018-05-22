/*package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Привет")
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
http://ident.me*/

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct")
	//resp, err := client.Head("http://fuel.truckmove.ru")
	//resp, err := client.Post("http://fuel.truckmove.ru")
	//, err := client.PostForm("http://fuel.truckmove.ru")
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
	}
	log.Println(string(bytes))
}
