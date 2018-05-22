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

/*package main

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
	/*	data := []byte(`{"foo":"bar"}`)
		r := bytes.NewReader(data)
		resp, err := http.Post("google.com", "application/json", r)
		if err != nil {
			return
		}
		fmt.Println(resp)
}

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
package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	templates := populateTemplates()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")

			if template != nil {
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
