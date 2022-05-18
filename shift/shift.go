package shift

import (
	"NCShiftSystem/model"
	"gorm.io/gorm"
)

// Shift 排班
/*  Shift 排班，给 指定时间段 所有没排的班次排上班
    首先查询到指定日期时间段的所有没有人值的班次
 	Shift 数据库中 Status {-1 表示还没有人值班}, {1 表示有人自己选了这个班次}
	那么我们要遍历这些班次，在每一次遍历中 查询出当前班次有空的人中，值班时长最少的网管，排上去
	过程： 拿到当前班次的week weekDay TimePeriod 查询空闲表中该时间段 Status = 0 的 UserID
	查询这些 UserID 的值班时长 升序 取第一个 UserID
	更新该班次 -> 遍历下一个班次
*/
func Shift(db *gorm.DB, startDate string, endDate string) {
	// 查询该时间段所有没有人值的班次
	unshift, line := selectUnshift(db, startDate, endDate)
	for i := int64(0); i < line; i++ {
		// 将该班次的时间转换 UNIX -> human
		//shiftDate, err := unshift[i].Date.Value()
		//if err != nil {
		//	fmt.Println(err)
		//}
		var SOT model.Occupation
		// 查询出来谁有空
		// 和 dutyHour 表做做一个左连接
		db.Select("user_id").Order("dutyhour.dutyhour ASC").Where("week = ? AND weekDay = ? AND time_period = ? AND Status = false", unshift[i].Week, unshift[i].Weekday, unshift[i].TimePeriod).Joins("left join dutyhour on dutyhour.id = occupation.user_id").Find(&SOT)
		unshift[i].UserID = SOT.UserID
		db.Save(&unshift[i])
	}
	//db.Save(&unshift)
}

//查询所有班次
//unshift := db.Find(&shift)

//测试输入的日期转化结果
//fmt.Println(tool.StringDateToGromDate(startDate).Value())
//fmt.Println(tool.StringDateToGromDate(endDate).Value())
