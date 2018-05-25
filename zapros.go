package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var znachVar2 []byte

type znach struct {
	Fueltank_volume         int `json:"fueltank_volume"`
	Consumption             int `json:"consumption"`
	Fuel_in_tank            int `json:"fuel_in_tank"`
	Min_fuel_balance_litres int `json:"min_fuel_balance_litres"`
	Min_fuel_balance_km     int `json:"min_fuel_balance_km"`
	Min_distance            int `json:"min_distance"`
	Min_benefit             int `json:"min_benefit"`
	Min_fuel_level_percent  int `json:"min_fuel_level_percent"`
	Max_distance_from_route int `json:"max_distance_from_route"`
	Fuel_types              int `json:"fuel_types"`

	Nav_points []string `json:"nav_points"`
}

func zapro() string {
	req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}

func zapr1(c chan string) {

	x := "37.6172999 55.755826"
	y := "30.315868, 59.939095"

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Nav_points:              []string{x, y}}

	znachVar2, _ = json.Marshal(znachVar1)

	znachVar2 := zapro()
	_ = znachVar2

}

func zapr2(c chan string) {

	x := "37.6172999 55.755826"
	y := "28.331786, 57.819365"

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Nav_points:              []string{x, y}}

	znachVar2, _ = json.Marshal(znachVar1)

	znachVar2 := zapro()
	_ = znachVar2
}

func main() {

	var c chan string = make(chan string)

	go zapr1(c)
	go zapr2(c)

	var input string
	fmt.Scanln(&input)

}
