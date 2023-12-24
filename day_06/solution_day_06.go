package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type DataSet struct {
	time []int
	distance []int
}

type SimulationParameters struct {
	hold_time int
	total_time int
	speed int
	opportunity int
	distance int
	record_distance int
}

func main() {

	data := readData("input2.txt")
	fmt.Println(RunSimulation(data))
}

func readData(input_file string) DataSet {
	f, err := os.Open(input_file)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var data DataSet
	line := 0
	for scanner.Scan() {
		s := scanner.Text()

		tokens := strings.Split(s, " ")
		numeric_tokens := make([]int, 0)

		for _, token := range tokens[1:len(tokens)] {
			if token == "" {
				continue
			}
			numeric_tokens = append(numeric_tokens, toInt(token))
		}

		if line == 0 {
			data.time = numeric_tokens
		} else {
			data.distance = numeric_tokens
		}

		line++
	}

	return data
}

func toInt(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return val
}

func generateSimulationParameters(time int, record_distance int, hold_time int) SimulationParameters {

	simulation_parameters := SimulationParameters {
		hold_time: hold_time,
		total_time: time,
		speed: hold_time,
		opportunity: time - hold_time,
		distance: hold_time * (time - hold_time),
		record_distance: record_distance,
	}

	return simulation_parameters
}

func isWinningStrategy(simulation_parameters SimulationParameters) int {
	if simulation_parameters.distance > simulation_parameters.record_distance {
		return 1
	}

	return 0
}

func RunSimulation(data DataSet) int {

	winning_strategies := make([]int, 0)

	for i := 0; i < len(data.time); i++ {
		ways_to_win := 0
		for j := 0; j <= data.time[i]; j++ {
			tmp := generateSimulationParameters(data.time[i], data.distance[i], j)
			//displaySimulationParams(tmp)
			ways_to_win += isWinningStrategy(tmp)
		}

		winning_strategies = append(winning_strategies, ways_to_win)

		/*fmt.Println("--------------------------------")
		fmt.Println("--------------------------------")
		fmt.Println("--------------------------------")
		fmt.Printf("%d WAYS TO WIN IN A %d SECOND RACE!\n", ways_to_win, data.time[i])
		fmt.Println("--------------------------------")
		fmt.Println("--------------------------------")
		fmt.Println("--------------------------------")*/
	}

	total_winning_strategies := 1
	for _, val := range winning_strategies {
		total_winning_strategies *= val
	}

	return total_winning_strategies
}

func displaySimulationParams(simulation_parameters SimulationParameters) {
	fmt.Printf("hold_time: %d,\n", simulation_parameters.hold_time)
	fmt.Printf("total_time: %d,\n", simulation_parameters.total_time)
	fmt.Printf("speed: %d,\n", simulation_parameters.speed)
	fmt.Printf("opportunity: %d,\n", simulation_parameters.opportunity)
	fmt.Printf("distance: %d,\n", simulation_parameters.distance)
	fmt.Printf("record_distance: %d,\n", simulation_parameters.record_distance)
	fmt.Printf("\n")
}