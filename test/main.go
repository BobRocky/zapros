//resp, err := http.PostForm("http://example.com/form", url.Values{"login_login": {"Value"}, "login_password": {"123"}})
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	i := 0
	for i <= 0 {

		i = i + 1

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

		switch i {
		case 1:
			x := "37.6172999 55.755826"
			y := "30.315868, 59.939095"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

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

			/*client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*36)
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
			fmt.Println(i)*/
		case 2:
			x := "37.6172999 55.755826"
			y := "28.331786, 57.819365"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 3:
			x := "37.6172999 55.755826"
			y := "40.516939, 64.539393"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 4:
			x := "37.6172999 55.755826"
			y := "48.033574, 46.347869"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 5:
			x := "37.6172999 55.755826"
			y := "38.974711 45.035566"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 6:
			x := "37.6172999 55.755826"
			y := "39.723062 43.585525"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 7:
			x := "37.6172999 55.755826"
			y := "47.504682 42.98306"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 8:
			x := "37.6172999 55.755826"
			y := "45.694909 43.317776"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 9:
			x := "37.6172999 55.755826"
			y := "44.270181 46.308309"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 10:
			x := "37.6172999 55.755826"
			y := "44.516939 48.707103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 11:
			x := "37.6172999 55.755826"
			y := "39.414526 57.185866"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 12:
			x := "37.6172999 55.755826"
			y := "50.101783 53.195538"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 13:
			x := "37.6172999 55.755826"
			y := "46.034158 51.533103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 14:
			x := "37.6172999 55.755826"
			y := "55.096955 51.768199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 15:
			x := "37.6172999 55.755826"
			y := "53.204843 56.852593"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 16:
			x := "37.6172999 55.755826"
			y := "56.229398 58.010374"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 17:
			x := "37.6172999 55.755826"
			y := "69.018902 61.00318"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 18:
			x := "37.6172999 55.755826"
			y := "76.569628 60.939716"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 19:
			x := "37.6172999 55.755826"
			y := "92.852572 56.010563"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 20:
			x := "37.6172999 55.755826"
			y := "61.40259 55.160026"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 21:
			x := "37.6172999 55.755826"
			y := "54.738227, 55.984619"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 22:
			x := "37.6172999 55.755826"
			y := "65.534328 57.153033"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 23:
			x := "37.6172999 55.755826"
			y := "73.368212 54.989342"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 24:
			x := "37.6172999 55.755826"
			y := "82.92043 55.030199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 25:
			x := "37.6172999 55.755826"
			y := "86.087314 55.354968"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 26:
			x := "37.6172999 55.755826"
			y := "83.769948 53.355084"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 27:
			x := "37.6172999 55.755826"
			y := "104.28066 52.286387"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 28:
			x := "37.6172999 55.755826"
			y := "107.584574 51.834464"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 29:
			x := "37.6172999 55.755826"
			y := "113.499432 52.033973"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 30:
			x := "37.6172999 55.755826"
			y := "135.071917 48.480223"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 31:
			x := "37.6172999 55.755826"
			y := "131.885485 43.115536"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 32:
			x := "37.6172999 55.755826"
			y := "33.074981 68.970682"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 33:
			x := "37.6172999 55.755826"
			y := "73.396221 61.254035"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 34:
			x := "37.6172999 55.755826"
			y := "76.680974 66.083963"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 35:
			x := "30.315868, 59.939095"
			y := "37.622504, 55.753215"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 36:
			x := "30.315868, 59.939095"
			y := "28.331786, 57.819365"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 37:
			x := "30.315868, 59.939095"
			y := "40.516939, 64.539393"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 38:
			x := "30.315868, 59.939095"
			y := "48.033574, 46.347869"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 39:
			x := "30.315868, 59.939095"
			y := "38.974711 45.035566"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 40:
			x := "30.315868, 59.939095"
			y := "39.723062 43.585525"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 41:
			x := "30.315868, 59.939095"
			y := "39.723062 43.585525"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 42:
			x := "30.315868, 59.939095"
			y := "47.504682 42.98306"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 43:
			x := "30.315868, 59.939095"
			y := "45.694909 43.317776"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 44:
			x := "30.315868, 59.939095"
			y := "44.270181 46.308309"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 45:
			x := "30.315868, 59.939095"
			y := "44.516939 48.707103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 46:
			x := "30.315868, 59.939095"
			y := "39.414526 57.185866"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 47:
			x := "30.315868, 59.939095"
			y := "50.101783 53.195538"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 48:
			x := "30.315868, 59.939095"
			y := "46.034158 51.533103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 49:
			x := "30.315868, 59.939095"
			y := "55.096955 51.768199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 50:
			x := "30.315868, 59.939095"
			y := "53.204843 56.852593"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 51:
			x := "30.315868, 59.939095"
			y := "56.229398 58.010374"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 52:
			x := "30.315868, 59.939095"
			y := "69.018902 61.00318"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 53:
			x := "30.315868, 59.939095"
			y := "76.569628 60.939716"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 54:
			x := "30.315868, 59.939095"
			y := "92.852572 56.010563"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 55:
			x := "30.315868, 59.939095"
			y := "61.40259 55.160026"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 56:
			x := "30.315868, 59.939095"
			y := "54.738227, 55.984619"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 57:
			x := "30.315868, 59.939095"
			y := "65.534328 57.153033"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 58:
			x := "30.315868, 59.939095"
			y := "73.368212 54.989342"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 59:
			x := "30.315868, 59.939095"
			y := "82.92043 55.030199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 60:
			x := "30.315868, 59.939095"
			y := "86.087314 55.354968"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 61:
			x := "30.315868, 59.939095"
			y := "83.769948 53.355084"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 62:
			x := "30.315868, 59.939095"
			y := "104.28066 52.286387"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 63:
			x := "30.315868, 59.939095"
			y := "107.584574 51.834464"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 64:
			x := "30.315868, 59.939095"
			y := "113.499432 52.033973"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 65:
			x := "30.315868, 59.939095"
			y := "135.071917 48.480223"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 66:
			x := "30.315868, 59.939095"
			y := "131.885485 43.115536"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 67:
			x := "30.315868, 59.939095"
			y := "33.074981 68.970682"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 68:
			x := "30.315868, 59.939095"
			y := "73.396221 61.254035"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 69:
			x := "30.315868, 59.939095"
			y := "76.680974 66.083963"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 70:
			x := "28.331786, 57.819365"
			y := "37.622504, 55.753215"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 71:
			x := "28.331786, 57.819365"
			y := "30.315868, 59.939095"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 72:
			x := "28.331786, 57.819365"
			y := "40.516939, 64.539393"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 73:
			x := "28.331786, 57.819365"
			y := "48.033574, 46.347869"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 74:
			x := "28.331786, 57.819365"
			y := "38.974711 45.035566"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 75:
			x := "28.331786, 57.819365"
			y := "39.723062 43.585525"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 76:
			x := "28.331786, 57.819365"
			y := "47.504682 42.98306"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 77:
			x := "28.331786, 57.819365"
			y := "45.694909 43.317776"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 78:
			x := "28.331786, 57.819365"
			y := "44.270181 46.308309"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 79:
			x := "28.331786, 57.819365"
			y := "44.516939 48.707103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 80:
			x := "28.331786, 57.819365"
			y := "39.414526 57.185866"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 81:
			x := "28.331786, 57.819365"
			y := "50.101783 53.195538"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 82:
			x := "28.331786, 57.819365"
			y := "46.034158 51.533103"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 83:
			x := "28.331786, 57.819365"
			y := "55.096955 51.768199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 84:
			x := "28.331786, 57.819365"
			y := "53.204843 56.852593"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 85:
			x := "28.331786, 57.819365"
			y := "56.229398 58.010374"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 86:
			x := "28.331786, 57.819365"
			y := "69.018902 61.00318"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 87:
			x := "28.331786, 57.819365"
			y := "76.569628 60.939716"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 88:
			x := "28.331786, 57.819365"
			y := "92.852572 56.010563"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 89:
			x := "28.331786, 57.819365"
			y := "61.40259 55.160026"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 90:
			x := "28.331786, 57.819365"
			y := "54.738227, 55.984619"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 91:
			x := "28.331786, 57.819365"
			y := "65.534328 57.153033"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 92:
			x := "28.331786, 57.819365"
			y := "73.368212 54.989342"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 93:
			x := "28.331786, 57.819365"
			y := "82.92043 55.030199"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 94:
			x := "28.331786, 57.819365"
			y := "86.087314 55.354968"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 95:
			x := "28.331786, 57.819365"
			y := "83.769948 53.355084"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 96:
			x := "28.331786, 57.819365"
			y := "104.28066 52.286387"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 97:
			x := "28.331786, 57.819365"
			y := "107.584574 51.834464"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 98:
			x := "28.331786, 57.819365"
			y := "113.499432 52.033973"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 99:
			x := "28.331786, 57.819365"
			y := "135.071917 48.480223"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
		case 100:
			x := "28.331786, 57.819365"
			y := "131.885485 43.115536"

			znachVar1 := &znach{
				Fueltank_volume:         600,
				Consumption:             30,
				Fuel_in_tank:            100,
				Min_fuel_balance_litres: 50,
				Nav_points:              []string{x, y}}

			fmt.Println(znachVar1)

			znachVar2, _ := json.Marshal(znachVar1)

			fmt.Println(string(znachVar2))

			req, err := http.NewRequest("POST", "https://tms-api-service-dev.redradar.ru/fuel_calc", bytes.NewBuffer(znachVar2))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bytes := make([]byte, 1024*10)
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
	}
}
