package stats

import (
	"reflect"
	"testing"

	"github.com/Umar2505/11.C.1/v2/pkg/bank/types"
)

func TestPeriodsDynamic_positive(t *testing.T) {
	first:= map[types.Category]types.Money{
		"auto":10,
		"food":20,
		"clothes":10,
	}
	second:= map[types.Category]types.Money{
		"food":20,
		"auto":50,
	}
	result:= PeriodsDynamic(first,second)

	expected := map[types.Category]types.Money{
		"auto"		: 40,
		"clothes"	: -10,
		"food"		: 0,
	}

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}