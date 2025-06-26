package state

import "sync"

var (
	stateInstance *CalculationState
	once          sync.Once
)

type CalculationState struct {
	Inputs Inputs
}

type Inputs struct {
	Revenue  float64
	TaxRate  float64
	Expenses float64
}

func CalculationStateInstanse() *CalculationState {
	once.Do(func() {
		stateInstance = &CalculationState{
			Inputs: Inputs{
				0,
				0,
				0,
			},
		}
	})

	return stateInstance
}
