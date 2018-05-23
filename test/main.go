//resp, err := http.PostForm("http://example.com/form", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func Derevna() string {
	//data := []byte(`{"37.622504":"55.753215"}, {"24.667522":"57.429639"}`)
	//r := bytes.NewReader(data)
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://geocode-maps.yandex.ru/1.x/?format=json&geocode=%D0%9D%D0%B0%D0%B1%D0%B5")
	//resp, err := client.Head("https://fonts.googleapis.com/icon?family=Material+Icons")
	//resp, err := client.Post("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json", r)
	//resp, err := client.PostForm("https://fonts.googleapis.com/icon?family=Material+Icons", url.Values{"49.119873046875": {"123"}, "55.819801652442436": {"123"}})
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
	return string(bytes)
}

func Moscow() string {
	//data := []byte(`{"37.622504":"55.753215"}, {"24.667522":"57.429639"}`)
	//r := bytes.NewReader(data)
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://geocode-maps.yandex.ru/1.x/?format=json&geocode=%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0")
	//resp, err := client.Head("https://fonts.googleapis.com/icon?family=Material+Icons")
	//resp, err := client.Post("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json", r)
	//resp, err := client.PostForm("https://fonts.googleapis.com/icon?family=Material+Icons", url.Values{"49.119873046875": {"123"}, "55.819801652442436": {"123"}})
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
	return string(bytes)
} // min_fuel_balance_km: 0, min_distance: 0, min_benefit: 0, min_fuel_level_percent:0 , max_distance_from_route: 0, fuel_types: 0

func main() {

	data := []byte(`{"fueltank_volume": 600,"consumption": 30,"fuel_in_tank": 100,"min_fuel_balance_litres": 50,"min_fuel_balance_km": 0,"min_distance": 0,"min_benefit": 0,"min_fuel_level_percent":0 ,"max_distance_from_route": 0,"fuel_types": 0, "nav_points": [0:"24.667522 57.429639", 1:37.622504 55.753215"]"}`)

	r := bytes.NewReader(data)
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	//resp, err := client.FormPost("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json", r)
	resp, err := client.Post("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json, text/plain, */*", r)
	//v := r
	//fmt.Println(v)
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
