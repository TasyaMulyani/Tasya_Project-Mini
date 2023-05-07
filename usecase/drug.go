package usecase

import (
	"Tasya_Project-Mini/model"
	"Tasya_Project-Mini/repository/database"
	"errors"
	"fmt"
)

func CreateBook(drug *model.Drug) error {

	// check title cannot be empty
	if drug.Title == "" {
		return errors.New("drug title cannot be empty")
	}

	// check creator
	if drug.Creator == "" {
		return errors.New("drug creator cannot be empty")
	}

	err := database.CreateDrug(drug)
	if err != nil {
		return err
	}

	return nil
}

func GetDrug(id uint) (drug model.Drug, err error) {
	drug, err = database.GetDrug(id)
	if err != nil {
		fmt.Println("GetDrug: Error getting drug from database")
		return
	}
	return
}

func GetListDrugs() (drugs []model.Drug, err error) {
	drugs, err = database.GetDrugs()
	if err != nil {
		fmt.Println("GetListDrugs: Error getting drugs from database")
		return
	}
	return
}

func UpdateDrug(drug *model.Drug) (err error) {
	err = database.UpdateDrug(drug)
	if err != nil {
		fmt.Println("UpdateDrug : Error updating drug, err: ", err)
		return
	}

	return
}

func DeleteDrug(id uint) (err error) {
	drug := model.Drug{}
	drug.ID = id
	err = database.DeleteDrug(&drug)
	if err != nil {
		fmt.Println("DeleteDrug : error deleting drug, err: ", err)
		return
	}

	return
}
