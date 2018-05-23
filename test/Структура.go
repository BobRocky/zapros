package main

import (
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
	}

	type nav_points struct {
		x  float64
		y  float64
		x1 float64
		y2 float64
	}

	cor := nav_points{37.622504, 55.753215, 49.106324, 55.798551}

	data := znach{600, 30, 100, 50, 0, 0, 0, 0, 0, 0}

	fmt.Println(data)
	fmt.Println(cor)
	fmt.Println("q", data, cor)
}
