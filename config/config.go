package config

import (
	"Tasya_Project-Mini/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// "DB_Username": "root",
// "DB_Password": "Qoala202!",
// "DB_Port":     "3306",
// "DB_Host":     "altadb.c3fzsopw9aux.ap-southeast-1.rds.amazonaws.com",
// "DB_Name":     "apotekdb",
func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "apotekdb",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Drug{}, &model.Sale{})
}
