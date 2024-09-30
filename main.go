package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
    Y "cc/functions"
)

const MAX_NUMBERS = 1000

func main() {
    //check numbers of arguments
	if len(os.Args) != 2 {
		fmt.Println("\x1b[33mUsage:----->> go run . data.txt<------\033[0m")
		return
	}

    // Read file and affiche err if we can't 
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Fields(string(data))
	numbers := make([]float64, 0, MAX_NUMBERS)

	for _, line := range lines {
		num, err := strconv.ParseFloat(line, 64)
		if err == nil && len(numbers) < MAX_NUMBERS {
			numbers = append(numbers, num)
		}
	}

	if len(numbers) == 0 {
		fmt.Println("No valid numbers found in the file.")
		return
	}

	avg := Y.Average(numbers)
	median := Y.Median(numbers)
	variance := Y.Variance(numbers)
	stdDev := Y.StandardDev(numbers)

    // print all data with round it first
	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
