package stats

import (
	"reflect"
	"testing"

	"github.com/Umar2505/11.C.1/v2/pkg/bank/types"
)

func TestCategoriesAvg_positive(t *testing.T) {
	payments:= []types.Payment{
		{
			Amount: 10_000,
			Category: "auto",
		},
		{
			Amount: 20_000,
			Category: "food",
		},
		{
			Amount: 50_000,
			Category: "cloth",
		},
		{
			Amount: 15_000,
			Category: "food",
		},
		{
			Amount: 200_000,
			Category: "cloth",
		},
	}
	result:= CategoriesAvg(payments)

	expected := map[types.Category]types.Money{
		"auto"	:	10_000,
		"cloth"	:	250_000,
		"food"	:	35_000,
	}

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}