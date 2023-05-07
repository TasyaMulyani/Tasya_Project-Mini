package database

import (
	"Tasya_Project-Mini/config"
	"Tasya_Project-Mini/model"
)

func CreateSale(sale *model.Sale) error {
	if err := config.DB.Create(sale).Error; err != nil {
		return err
	}
	return nil
}

func GetSales() (sales []model.Sale, err error) {
	if err = config.DB.Find(&sales).Error; err != nil {
		return
	}
	return
}

func GetSale(id uint) (sale model.Sale, err error) {
	sale.ID = id
	if err = config.DB.First(&sale).Error; err != nil {
		return
	}
	return
}

func GetSalesByUserId(userID uint) (sale model.Sale, err error) {
	sale.UserID = userID
	if err = config.DB.Find(&sale).Error; err != nil {
		return
	}
	return
}

func UpdateSale(sale *model.Sale) error {
	if err := config.DB.Updates(sale).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSale(sale *model.Sale) error {
	if err := config.DB.Delete(sale).Error; err != nil {
		return err
	}
	return nil
}
