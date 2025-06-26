package cmd

import (
	"fmt"

	"github.com/mr3iscuit/profit-calc/logic"
	"github.com/mr3iscuit/profit-calc/state"
	"github.com/spf13/cobra"
)

var profitCmd = &cobra.Command{
	Use:   "profit",
	Short: "Calculate net profit after tax and expenses.",
	Long:  `Calculates the net profit by subtracting tax and expenses from revenue. This command helps you determine the actual profit earned after all deductions.`,
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
	RootCmd.AddCommand(profitCmd)
}
