package calc

import (
	"fmt"

	"github.com/neldad1/calc/pkg/calc"
	"github.com/spf13/cobra"
)

func additionCommand() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add all the input numbers.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			sum, err := calc.Addition(args)
			if err == nil {
				fmt.Println("Sum: ", sum)
			} else {
				fmt.Println("Error: ", err)
			}
		},
	}
	return addCmd
}

func subtractionCommand() *cobra.Command {
	subtractCmd := &cobra.Command{
		Use:   "sub",
		Short: "Subtract two numbers.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			dif, err := calc.Subtraction(args[0], args[1])

			if err == nil {
				fmt.Println("Difference: ", dif)
			} else {
				fmt.Println("Error: ", err)
			}
		},
	}
	return subtractCmd
}

func multiplicationCommand() *cobra.Command {
	multiplyCmd := &cobra.Command{
		Use:   "mul",
		Short: "Multiply two positive numbers.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			mul, err := calc.Multiplication(args[0], args[1])

			if err == nil {
				fmt.Println("Product: ", mul)
			} else {
				fmt.Println("Error: ", err)
			}
		},
	}
	return multiplyCmd
}

func divisionCommand() *cobra.Command {
	divideCmd := &cobra.Command{
		Use:   "div",
		Short: "Divide two numbers.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			quotient, err := calc.Division(args[0], args[1])

			if err == nil {
				fmt.Println("Quotient: ", quotient)
			} else {
				fmt.Println("Error: ", err)
			}
		},
	}
	return divideCmd
}

func findMaxCommand() *cobra.Command {
	findMaxCmd := &cobra.Command{
		Use:   "max",
		Short: "Find the maximum number.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			max := calc.FindMax(args)
			fmt.Println("Maximum number is: ", max)
		},
	}
	return findMaxCmd
}

func findMinCommand() *cobra.Command {
	findMinCmd := &cobra.Command{
		Use:   "min",
		Short: "Find the minimum number.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			min := calc.FindMin(args)
			fmt.Println("Minimum number is: ", min)
		},
	}
	return findMinCmd
}

func findMeanCommand() *cobra.Command {
	findMeanCmd := &cobra.Command{
		Use:   "mean",
		Short: "Find the mean.",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			mean := calc.FindMean(args)
			fmt.Println("Mean is: ", mean)
		},
	}
	return findMeanCmd
}
