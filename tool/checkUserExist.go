package tool

import (
	"NCShiftSystem/model"
	"fmt"
	"gorm.io/gorm"
)

// CheckUserExist 判断这个用户是否已经存在
func CheckUserExist(key string, val string) bool {
	if val == "" {
		return false
	}
	db := model.DB
	var user model.User
	// check 一个 User 无非也就 check 学号 nickname 电话 邮箱
	var result *gorm.DB
	switch key {
	case "nickname":
		result = db.Where("nickname = ?", val).Find(&user)
	case "stuID":
		result = db.Where("id = ?", val).Find(&user)
	case "telephone":
		result = db.Where("telephone = ?", val).Find(&user)
	case "email":
		result = db.Where("email = ?", val).Find(&user)
	default:
		fmt.Println("error param")
	}

	// 用下面这个方法
	//result := db.Where("? = ?", key, val).Find(&user)
	// SQL 语句为↓,查询不出结果，这里我们要对key处理一下
	// SELECT * FROM `user` WHERE 'nickname' = 'guochuankun';
	//
	//fmt.Println(result.RowsAffected)
	//fmt.Println(&user)
	if result.RowsAffected == 0 {
		// 返回不存在
		return false
	} else {
		return true
	}
}
