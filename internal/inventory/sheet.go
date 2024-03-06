package inventory

import (
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string `gorm:"check:name_checker,name <> ''"`
	Quantity int    `gorm:"check:quantity_checker,quantity > 0"`
}

func AddItem(dbClient *gorm.DB, itemToRecord *Item) error {
	if err := dbClient.Create(&itemToRecord).Error; err != nil {
		return fmt.Errorf("AddItem failed to add record: %w", err)
	}
	return nil
}
