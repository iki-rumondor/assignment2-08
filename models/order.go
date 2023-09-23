package models

import "time"

type Order struct{
	Id int `gorm:"primaryKey"`
	CustomerName string `gorm:"not null; type:varchar(120)"`
	Items []Item

	CreatedAt time.Time
	UpdatedAt time.Time
}