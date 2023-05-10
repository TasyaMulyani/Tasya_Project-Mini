package route

import (
	"Tasya_Project-Mini/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	// drug collection
	drug := e.Group("/drugs")
	drug.GET("", controller.GetDrugController)
	drug.GET("/:id", controller.GetDrugController)
	drug.POST("", controller.CreateDrugController)
	drug.PUT("/:id", controller.UpdateDrugController)
	drug.DELETE("/:id", controller.DeleteDrugController)

	// sale collection
	sale := e.Group("/sales")
	sale.GET("", controller.GetSalesController)
	sale.GET("/:id", controller.GetSaleController)
	sale.POST("", controller.CreateSaleController)
	sale.PUT("/:id", controller.UpdateSaleController)
	sale.DELETE("/:id", controller.DeleteSaleController)

}
