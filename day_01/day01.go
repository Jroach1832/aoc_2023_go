package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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
		first := getFirstDigit(s)
		last := getLastDigit(s)
		result := toInt(first + last)
		total += result
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func getFirstDigit(s string) string {
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			return string(ch)
		}
	}
	return ""
}

func getLastDigit(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return string(s[i])
		}
	}
	return ""
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
