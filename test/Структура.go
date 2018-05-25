package znacheniya

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

/*func zikl() {

	if i == 0 {
		q := 10
		return q
	} else if i == 1 {
		q := 11
	}

}

func main() {

	i := 0

	for i <= 4 {
		i = i + 1

		//type q struct {
		//	x float64
		//	y float64
		//}

		Moscow := "37.622504, 55.753215"
		Peter := "30.315868, 59.939095"

		//Moscow := q{37.622504, 55.753215}
		//Peter := q{30.315868, 59.939095}
		//Pskov := q{28.331786, 57.819365}
		//Arkhangelsk := q{40.516939, 64.539393}
		//Astrakhan := q{48.033574, 46.347869}

		/*Krasnodar = 38.974711 45.035566
		Sochi = 39.723062 43.585525
		molensk = 39.723062 43.585525
		Makhachkala = 47.504682 42.98306
		Terrible = 45.694909 43.317776
		Elista = 44.270181 46.308309
		Volgograd = 44.516939 48.707103
		Rostov = 39.414526 57.185866
		Samara = 50.101783 53.195538
		Saratov = 46.034158 51.533103
		Orenburg = 55.096955 51.768199
		Izhevsk = 53.204843 56.852593
		Perm = 56.229398 58.010374
		KhantyMansiysk = 69.018902 61.00318
		Nizhnevartovsk = 76.569628 60.939716
		Krasnoyarsk = 92.852572 56.010563
		Chelyabinsk = 61.40259 55.160026
		Ufa = 54.738227, 55.984619
		Tyumen = 65.534328 57.153033
		Omsk = 73.368212 54.989342
		Novosibirsk = 82.92043 55.030199
		Kemerovo = 86.087314 55.354968
		Barnaul = 83.769948 53.355084
		Irkutsk = 104.28066 52.286387
		UlanUde = 107.584574 51.834464
		Chita = 113.499432 52.033973
		Khabarovsk = 135.071917 48.480223
		Vladivostok = 131.885485 43.115536
		Murmansk = 33.074981 68.970682
		Surgut = 73.396221 61.254035
		Urengoy = 76.680974 66.083963

		fmt.Println(q)

		type znach struct {
			fueltank_volume         int
			consumption             int
			fuel_in_tank            int
			min_fuel_balance_litres int
			min_fuel_balance_km     string
			min_distance            string
			min_benefit             string
			min_fuel_level_percent  string
			max_distance_from_route string
			fuel_types              string
			nav_points              []string
		}

		from := Moscow
		to := Peter

		znacheniya := znach{
			fueltank_volume:         600,
			consumption:             30,
			fuel_in_tank:            100,
			min_fuel_balance_litres: 50,
			min_fuel_balance_km:     "null",
			min_distance:            "null",
			min_benefit:             "null",
			min_fuel_level_percent:  "null",
			max_distance_from_route: "null",
			fuel_types:              "null",
			nav_points:              []string{from, to}}

		fmt.Println(znacheniya)
	}
}*/
