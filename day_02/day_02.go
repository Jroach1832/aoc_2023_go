package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// close file
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		s := scanner.Text()
		gameId := getGameId(s)
		total += compare(gameId, tokenize(s), 12, 13, 14)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getGameId(s string) int {
	colonIndex := strings.Index(s, ":")
	return toInt(s[5:colonIndex])
}

func tokenizeByColon(s string) []string {

	//initial step
	start := strings.Index(s, ":") + 2
	s = s[start:]

	tokens := strings.Split(s, ";")
	
	return tokens
}

func tokenizeByComma(s string) []string {

	tokens := strings.Split(s, ",")
	
	return tokens
}

func tokenize(s string) [][]string {
	tokens := make([][]string, 0)
	level_one := tokenizeByColon(s)

	for _, val := range level_one {
		tokens = append(tokens, tokenizeByComma(val))
	}

	return tokens
} 

func getColorNumberPair(s string) (int, string){
	tokens := strings.Split(strings.TrimSpace(s), " ")
	return toInt(tokens[0]), tokens[1]
}

func compare(gameId int, tokens [][]string, red int, green int, blue int) int {
	for _, token := range tokens {
		for _, subtoken := range token {
			number, color := getColorNumberPair(subtoken)

			switch color {
			case "red":
				if number > red {
					return 0
				}
			case "green":
				if number > green {
					return 0
				}
			case "blue":
				if number > blue {
					return 0
				}
			}
		}
	}

	return gameId
}