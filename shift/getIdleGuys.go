package shift

import (
	"gorm.io/gorm"
)

func getIdleGuys(db *gorm.DB, week int, weekday int, timePeriod int) {
	//var idleGuys []model.User
	//var idleOT []model.Occupation
	//queryOT := db.Where("week = ? and weekday = ? and ? timePeriod = ? and status = 0", week, weekday, timePeriod).Find(&idleOT)
	//line := queryOT.RowsAffected
	//if queryOT.Error != nil {
	//	fmt.Println(queryOT.Error)
	//}
	//
	//fmt.Println(line)

}
