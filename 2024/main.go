package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	day_one "github.com/louisandrew/advent-of-code/2024/01"
	day_two "github.com/louisandrew/advent-of-code/2024/02"
	day_three "github.com/louisandrew/advent-of-code/2024/03"
	day_four "github.com/louisandrew/advent-of-code/2024/04"
	day_five "github.com/louisandrew/advent-of-code/2024/05"
	day_six "github.com/louisandrew/advent-of-code/2024/06"
	"github.com/louisandrew/advent-of-code/2024/utils"
)

var solutionsMap = map[utils.Day]utils.Solution{
	utils.DAY_ONE:   day_one.Solution,
	utils.DAY_TWO:   day_two.Solution,
	utils.DAY_THREE: day_three.Solution,
	utils.DAY_FOUR:  day_four.Solution,
	utils.DAY_FIVE:  day_five.Solution,
	utils.DAY_SIX:   day_six.Solution,
}

func runSolution(dayStr string, inputFilePath string) (int, error) {
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		return 0, fmt.Errorf("Error opening file: %v", err)
	}

	solution, ok := solutionsMap[utils.Day(dayStr)]
	if !ok {
		return 0, fmt.Errorf("No solution found for day %v", dayStr)
	}

	return solution(string(input)), nil
}

func main() {
	// You have to use the pointers when accessing. Otherwise it'll just use the default value.
	day := flag.Int("day", 1, "Day of advent (1-25)")
	test := flag.Bool("test", false, "Run test input instead of real input")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Println("Day must be between 1 and 25")
		os.Exit(1)
	}

	dayStr := fmt.Sprintf("%02d", *day)

	inputFileName := "input.txt"
	if *test {
		inputFileName = "test.txt"
	}

	inputFilePath := filepath.Join(dayStr, inputFileName)
	result, err := runSolution(dayStr, inputFilePath)
	if err != nil {
		fmt.Printf("Error running solution: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", result)
}
