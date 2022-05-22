package homework

import "testing"

func TestAverage(t *testing.T) {
	var target = [15]float32{1, 3, 4, 5, 6, 2, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	res := average(target)
	if res != 8.0 {
		t.Errorf("want: 8.0, have: %f", res)
	}
}
