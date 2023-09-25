package repositories

import "github.com/iki-rumondor/assignment2-GLNG-KS-08-08/domains"

type OrderRepository interface {
	FindAll() (*[]domains.Order, error)
	Find(*int) (*domains.Order, error)
	Save(*domains.Order) error
	Update(*domains.Order) (*domains.Order, error)
	Delete(*domains.Order) error
}
