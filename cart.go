// Package main is a small checkout service.
package main

// Item is one line of a cart: a price in cents and a quantity.
type Item struct {
	SKU      string `json:"sku"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

// Subtotal is the sum of every line's price times its quantity, in cents.
func Subtotal(items []Item) int {
	total := 0
	for _, it := range items {
		total += it.Price * it.Quantity
	}
	return total
}

// Total applies a percentage discount to the subtotal, rounding down to the cent.
func Total(items []Item, discountPct int) int {
	sub := Subtotal(items)
	return sub * (100 - discountPct) / 100
}
