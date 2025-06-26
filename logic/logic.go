package logic

func Ebt(revenue, taxRate, expenses float64) (float64, error) {
	return revenue - expenses, nil
}

func Profit(revenue float64, taxRate float64, expenses float64) (float64, error) {
	ebt, err := Ebt(revenue, taxRate, expenses)
	if err != nil {
		return 0, err
	}

	return ebt * (1 * taxRate / 100), nil
}

func Ratio(revenue, taxRate, expenses float64) (float64, error) {
	ebt, err := Ebt(revenue, taxRate, expenses)
	if err != nil {
		return 0, err
	}

	profit, err := Profit(revenue, taxRate, expenses)
	if err != nil {
		return 0, err
	}

	return ebt / profit, nil
}
