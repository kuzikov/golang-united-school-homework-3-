package homework

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []int64{1, 2, 5, 15}

	expected := []int64{15, 5, 2, 1}
	actual := reverse(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual:\t%v\nExpected\t%v\n", actual, expected)
	}
}
