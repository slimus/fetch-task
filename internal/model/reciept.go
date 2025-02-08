package model

import "time"

type Item struct {
	ShortDescription string
	Price            float64
}

type Reciept struct {
	ID           string
	Retailer     string
	PurchaseTime time.Time
	Items        []Item
	Total        float64
}
