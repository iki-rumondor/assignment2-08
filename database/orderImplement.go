package database

import (
	"github.com/iki-rumondor/assignment2-GLNG-KS-08-08/domains"
)

func (p *postgresDB) FindAll() (*[]domains.Order, error) {

	var orders []domains.Order
	err := p.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (p *postgresDB) Find(id *int) (*domains.Order, error) {
	var order domains.Order
	err := p.db.Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (p *postgresDB) Save(o *domains.Order) error {
	return p.db.Create(o).Error
}

func (p *postgresDB) Update(o *domains.Order) (*domains.Order, error) {
	var order domains.Order
	if err := p.db.First(&order, "id = ?", o.Id).Error; err != nil {
		return nil, err
	}

	order.Items = o.Items

	err := p.db.Model(&order).Where("id = ?", o.Id).Updates(o).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (p *postgresDB) Delete(o *domains.Order) error {
	var order domains.Order
	if err := p.db.First(&order, "id = ?", o.Id).Error; err != nil {
		return err
	}

	err := p.db.Select("Items").Where("id = ?", o.Id).Delete(&o).Error
	if err != nil {
		return err
	}

	return nil
}
