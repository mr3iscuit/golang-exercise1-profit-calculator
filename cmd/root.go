package cmd

import (
	"fmt"
	"os"

	"github.com/mr3iscuit/golang-exercise1-profit-calculator/state"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:          getCommandName(),
	Short:        "Profit Calculator CLI – Calculate profit, EBT, and ratios from revenue, tax, and expenses.",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()

		revenue, err := flags.GetFloat64("revenue")
		if err != nil {
			return err
		}

		taxRate, err := flags.GetFloat64("tax-rate")
		if err != nil {
			return err
		}

		expenses, err := flags.GetFloat64("expenses")
		if err != nil {
			return err
		}

		if revenue < 0 {
			return fmt.Errorf("revenue must be non-negative")
		}
		if expenses < 0 {
			return fmt.Errorf("expenses must be non-negative")
		}
		if taxRate < 0 || taxRate > 100 {
			return fmt.Errorf("tax-rate must be between 0 and 100 (inclusive)")
		}

		calcState := state.CalculationStateInstanse()
		calcState.Inputs = state.Inputs{
			Revenue:  revenue,
			TaxRate:  taxRate,
			Expenses: expenses,
		}

		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func getCommandName() string {
	return os.Args[0]
}

func init() {
	flags := RootCmd.PersistentFlags()

	flags.Float64P("revenue", "r", 0, "Total revenue (income) amount")
	flags.Float64P("tax-rate", "t", 0, "Tax rate as a percentage (e.g., 15 for 15%)")
	flags.Float64P("expenses", "e", 0, "Total expenses amount")
}
