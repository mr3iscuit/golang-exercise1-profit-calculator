package cmd

import (
	"fmt"

	"github.com/mr3iscuit/profit-calc/logic"
	"github.com/mr3iscuit/profit-calc/state"
	"github.com/spf13/cobra"
)

var ebtCmd = &cobra.Command{
	Use:   "ebt",
	Short: "Calculate Earnings Before Tax (EBT) based on revenue, tax rate, and expenses.",
	Long:  `Calculates the Earnings Before Tax (EBT) using the provided revenue, tax rate, and expenses. EBT is a key profitability metric that shows a company's earnings before tax expenses are deducted.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEbt(cmd, args)
	},
}

func runEbt(cmd *cobra.Command, args []string) error {
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
}

func init() {
	RootCmd.AddCommand(ebtCmd)
}
