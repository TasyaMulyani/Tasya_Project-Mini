package controller

import (
	"Tasya_Project-Mini/model"
	"Tasya_Project-Mini/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSalesController(c echo.Context) error {
	sales, e := usecase.GetListSales()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"sales":  sales,
	})
}

func GetSaleController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	sale, err := usecase.GetSale(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get sale",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"sales":  sale,
	})
}

// create new sale
func CreateSaleController(c echo.Context) error {
	sale := model.Sale{}
	c.Bind(&sale)

	if err := usecase.CreateSale(&sale); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create sale",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new sale",
		"sale":    sale,
	})
}

// delete sale by id
func DeleteSaleController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteSale(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete sale",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf penjualan tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete sale",
	})
}

// update sale by id
func UpdateSaleController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	sale := model.Sale{}
	c.Bind(&sale)
	sale.ID = uint(id)
	if err := usecase.UpdateSale(&sale); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update sale",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf penjualan tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update sale",
	})
}
