package responses

import "time"


type OrderRes struct {
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CustomerName string    `json:"customer_name"`
	Items        []Item    `json:"Items"`
}

type Item struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"order_id"`
}
