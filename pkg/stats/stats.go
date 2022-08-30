package stats

import (
	"github.com/Umar2505/11.C.1/v2/pkg/bank/types"
)

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}

	for _, v := range payments {
		result[v.Category]+=v.Amount
	}
	return result
}