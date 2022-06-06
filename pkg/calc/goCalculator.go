package calc

import (
	"fmt"
	"math"
	"strconv"

	"github.com/greenpau/go-calculator"
)

func initCalc(args []string) *calculator.Cell {
	var numbers []float64
	for _, a := range args {
		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			fmt.Println("failed to initialize calculator")
			return nil
		}
		numbers = append(numbers, i)
	}
	calc := calculator.New(numbers)

	return calc
}

func FindMax(args []string) float64 {
	calc := initCalc(args)
	if calc == nil {
		return -1
	}
	return calc.Max().Register.MaxValue
}

func FindMin(args []string) float64 {
	calc := initCalc(args)
	if calc == nil {
		return -1
	}
	return calc.Min().Register.MinValue
}

func FindMean(args []string) float64 {
	calc := initCalc(args)
	if calc == nil {
		return -1
	}
	val := math.Round(calc.Mean().Register.Mean*100) / 100
	return val
}
