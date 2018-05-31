package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var znachVar2 []byte

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

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

	x := "37.622504 55.753215"
	y := "30.315868 59.939095"

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Nav_points:              []string{x, y}}

	znachVar2, _ = json.Marshal(znachVar1)

	znachVar2 := zapro()
	_ = znachVar2
	fmt.Println("1")

}

func zapr2(c chan string) {

	x := "37.622504 55.753215"
	y := "28.331786 57.819365"

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Nav_points:              []string{x, y}}

	znachVar2, _ = json.Marshal(znachVar1)

	znachVar2 := zapro()
	_ = znachVar2
	fmt.Println("2")
}

func zapr3(c chan string) {

	x := "37.622504 55.753215"
	y := "40.516939 64.539393"

	znachVar1 := &znach{
		Fueltank_volume:         600,
		Consumption:             30,
		Fuel_in_tank:            100,
		Min_fuel_balance_litres: 50,
		Nav_points:              []string{x, y}}

	znachVar2, _ = json.Marshal(znachVar1)

	znachVar2 := zapro()
	_ = znachVar2

	fmt.Println("3")

}

func zapr4(c chan string) {

	x := "37.6172999 55.755826"
	y := "48.033574 46.347869"

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
func zapr5(c chan string) {

	x := "37.6172999 55.755826"
	y := "38.974711 45.035566"

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

func zapr6(c chan string) {

	x := "37.6172999 55.755826"
	y := "47.504682 42.98306"

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
func zapr7(c chan string) {

	x := "37.6172999 55.755826"
	y := "45.694909 43.317776"

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

func zapr8(c chan string) {

	x := "37.6172999 55.755826"
	y := "44.270181 46.308309"

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
func zapr9(c chan string) {

	x := "37.6172999 55.755826"
	y := "44.516939 48.707103"

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

func zapr10(c chan string) {

	x := "37.6172999 55.755826"
	y := "39.414526 57.185866"

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
func zapr11(c chan string) {

	x := "37.6172999 55.755826"
	y := "50.101783 53.195538"

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

func zapr12(c chan string) {

	x := "37.6172999 55.755826"
	y := "46.034158 51.533103"

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

func zapr13(c chan string) {

	x := "37.6172999 55.755826"
	y := "55.096955 51.768199"

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

func zapr14(c chan string) {

	x := "37.6172999 55.755826"
	y := "53.204843 56.852593"

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
func zapr15(c chan string) {

	x := "37.6172999 55.755826"
	y := "56.229398 58.010374"

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

func zapr16(c chan string) {

	x := "73.396221 61.254035"
	y := "76.680974 66.083963"

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
func zapr17(c chan string) {

	x := "37.6172999 55.755826"
	y := "69.018902 61.00318"

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

func zapr18(c chan string) {

	x := "37.6172999 55.755826"
	y := "76.569628 60.939716"

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
func zapr19(c chan string) {

	x := "37.6172999 55.755826"
	y := "92.852572 56.010563"

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

func zapr20(c chan string) {

	x := "37.6172999 55.755826"
	y := "61.40259 55.160026"

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
func zapr21(c chan string) {

	x := "37.6172999 55.755826"
	y := "54.738227 55.984619"

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

func zapr22(c chan string) {

	x := "37.6172999 55.755826"
	y := "65.534328 57.153033"

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

func zapr23(c chan string) {

	x := "37.6172999 55.755826"
	y := "73.368212 54.989342"

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

func zapr24(c chan string) {

	x := "37.6172999 55.755826"
	y := "82.92043 55.030199"

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
func zapr25(c chan string) {

	x := "37.6172999 55.755826"
	y := "86.087314 55.354968"

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

func zapr26(c chan string) {

	x := "37.6172999 55.755826"
	y := "83.769948 53.355084"

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
func zapr27(c chan string) {

	x := "37.6172999 55.755826"
	y := "104.28066 52.286387"

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

func zapr28(c chan string) {

	x := "37.6172999 55.755826"
	y := "107.584574 51.834464"

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
func zapr29(c chan string) {

	x := "37.6172999 55.755826"
	y := "113.499432 52.033973"

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

func zapr30(c chan string) {

	x := "37.6172999 55.755826"
	y := "135.071917 48.480223"

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
func zapr31(c chan string) {

	x := "37.6172999 55.755826"
	y := "131.885485 43.115536"

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

func zapr32(c chan string) {

	x := "37.6172999 55.755826"
	y := "33.074981 68.970682"

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

func zapr33(c chan string) {

	x := "37.6172999 55.755826"
	y := "73.396221 61.254035"

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

func zapr34(c chan string) {

	x := "37.6172999 55.755826"
	y := "76.680974 66.083963"

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
func zapr35(c chan string) {

	x := "47.504682 42.98306"
	y := "104.28066 52.286387"

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

func zapr36(c chan string) {

	x := "30.315868 59.939095"
	y := "37.622504 55.753215"

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
func zapr37(c chan string) {

	x := "30.315868 59.939095"
	y := "28.331786 57.819365"

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

func zapr38(c chan string) {

	x := "30.315868 59.939095"
	y := "40.516939 64.539393"

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
func zapr39(c chan string) {

	x := "30.315868 59.939095"
	y := "48.033574 46.347869"

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

func zapr40(c chan string) {

	x := "30.315868 59.939095"
	y := "38.974711 45.035566"

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
func zapr41(c chan string) {

	x := "30.315868 59.939095"
	y := "39.723062 43.585525"

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

func zapr42(c chan string) {

	x := "30.315868 59.939095"
	y := "47.504682 42.98306"

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

func zapr43(c chan string) {

	x := "30.315868 59.939095"
	y := "45.694909 43.317776"

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

func zapr44(c chan string) {

	x := "30.315868 59.939095"
	y := "44.270181 46.308309"

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
func zapr45(c chan string) {

	x := "30.315868 59.939095"
	y := "44.516939 48.707103"

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

func zapr46(c chan string) {

	x := "30.315868 59.939095"
	y := "39.414526 57.185866"

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
func zapr47(c chan string) {

	x := "30.315868 59.939095"
	y := "50.101783 53.195538"

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

func zapr48(c chan string) {

	x := "30.315868 59.939095"
	y := "46.034158 51.533103"

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
func zapr49(c chan string) {

	x := "30.315868 59.939095"
	y := "55.096955 51.768199"

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

func zapr50(c chan string) {

	x := "30.315868 59.939095"
	y := "53.204843 56.852593"

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
func zapr51(c chan string) {

	x := "30.315868 59.939095"
	y := "56.229398 58.010374"

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

func zapr52(c chan string) {

	x := "54.738227 55.984619"
	y := "92.852572 56.010563"

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

func zapr53(c chan string) {

	x := "30.315868 59.939095"
	y := "69.018902 61.00318"

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

func zapr54(c chan string) {

	x := "30.315868 59.939095"
	y := "76.569628 60.939716"

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
func zapr55(c chan string) {

	x := "30.315868 59.939095"
	y := "92.852572 56.010563"

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

func zapr56(c chan string) {

	x := "30.315868 59.939095"
	y := "61.40259 55.160026"

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
func zapr57(c chan string) {

	x := "30.315868 59.939095"
	y := "54.738227 55.984619"

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

func zapr58(c chan string) {

	x := "30.315868 59.939095"
	y := "65.534328 57.153033"

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
func zapr59(c chan string) {

	x := "30.315868 59.939095"
	y := "73.368212 54.989342"

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

func zapr60(c chan string) {

	x := "30.315868 59.939095"
	y := "82.92043 55.030199"

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
func zapr61(c chan string) {

	x := "30.315868 59.939095"
	y := "86.087314 55.354968"

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

func zapr62(c chan string) {

	x := "30.315868 59.939095"
	y := "83.769948 53.355084"

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

func zapr63(c chan string) {

	x := "30.315868 59.939095"
	y := "104.28066 52.286387"

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

func zapr64(c chan string) {

	x := "30.315868 59.939095"
	y := "107.584574 51.834464"

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
func zapr65(c chan string) {

	x := "30.315868 59.939095"
	y := "113.499432 52.033973"

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

func zapr66(c chan string) {

	x := "30.315868 59.939095"
	y := "135.071917 48.480223"

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
func zapr67(c chan string) {

	x := "30.315868 59.939095"
	y := "131.885485 43.115536"

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

func zapr68(c chan string) {

	x := "30.315868 59.939095"
	y := "73.396221 61.254035"

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
func zapr69(c chan string) {

	x := "30.315868 59.939095"
	y := "76.680974 66.083963"

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

func zapr70(c chan string) {

	x := "45.694909 43.317776"
	y := "28.331786 57.819365"

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
func zapr71(c chan string) {

	x := "76.680974 66.083963"
	y := "30.315868 59.939095"

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

func zapr72(c chan string) {

	x := "54.738227 55.984619"
	y := "28.331786 57.819365"

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

func zapr73(c chan string) {

	x := "28.331786 57.819365"
	y := "30.315868 59.939095"

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

func zapr74(c chan string) {

	x := "28.331786 57.819365"
	y := "37.622504 55.753215"

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
func zapr75(c chan string) {

	x := "28.331786 57.819365"
	y := "40.516939 64.539393"

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

func zapr76(c chan string) {

	x := "28.331786 57.819365"
	y := "48.033574 46.347869"

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
func zapr77(c chan string) {

	x := "28.331786 57.819365"
	y := "38.974711 45.035566"

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

func zapr78(c chan string) {

	x := "28.331786 57.819365"
	y := "39.723062 43.585525"

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
func zapr79(c chan string) {

	x := "28.331786 57.819365"
	y := "47.504682 42.98306"

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

func zapr80(c chan string) {

	x := "28.331786 57.819365"
	y := "45.694909 43.317776"

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
func zapr81(c chan string) {

	x := "28.331786 57.819365"
	y := "44.270181 46.308309"

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

func zapr82(c chan string) {

	x := "28.331786 57.819365"
	y := "44.516939 48.707103"

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

func zapr83(c chan string) {

	x := "28.331786 57.819365"
	y := "39.414526 57.185866"

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
func zapr84(c chan string) {

	x := "28.331786 57.819365"
	y := "50.101783 53.195538"

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

func zapr85(c chan string) {

	x := "28.331786 57.819365"
	y := "46.034158 51.533103"

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
func zapr86(c chan string) {

	x := "28.331786 57.819365"
	y := "55.096955 51.768199"

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

func zapr87(c chan string) {

	x := "28.331786 57.819365"
	y := "53.204843 56.852593"

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
func zapr88(c chan string) {

	x := "28.331786 57.819365"
	y := "56.229398 58.010374"

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

func zapr89(c chan string) {

	x := "28.331786 57.819365"
	y := "69.018902 61.00318"

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
func zapr90(c chan string) {

	x := "28.331786 57.819365"
	y := "76.569628 60.939716"

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

func zapr91(c chan string) {

	x := "28.331786 57.819365"
	y := "54.738227 55.984619"

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
func zapr92(c chan string) {

	x := "28.331786 57.819365"
	y := "65.534328 57.153033"

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

func zapr93(c chan string) {

	x := "28.331786 57.819365"
	y := "73.368212 54.989342"

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

func zapr94(c chan string) {

	x := "28.331786 57.819365"
	y := "82.92043 55.030199"

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

func zapr95(c chan string) {

	x := "28.331786 57.819365"
	y := "83.769948 53.355084"

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
func zapr96(c chan string) {

	x := "28.331786 57.819365"
	y := "104.28066 52.286387"

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

func zapr97(c chan string) {

	x := "28.331786 57.819365"
	y := "107.584574 51.834464"

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
func zapr98(c chan string) {

	x := "28.331786 57.819365"
	y := "113.499432 52.033973"

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
func zapr99(c chan string) {

	x := "28.331786 57.819365"
	y := " 35.071917 48.480223"

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

func zapr100(c chan string) {

	x := "28.331786 57.819365"
	y := "33.074981 68.970682"

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

func zapr101(c chan string) {

	x := "76.680974 66.083963"
	y := "73.396221 61.254035"

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

func zapr102(c chan string) {

	x := "76.680974 66.083963"
	y := "33.074981 68.970682"

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

func zapr103(c chan string) {

	x := "76.680974 66.083963"
	y := "131.885485 43.115536"

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

func zapr104(c chan string) {

	x := "76.680974 66.083963"
	y := "135.071917 48.480223"

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
func zapr105(c chan string) {

	x := "76.680974 66.083963"
	y := "113.499432 52.033973"

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

func zapr106(c chan string) {

	x := "76.680974 66.083963"
	y := "107.584574 51.834464"

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
func zapr107(c chan string) {

	x := "76.680974 66.083963"
	y := "104.28066 52.286387"

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

func zapr108(c chan string) {

	x := "76.680974 66.083963"
	y := "83.769948 53.355084"

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
func zapr109(c chan string) {

	x := "76.680974 66.083963"
	y := "86.087314 55.354968"

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

func zapr110(c chan string) {

	x := "76.680974 66.083963"
	y := "82.92043 55.030199"

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
func zapr111(c chan string) {

	x := "76.680974 66.083963"
	y := "73.368212 54.989342"

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

func zapr112(c chan string) {

	x := "76.680974 66.083963"
	y := "65.534328 57.153033"

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

func zapr113(c chan string) {

	x := "76.680974 66.083963"
	y := "54.738227 55.984619"

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

func zapr114(c chan string) {

	x := "76.680974 66.083963"
	y := "61.40259 55.160026"

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
func zapr115(c chan string) {

	x := "76.680974 66.083963"
	y := "92.852572 56.010563"

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

func zapr116(c chan string) {

	x := "76.680974 66.083963"
	y := "76.569628 60.939716"

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
func zapr117(c chan string) {

	x := "76.680974 66.083963"
	y := "69.018902 61.00318"

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

func zapr118(c chan string) {

	x := "76.680974 66.083963"
	y := "56.229398 58.010374"

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
func zapr119(c chan string) {

	x := "76.680974 66.083963"
	y := "53.204843 56.852593"

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

func zapr120(c chan string) {

	x := "76.680974 66.083963"
	y := "55.096955 51.768199"

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
func zapr121(c chan string) {

	x := "76.680974 66.083963"
	y := "46.034158 51.533103"

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

func zapr122(c chan string) {

	x := "76.680974 66.083963"
	y := "50.101783 53.195538"

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

func zapr123(c chan string) {

	x := "76.680974 66.083963"
	y := "39.414526 57.185866"

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

func zapr124(c chan string) {

	x := "76.680974 66.083963"
	y := "44.516939 48.707103"

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
func zapr125(c chan string) {

	x := "76.680974 66.083963"
	y := "44.270181 46.308309"

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

func zapr126(c chan string) {

	x := "76.680974 66.083963"
	y := "45.694909 43.317776"

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
func zapr127(c chan string) {

	x := "76.680974 66.083963"
	y := "47.504682 42.98306"

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

func zapr128(c chan string) {

	x := "76.680974 66.083963"
	y := "39.723062 43.585525"

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
func zapr129(c chan string) {

	x := "76.680974 66.083963"
	y := "38.974711 45.035566"

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

func zapr130(c chan string) {

	x := "76.680974 66.083963"
	y := "48.033574 46.347869"

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
func zapr131(c chan string) {

	x := "76.680974 66.083963"
	y := "40.516939 64.539393"

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

func zapr132(c chan string) {

	x := "76.680974 66.083963"
	y := "28.331786 57.819365"

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

func zapr133(c chan string) {

	x := "76.680974 66.083963"
	y := "30.315868 59.939095"

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

func zapr134(c chan string) {

	x := "76.680974 66.083963"
	y := "37.622504 55.753215"

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
func zapr135(c chan string) {

	x := "73.396221 61.254035"
	y := "33.074981 68.970682"

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

func zapr136(c chan string) {

	x := "73.396221 61.254035"
	y := "131.885485 43.115536"

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
func zapr137(c chan string) {

	x := "73.396221 61.254035"
	y := "135.071917 48.480223"

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

func zapr138(c chan string) {

	x := "73.396221 61.254035"
	y := "113.499432 52.033973"

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
func zapr139(c chan string) {

	x := "73.396221 61.254035"
	y := "107.584574 51.834464"

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

func zapr140(c chan string) {

	x := "73.396221 61.254035"
	y := "104.28066 52.286387"

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
func zapr141(c chan string) {

	x := "73.396221 61.254035"
	y := "83.769948 53.355084"

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

func zapr142(c chan string) {

	x := "73.396221 61.254035"
	y := "86.087314 55.354968"

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

func zapr143(c chan string) {

	x := "73.396221 61.254035"
	y := "82.92043 55.030199"

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

func zapr144(c chan string) {

	x := "73.396221 61.254035"
	y := "65.534328 57.153033"

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
func zapr145(c chan string) {

	x := "73.396221 61.254035"
	y := "54.738227 55.984619"

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

func zapr146(c chan string) {

	x := "73.396221 61.254035"
	y := "61.40259 55.160026"

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
func zapr147(c chan string) {

	x := "73.396221 61.254035"
	y := "92.852572 56.010563"

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

func zapr148(c chan string) {

	x := "73.396221 61.254035"
	y := "76.569628 60.939716"

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
func zapr149(c chan string) {

	x := "73.396221 61.254035"
	y := "69.018902 61.00318"

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

func zapr150(c chan string) {

	x := "73.396221 61.254035"
	y := "56.229398 58.010374"

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

var w string

func main() {
	var c chan string = make(chan string)
	go zapr1(c)
	go zapr2(c)
	go zapr3(c)
	go zapr4(c)
	go zapr5(c)
	go zapr6(c)
	go zapr7(c)
	go zapr8(c)
	go zapr9(c)
	go zapr10(c)
	go zapr11(c)
	go zapr12(c)
	go zapr13(c)
	go zapr14(c)
	go zapr15(c)
	go zapr16(c)
	go zapr17(c)
	go zapr18(c)
	go zapr19(c)
	go zapr20(c)
	go zapr21(c)
	go zapr22(c)
	go zapr23(c)
	go zapr24(c)
	go zapr25(c)
	go zapr26(c)
	go zapr27(c)
	go zapr28(c)
	go zapr29(c)
	go zapr30(c)
	go zapr31(c)
	go zapr32(c)
	go zapr33(c)
	go zapr34(c)
	go zapr35(c)
	go zapr36(c)
	go zapr37(c)
	go zapr38(c)
	go zapr39(c)
	go zapr40(c)
	go zapr41(c)
	go zapr42(c)
	go zapr43(c)
	go zapr44(c)
	go zapr45(c)
	go zapr46(c)
	go zapr47(c)
	go zapr48(c)
	go zapr49(c)
	go zapr50(c)
	go zapr51(c)
	go zapr52(c)
	go zapr53(c)
	go zapr54(c)
	go zapr55(c)
	go zapr56(c)
	go zapr57(c)
	go zapr58(c)
	go zapr59(c)
	go zapr60(c)
	go zapr61(c)
	go zapr62(c)
	go zapr63(c)
	go zapr64(c)
	go zapr65(c)
	go zapr66(c)
	go zapr67(c)
	go zapr68(c)
	go zapr69(c)
	go zapr70(c)
	go zapr71(c)
	go zapr72(c)
	go zapr73(c)
	go zapr74(c)
	go zapr75(c)
	go zapr76(c)
	go zapr77(c)
	go zapr78(c)
	go zapr79(c)
	go zapr80(c)
	go zapr81(c)
	go zapr82(c)
	go zapr83(c)
	go zapr84(c)
	go zapr85(c)
	go zapr86(c)
	go zapr87(c)
	go zapr88(c)
	go zapr89(c)
	go zapr90(c)
	go zapr91(c)
	go zapr92(c)
	go zapr93(c)
	go zapr94(c)
	go zapr95(c)
	go zapr96(c)
	go zapr97(c)
	go zapr98(c)
	go zapr99(c)
	go zapr100(c)
	/*
		go zapr101(c)
		go zapr102(c)
		go zapr103(c)
		go zapr104(c)
		go zapr105(c)
		go zapr106(c)
		go zapr107(c)
		go zapr108(c)
		go zapr109(c)
		go zapr110(c)
		go zapr111(c)
		go zapr112(c)
		go zapr113(c)
		go zapr114(c)
		go zapr115(c)
		go zapr116(c)
		go zapr117(c)
		go zapr118(c)
		go zapr119(c)
		go zapr120(c)
		go zapr121(c)
		go zapr122(c)
		go zapr123(c)
		go zapr124(c)
		go zapr125(c)
		go zapr126(c)
		go zapr127(c)
		go zapr128(c)
		go zapr129(c)
		go zapr130(c)
		go zapr131(c)
		go zapr132(c)
		go zapr133(c)
		go zapr134(c)
		go zapr135(c)
		go zapr136(c)
		go zapr137(c)
		go zapr138(c)
		go zapr139(c)
		go zapr140(c)
		go zapr141(c)
		go zapr142(c)
		go zapr143(c)
		go zapr144(c)
		go zapr145(c)
		go zapr146(c)
		go zapr147(c)
		go zapr148(c)
		go zapr149(c)
		go zapr150(c)
	*/
	var input string
	fmt.Scanln(&input)

}
