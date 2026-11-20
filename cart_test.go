package main

import "testing"

func TestTotal(t *testing.T) {
	items := []Item{{SKU: "mug", Price: 1250, Quantity: 2}, {SKU: "tee", Price: 2000, Quantity: 1}}
	for _, tc := range []struct{ pct, want int }{{0, 4500}, {10, 4050}, {100, 0}} {
		if got := Total(items, tc.pct); got != tc.want {
			t.Errorf("Total(%d%%) = %d, want %d", tc.pct, got, tc.want)
		}
	}
}
