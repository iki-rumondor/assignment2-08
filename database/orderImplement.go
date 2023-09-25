package database

import (
	"fmt"

	"github.com/iki-rumondor/assignment2-08/domains"
	"gorm.io/gorm"
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
	// Get order berdasarkan data model yang dikirim
	var order domains.Order
	if err := p.db.Preload("Items").First(&order, "id = ?", o.Id).Error; err != nil {
		return nil, err
	}

	// Simpan item yang dari body request ke variabel
	item := o.Items[0]
	item.OrderId = order.Id

	// Buat transaction untuk konsistensi data
	err := p.db.Transaction(func(tx *gorm.DB) error {
		// Update order function
		if err := tx.Model(&order).Where("id = ?", o.Id).Updates(o).Error; err != nil {
			return err
		}

		// Update or Insert condition
		var isUpdate = false
		for _, val := range order.Items {
			fmt.Println(val.ItemCode, item.ItemCode)
			if val.ItemCode == item.ItemCode {
				isUpdate = true
				break
			}
		}

		// If isUpdate
		if isUpdate {
			err := tx.Model(&item).Where("item_code = ?", item.ItemCode).Updates(o.Items[0]).Error
			if err != nil {
				return err
			}
			return nil
		}

		// If !isUpdate
		if err := tx.Create(&item).Error; err != nil {
			return err
		}

		return nil
	})

	// Jika transaction gagal kembalikan nilai error
	if err != nil {
		return nil, err
	}

	// Buat updatedOrder untuk dikirimkan sebagai response dari update function
	var updatedOrder domains.Order
	if err := p.db.Preload("Items").First(&updatedOrder, "id = ?", o.Id).Error; err != nil {
		return nil, err
	}

	return &updatedOrder, nil
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
