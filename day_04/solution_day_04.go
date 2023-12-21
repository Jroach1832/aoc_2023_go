package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	f, err := os.Open("input2.txt")
	
	if err != nil {
		log.Fatal(err)														
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	outer_count := 0.0
	for scanner.Scan() {
		s := scanner.Text()
		winning_list := generateWinningList(s)
		have_list := generateHaveList(s)
		inner_count := 0.0
		inner_total := 0.0
		for _, val := range have_list {
			if contains(winning_list, val) {
				inner_count++
				inner_total = math.Pow(2, inner_count - 1)
			}
		}
		outer_count += inner_total
	}

	fmt.Println(outer_count)

}

func generateWinningList(s string) []int {

	start := strings.Index(s, ":") + 2
	end := strings.Index(s, "|") - 1
	s = s[start:end]

	tokens := strings.Split(s, " ")

	nmb_tokens := make([]int, 0)

	for _, value := range tokens {

		if value == "" {
			continue
		}

		nmb_tokens = append(nmb_tokens, toInt(value))
	}
	
	return nmb_tokens
}

func generateHaveList(s string) []int {

	start := strings.Index(s, "|") + 2
	s = s[start:]

	tokens := strings.Split(s, " ")

	nmb_tokens := make([]int, 0)

	for _, value := range tokens {

		if value == "" {
			continue
		}

		nmb_tokens = append(nmb_tokens, toInt(value))
	}
	
	return nmb_tokens
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func contains(list []int, x int) bool {
	for _, value := range list {
		if value == x {
			return true
		}
	}
	return false
}