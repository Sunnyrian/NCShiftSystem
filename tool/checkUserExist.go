package tool

import (
	"NCShiftSystem/model"
)

// CheckUserExist 判断这个用户是否已经存在
func CheckUserExist(s string, key string) bool {
	db := model.ConnectToMySQL()
	var user model.User

	result := db.Where("? = ?", key, s).Find(&user)
	//fmt.Println(result.RowsAffected)
	//fmt.Println(&user)
	if result.RowsAffected == 0 {
		// 返回不存在
		return false
	} else {
		return true
	}
}