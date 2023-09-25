package services

import (
	"github.com/iki-rumondor/assignment2-GLNG-KS-08-08/domains"
	"github.com/iki-rumondor/assignment2-GLNG-KS-08-08/repositories"
)

type OrderServices struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderServicve(repo repositories.OrderRepository) *OrderServices {
	return &OrderServices{OrderRepository: repo}
}

func (s *OrderServices) GetAllOrders() (*[]domains.Order, error) {
	orders, err := s.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderServices) GetOrder(id *int) (*domains.Order, error) {
	order, err := s.OrderRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServices) CreateOrder(order *domains.Order) error {
	if err := s.OrderRepository.Save(order); err != nil {
		return err
	}

	return nil
}

func (s *OrderServices) UpdateOrder(order *domains.Order) (*domains.Order, error) {
	order, err := s.OrderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServices) DeleteOrder(order *domains.Order) error {
	err := s.OrderRepository.Delete(order)
	if err != nil {
		return err
	}

	return nil
}
