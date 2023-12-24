package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

type MapRange struct {
	delta int64
	destination_start int64
	destination_end int64
	source_start int64
	source_end int64
	length int64
}

type InequalityTuple struct {
	first int64
	second int64
}

type LinearEqParams struct {
	inequalities []InequalityTuple
	additive_elements []int64
}

type LinearEqParamsCollection struct {
	seed_to_soil_le_eq_params LinearEqParams
	soil_to_fertilizer_le_eq_params LinearEqParams
	fertilizer_to_water_le_eq_params LinearEqParams
	water_to_light_le_eq_params LinearEqParams
	light_to_temperature_le_eq_params LinearEqParams
	temperature_to_humidity_le_eq_params LinearEqParams
	humidity_to_location_le_eq_params LinearEqParams
}

func main() {
	data := readData("input2.txt")

	var seed_ids []int64

	var seed_to_soil_le_eq_params LinearEqParams
	var soil_to_fertilizer_le_eq_params LinearEqParams
	var fertilizer_to_water_le_eq_params LinearEqParams
	var water_to_light_le_eq_params LinearEqParams
	var light_to_temperature_le_eq_params LinearEqParams
	var temperature_to_humidity_le_eq_params LinearEqParams
	var humidity_to_location_le_eq_params LinearEqParams

	for i, chunk := range data {
		if i == 0 {
			seed_ids = getSeedIds(chunk[0])
		} else {
			map_range_components := tointegerListofLists(chunk[1:])
			map_range_list := generateMapRangeList(map_range_components)

			switch i {
			case 1:
				seed_to_soil_le_eq_params = getLinearEqParams(map_range_list)
			case 2:
				soil_to_fertilizer_le_eq_params = getLinearEqParams(map_range_list)
			case 3:
				fertilizer_to_water_le_eq_params = getLinearEqParams(map_range_list)
			case 4:
				water_to_light_le_eq_params = getLinearEqParams(map_range_list)
			case 5:
				light_to_temperature_le_eq_params = getLinearEqParams(map_range_list)
			case 6:
				temperature_to_humidity_le_eq_params = getLinearEqParams(map_range_list)
			case 7:
				humidity_to_location_le_eq_params = getLinearEqParams(map_range_list)
			default:
				fmt.Println(i, "This should never happen")
			}
		}
	}


	equations := LinearEqParamsCollection {
		seed_to_soil_le_eq_params: seed_to_soil_le_eq_params,
		soil_to_fertilizer_le_eq_params: soil_to_fertilizer_le_eq_params,
		fertilizer_to_water_le_eq_params: fertilizer_to_water_le_eq_params,
		water_to_light_le_eq_params: water_to_light_le_eq_params,
		light_to_temperature_le_eq_params: light_to_temperature_le_eq_params,
		temperature_to_humidity_le_eq_params: temperature_to_humidity_le_eq_params,
		humidity_to_location_le_eq_params: humidity_to_location_le_eq_params,
	}
	
	min := seedToLocation(equations, seed_ids[0])
	for _, seed_id := range seed_ids {
		tmp := seedToLocation(equations, seed_id)
		if tmp < min {
			min = tmp
		}
	}

	fmt.Println(min)
}

func readData(filename string) [][]string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		
		if s != "" {
			chunk = append(chunk, s)
		} else {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			continue
		}

	}

	return chunks
}

func toint64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func getSeedIds(s string) []int64 {
	colon_index := strings.Index(s, ":")
	start := colon_index + 2
	seed_id_string := s[start:]
	string_seed_ids := strings.Split(seed_id_string, " ")
	numeric_seed_ids := make([]int64, 0)

	for _, seed_id := range string_seed_ids {
		numeric_seed_ids = append(numeric_seed_ids, toint64(seed_id))
	}

	return numeric_seed_ids
}

func tointegerListofLists(params []string) [][]int64{

	int64eger_list_of_lists := make([][]int64, 0)

	for _, param := range params {
		tmp := strings.Split(param, " ")
		tmp_numeric := make([]int64, 0)
		for _, str := range tmp {
			tmp_numeric = append(tmp_numeric, toint64(str))
		}
		int64eger_list_of_lists = append(int64eger_list_of_lists, tmp_numeric)
	}

	return int64eger_list_of_lists
}

func toMapRange(list []int64) MapRange {

	dstart := list[0]
	lnt := list[2]
	sstart := list[1]
	dlt := dstart - sstart

	map_range := MapRange {
		delta: dlt,
		destination_start: dstart,
		destination_end: dstart + lnt - 1,
		source_start: sstart,
		source_end: sstart + lnt - 1,
		length: lnt,
	}

	return map_range
}

func generateMapRangeList(params [][]int64) []MapRange {
	
	map_range_list := make([]MapRange, 0)

	for _, param := range params {
		map_range_list = append(map_range_list, toMapRange(param))
	}

	return map_range_list

}

func getLinearEqParams(map_ranges []MapRange) LinearEqParams{

	var linear_eq_params LinearEqParams

	for _, map_range := range map_ranges {
		linear_eq_params.inequalities = append(linear_eq_params.inequalities, InequalityTuple{first: map_range.source_start, second: map_range.source_end})
		linear_eq_params.additive_elements = append(linear_eq_params.additive_elements, map_range.delta)
	}

	return linear_eq_params
}

func getDecodedValue(linear_eq_params LinearEqParams, input_value int64) int64 {
	inequality_values := make([]int64, 0)

	for _, inequality_nmbs := range linear_eq_params.inequalities {
		var val int64 = 0

		if (inequality_nmbs.first <= input_value && input_value <= inequality_nmbs.second) {
			val = 1
		}
		inequality_values = append(inequality_values, val)
	}

	var final_addition int64 = 0

	for i, additive_element := range linear_eq_params.additive_elements {
		if inequality_values[i] * additive_element != 0 {
			final_addition += inequality_values[i] * additive_element
		}
	}

	return input_value + final_addition
	
}

func seedToLocation(equations LinearEqParamsCollection, seed int64) int64 {
	soil := getDecodedValue(equations.seed_to_soil_le_eq_params, seed)
	fertilizer := getDecodedValue(equations.soil_to_fertilizer_le_eq_params, soil)
	water := getDecodedValue(equations.fertilizer_to_water_le_eq_params, fertilizer)
	light := getDecodedValue(equations.water_to_light_le_eq_params, water)
	temperature := getDecodedValue(equations.light_to_temperature_le_eq_params, light)
	humidity := getDecodedValue(equations.temperature_to_humidity_le_eq_params, temperature)
	location := getDecodedValue(equations.humidity_to_location_le_eq_params, humidity)

	//fmt.Printf("%d -> %d -> %d -> %d -> %d -> %d -> %d -> %d\n", seed, soil, fertilizer, water, light, temperature, humidity, location)

	return location
}
