package requests

import (
	"time"
)

type CreateOrder struct {
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

type UpdateOrder struct {
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
