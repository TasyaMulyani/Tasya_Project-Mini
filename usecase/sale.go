package usecase

import (
	"Tasya_Project-Mini/model"
	"Tasya_Project-Mini/repository/database"
	"errors"
	"fmt"
)

type SaleUsecase interface {
	CreateSale(sale *model.Sale) error
	GetSale(id uint) (sale model.Sale, err error)
	GetListSales() (sales []model.Sale, err error)
	UpdateSale(sale *model.Sale) (err error)
	DeleteSale(id uint) (err error)
}

func CreateSale(sale *model.Sale) error {

	// check status cannot be empty
	if sale.Status == "" {
		return errors.New("sale status cannot be empty")
	}

	// check qty
	if sale.Qty == "" {
		return errors.New("sale qty cannot be empty")
	}

	// check pay
	if sale.Pay == "" {
		return errors.New("sale pay cannot be empty")
	}

	err := database.CreateSale(sale)
	if err != nil {
		return err
	}

	return nil
}

func GetSale(id uint) (sale model.Sale, err error) {
	sale, err = database.GetSale(id)
	if err != nil {
		fmt.Println("GetSale: Error getting sale from database")
		return
	}
	return
}

func GetListSales() (sales []model.Sale, err error) {
	sales, err = database.GetSales()
	if err != nil {
		fmt.Println("GetListSales: Error getting sales from database")
		return
	}
	return
}

func UpdateSale(sale *model.Sale) (err error) {
	err = database.UpdateSale(sale)
	if err != nil {
		fmt.Println("UpdateSale : Error updating sale, err: ", err)
		return
	}

	return
}

func DeleteSale(id uint) (err error) {
	sale := model.Sale{}
	sale.ID = id
	err = database.DeleteSale(&sale)
	if err != nil {
		fmt.Println("DeleteSale : error deleting sale, err: ", err)
		return
	}

	return
}
