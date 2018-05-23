package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
/*
package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	data := []byte(`{"foo":"bar"}`)
	r := bytes.NewReader(data)
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct")
	//resp, err := client.Head("http://fuel.truckmove.ru")
	resp, err := client.Post("vk.com", "application/json", r)
	//resp, err := client.PostForm("http://google.com", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
	if err != nil {
		log.Panic("Responce:", resp, "\nError:", err)
	}
	defer resp.Body.Close()
	bytes := make([]byte, 100)
	for {
		bytes = bytes[:cap(bytes)]
		n, err := resp.Body.Read(bytes)
		//if err != nil {
		//if err == io.EOF {
		//	break
		//}
		//	log.Panic(err)
		//}
		bytes = bytes[:n]
	}
	log.Println(string(bytes))
	data := []byte(`{"foo":"bar"}`)
		r := bytes.NewReader(data)
		resp, err := http.Post("google.com", "application/json", r)
		if err != nil {
			return
		}
		fmt.Println(resp)
}*/
/*
package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	data := []byte(`{"foo":"bar"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post("https://www.google.com", "application/octet-stream", r)
	fmt.Printf("%v %v", err, resp)
}
*/

func main() {
	var username string = "login_login"
	var passwd string = "login_password"
	v := url.Values{}
	v.Add(username, "login")
	v.Add(passwd, "passwd")
	client := &http.Client{}
	resp, err := http.PostForm("http://example.com/form", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s = string(bodyText)
	fmt.Println(s)
}
