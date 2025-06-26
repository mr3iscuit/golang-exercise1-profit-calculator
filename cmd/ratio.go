package cmd

import (
	"fmt"

	"github.com/mr3iscuit/profit-calc/logic"
	"github.com/mr3iscuit/profit-calc/state"
	"github.com/spf13/cobra"
)

var ratioCmd = &cobra.Command{
	Use:   "ratio",
	Short: "Calculate profit ratio based on inputs.",
	Long:  `Computes the profit ratio using the provided revenue, tax rate, and expenses. The profit ratio is useful for analyzing profitability relative to revenue.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		calcState := state.CalculationStateInstanse()
		revenue := calcState.Inputs.Revenue
		taxRate := calcState.Inputs.TaxRate
		expenses := calcState.Inputs.Expenses

		profit, err := logic.Profit(revenue, taxRate, expenses)
		if err != nil {
			return err
		}

		fmt.Printf("Earnings Before Tax: %f\n", profit)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ratioCmd)
}
