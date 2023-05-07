package controller

import (
	"Tasya_Project-Mini/model"
	"net/http"
	"strconv"

	"Tasya_Project-Mini/usecase"

	"github.com/labstack/echo/v4"
)

func GetDrugcontroller(c echo.Context) error {
	drugs, e := usecase.GetListDrugs()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"drugs":  drugs,
	})
}

func GetDrugController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	drug, err := usecase.GetDrug(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get drug",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"drugs":  drug,
	})
}

// create new drug
func CreateDrugController(c echo.Context) error {
	drug := model.Drug{}
	c.Bind(&drug)

	if err := usecase.CreateDrug(&drug); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create drug",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new drug",
		"drug":    drug,
	})
}

// delete drug by id
func DeleteDrugController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteDrug(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete drug",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf obat tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete drug",
	})
}

// update drug by id
func UpdateDrugController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	drug := model.Drug{}
	c.Bind(&drug)
	drug.ID = uint(id)
	if err := usecase.UpdateDrug(&drug); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update drug",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf obat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update drug",
	})
}
