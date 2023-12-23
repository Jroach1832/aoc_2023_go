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
	destination_start int
	source_start int
	length int
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)

		chunk = append(chunk, s)

		if s == "" {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			continue
		} 

		//if line == 0 {
		//	fmt.Println(getSeedIds(s))
		//} 

		line++
	}

	fmt.Println(chunks)
	fmt.Println(chunks[1][1])
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getSeedIds(s string) []int {
	colon_index := strings.Index(s, ":")
	start := colon_index + 2
	seed_id_string := s[start:]
	string_seed_ids := strings.Split(seed_id_string, " ")
	numeric_seed_ids := make([]int, 0)

	for _, seed_id := range string_seed_ids {
		numeric_seed_ids = append(numeric_seed_ids, toInt(seed_id))
	}

	return numeric_seed_ids
}

func toMapRange(list []int) MapRange {
	map_range := MapRange {
		destination_start:  list[0],
		source_start: list[1],
		length: list[2],
	}

	return map_range
}

func generateMap(params []MapRange, empty_map map[int]int) {
	
}