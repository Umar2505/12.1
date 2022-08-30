package stats

import (
	"github.com/Umar2505/11.C.1/v2/pkg/bank/types"
)

func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category]types.Money {
		result := map[types.Category]types.Money{}

		for v1, v2 := range first {
			result[v1]=second[v1]-v2
		}
		return result
}