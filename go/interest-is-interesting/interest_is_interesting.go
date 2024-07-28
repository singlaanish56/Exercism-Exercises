package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	
	var ans float32

	switch {
	case balance < 0:
			ans =  float32(3.213)
	case balance>=0 && balance<1000:
			ans = float32(0.5)
	case balance>=1000 && balance<5000:
		ans = float32(1.621)
	case balance>=5000:
		ans = float32(2.475)
	}

	// switch {
	// case balance < 0:
	// 		ans = balance + (balance * float64(0.03213))
	// case balance>=0 && balance<1000:
	// 		ans = balance + (balance * float64(0.005))
	// case balance>=1000 && balance<5000:
	// 	ans = balance + (balance * float64(0.01621))
	// case balance>=5000:
	// 	ans = balance + (balance * float64(0.02475))
	// }

	return ans
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64((InterestRate(balance) / 100))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	if balance >= targetBalance{
		return 0
	}

	var y int

	for targetBalance > balance{
		balance=AnnualBalanceUpdate(balance)
		y++
	}

	return y
}
