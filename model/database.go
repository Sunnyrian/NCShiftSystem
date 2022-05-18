package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMySQL() *gorm.DB {
	//dsn := "root:ZA102_Minecraft@tcp(100.102.100.6:3306)/shift?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:ZA102_Minecraft@tcp(100.88.1.4:3306)/shift?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}