package readXls

import (
	"NCShiftSystem/model"
	"fmt"
	"gorm.io/gorm"
)



func ReadAllXls(db *gorm.DB){
	var users []model.User
	line := db.Find(&users).RowsAffected
	fmt.Println(line)
	for i := int64(0); i < line; i++{
		fmt.Println(users[i].ID,users[i].Name)
		readSingleXls(users[i].ID, users[i].Name, db)
	}
}





