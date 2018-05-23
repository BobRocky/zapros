package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type znach struct {
		fueltank_volume         int
		consumption             int
		fuel_in_tank            int
		min_fuel_balance_litres int
		min_fuel_balance_km     int
		min_distance            int
		min_benefit             int
		min_fuel_level_percent  int
		max_distance_from_route int
		fuel_types              int
		nav_points              []float64
	}

	znacheniya := znach{
		fueltank_volume:         600,
		consumption:             30,
		fuel_in_tank:            100,
		min_fuel_balance_litres: 50,
		min_fuel_balance_km:     0,
		min_distance:            0,
		min_benefit:             0,
		min_fuel_level_percent:  0,
		max_distance_from_route: 0,
		fuel_types:              0,
		nav_points:              []float64{37.622504, 55.753215, 49.106324, 55.798551}}

	fmt.Println(znacheniya)

	data, _ := json.Marshal(znacheniya)

	fmt.Println(data)
}
