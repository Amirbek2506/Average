package stats

import "github.com/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var summa types.Money = 0
	for _, payment := range payments {
		summa+=payment.Amount
	}
	return summa/types.Money(len(payments))
}
