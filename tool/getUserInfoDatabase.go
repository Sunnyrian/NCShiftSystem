package tool

import "NCShiftSystem/model"

func GetUserInfo(stuID string) model.User{
	db := model.DB
	var user model.User
	db.Where("stuid = ?", stuID).Find(&user)
	return user
}