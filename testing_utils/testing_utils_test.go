package testing_utils

import "testing"

func TestCompareInts(t *testing.T) {
	CompareInts(t, 1, 1)
}

func TestCompareBools(t *testing.T) {
	CompareBools(t, true, true)
	CompareBools(t, false, false)
}

func TestCompareStrings(t *testing.T) {
	CompareStrings(t, "string", "string")
}

func TestDeepCompareInts(t *testing.T) {
	DeepCompareInts(t, []int{1, 2}, []int{1, 2})
}

func TestDeepCompareStrings(t *testing.T) {
	DeepCompareStrings(t, []string{"one", "two"}, []string{"one", "two"})
}
