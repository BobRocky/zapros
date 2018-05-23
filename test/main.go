//resp, err := http.PostForm("http://example.com/form", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	type znach struct {
		Fueltank_volume         int    `json:"fueltank_volume"`
		Consumption             int    `json:"consumption"`
		Fuel_in_tank            int    `json:"fuel_in_tank"`
		Min_fuel_balance_litres int    `json:"min_fuel_balance_litres"`
		Min_fuel_balance_km     string `json:"min_fuel_balance_km"`
		Min_distance            string `json:"min_distance"`
		Min_benefit             string `json:"min_benefit"`
		Min_fuel_level_percent  string `json:"min_fuel_level_percent"`
		Max_distance_from_route string `json:"max_distance_from_route"`
		Fuel_types              string `json:"fuel_types"`

		Nav_points []string `json:"nav_points"`
	}

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Min_fuel_balance_km:     "null",
		Min_distance:            "null",
		Min_benefit:             "null",
		Min_fuel_level_percent:  "null",
		Max_distance_from_route: "null",
		Fuel_types:              "null",
		Nav_points:              []string{"37.622504 55.753215", "49.106324 55.798551"}}

	//fmt.Println(znachVar1)

	znachVar2, _ := json.Marshal(znachVar1)

	fmt.Println(string(znachVar2))

	r := bytes.NewReader(znachVar2)
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	//resp, err := client.FormPost("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json", r)
	resp, err := client.Post("https://tms-api-service-dev.redradar.ru/fuel_calc", "application/json, text/plain, */*", r)
	//v := r
	//fmt.Println(r)
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
