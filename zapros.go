package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://geocode-maps.yandex.ru/1.x/?format=json&geocode=Москва")
	if err != nil {
		log.Panic("Responce:", resp, "\nError:", err)
	}
	defer resp.Body.Close()
	bytes := make([]byte, 1024)
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
