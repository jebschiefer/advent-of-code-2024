package testing_utils

import (
	"reflect"
	"testing"
)

func CompareInts(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func CompareBools(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func CompareStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func DeepCompareInts(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func DeepCompareStrings(t *testing.T, got []string, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
