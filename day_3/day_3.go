package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"unicode"
	"strconv"
)

type PartNumber struct {
    Row   int
    Cols  []int
    Value int
}

type Symbol struct {
	Row int
	Col int
}

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

	row := 0

	part_numbers := make([]PartNumber, 0)
	symbols := make([]Symbol, 0)

	for scanner.Scan() {
		s := scanner.Text()
		//fmt.Println(s)

		for index, character := range s {
			schar := string(character)
			cond := false

			if index != 0 {
				cond = unicode.IsNumber(rune(s[index - 1])) && !isSpecialCharacter(rune(s[index]))
			}

			if schar == "." || cond {
				continue
			} else if unicode.IsNumber(character) {

				tmp_value := grabNumber(s[index:])
				tmp_str_value := strconv.Itoa(tmp_value)
				list_of_indices := grabListOfIndices(tmp_str_value, index)

				tmp_pt_nmb := PartNumber{
					Row: row, 
					Cols: list_of_indices, 
					Value: tmp_value,
				}
				part_numbers = append(part_numbers, tmp_pt_nmb)
			
			} else {
				tmp_symbol := Symbol{
					Row: row,
					Col: index,
				}

				symbols = append(symbols, tmp_symbol)
			}
		}

		row++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	/*fmt.Println()
	fmt.Println("Part numbers:", part_numbers)
	fmt.Println()
	fmt.Println("Symbols:", symbols)
	fmt.Println()*/

	total := 0

	for _, symbol_value := range symbols {
		for _, part_value := range part_numbers {
			row_cond := checkRowCondition(symbol_value.Row, part_value.Row)
			col_cond := checkColumnsCondition(symbol_value.Col, part_value.Cols)

			if row_cond && col_cond {
				total += part_value.Value
			}

		}
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

func grabNumber(s string) int {

	n := make([]rune, 0)

	for _, character := range s {
		schar := character
		if schar == 46 || !unicode.IsNumber(character) {
			break
		} else {
			n = append(n, schar)
		}
	}

	return toInt(string(n))
}

func grabListOfIndices(s string, start_index int) []int {
	list_of_indices := make([]int, 0)
	for index, _ := range s {
		list_of_indices = append(list_of_indices, start_index + index)
	}

	return list_of_indices
}

func isSpecialCharacter(char rune) bool {
	return !unicode.IsLetter(char) && !unicode.IsNumber(char)
}

func checkRowCondition(symbol_row int, part_row int) bool {
	return part_row >= symbol_row - 1 && part_row <= symbol_row + 1
}

func checkColumnsCondition(symbol_column int, part_columns []int) bool {
	conds := make([]bool, 0)

	for _, part_column := range part_columns {
		conds = append(conds, checkRowCondition(symbol_column, part_column))
	}

	return any(conds)
}

func any(conds []bool) bool {
	match := false

	for i := 0; i < len(conds); i++ {
		if conds[i] {
			match := true
			return match
		}
	}
	
	return match
}