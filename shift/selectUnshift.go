package shift

import (
	"NCShiftSystem/model"
	"NCShiftSystem/tool"
	"gorm.io/gorm"
)

// selectUnshift 查询该时间段所有没有人值的班次
func selectUnshift(db *gorm.DB, startDate string, endDate string) ([]model.Shift, int64) {
	var shift []model.Shift

	unshift := db.Where("user_id = 0 AND Date BETWEEN ? AND ?", tool.StringDateToGromDate(startDate), tool.StringDateToGromDate(endDate)).Find(&shift)
	line := unshift.RowsAffected

	return shift, line
}
