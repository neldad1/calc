package calc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := initRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing the calc CLI '%s'", err)
		os.Exit(1)
	}
}
func initRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calc",
		Short: "calc - a simple CLI to perform add, subtract, and finding the min, max and mean.",
		Long:  `calc is a CLI that performs addition, subtraction and finding the min or max or mean from a given set of numbers`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	rootCmd.AddCommand(additionCommand())
	rootCmd.AddCommand(subtractionCommand())
	rootCmd.AddCommand(multiplicationCommand())
	rootCmd.AddCommand(divisionCommand())
	rootCmd.AddCommand(findMaxCommand())
	rootCmd.AddCommand(findMinCommand())
	rootCmd.AddCommand(findMeanCommand())
	return rootCmd
}
