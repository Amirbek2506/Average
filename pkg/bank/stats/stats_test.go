package stats

import (
	"fmt"
	"github.com/pkg/bank/types"
)

func StatsAvg()  {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "pay",
		},
		{
			ID:       2,
			Amount:   200,
			Category: "pay",
		},
	}

		fmt.Println(Avg(payments))

	// Output:
	// 150
}
