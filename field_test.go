package awql

import (
	"reflect"
	"testing"
)

func TestGetFields(t *testing.T) {
	want := []struct {
		q string
		f []string
	}{
		{"select one, two, three, four from test", []string{"one", "two", "three", "four"}},
		{"select one, two, three,four order by test", []string{"one", "two", "three", "four"}},
		{"select one, TWO, three,four limit 1", []string{"one", "TWO", "three", "four"}},
	}

	for _, w := range want {
		f := getFields(w.q)
		if !reflect.DeepEqual(f, w.f) {
			t.Errorf("want %v got %v", w.f, f)
		}
	}
}
