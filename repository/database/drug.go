package database

import (
	"Tasya_Project-Mini/config"
	"Tasya_Project-Mini/model"
)

func CreateDrug(drug *model.Drug) error {
	if err := config.DB.Create(drug).Error; err != nil {
		return err
	}
	return nil
}

func GetDrugs() (drugs []model.Drug, err error) {
	if err = config.DB.Find(&drugs).Error; err != nil {
		return
	}
	return
}

func GetDrug(id uint) (drug model.Drug, err error) {
	drug.ID = id
	if err = config.DB.First(&drug).Error; err != nil {
		return
	}
	return
}

func UpdateDrug(drug *model.Drug) error {
	if err := config.DB.Updates(drug).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDrug(drug *model.Drug) error {
	if err := config.DB.Delete(drug).Error; err != nil {
		return err
	}
	return nil
}
