package shift

import (
	shiftConst "NCShiftSystem/const"
	"NCShiftSystem/model"
	"gorm.io/gorm"
)
// InitOccupation 初始化每个人的OT表
func InitOccupation (db *gorm.DB) {
	// 这里先清空 OT 表
	// 应该启用事务，如果更新 OT表失败了 回滚删除操作
	db.Where("1=1").Delete(&model.Occupation{})
	// 每个人 17 个周 每周 7 天 每天 7 个时间段
	var users []model.User
	line := db.Find(&users).RowsAffected
	var OT []model.Occupation
	// 遍历所有网管 i
	for i := int64(0); i < line; i ++ {
		// 遍历所有周 j
		for j := 0; j < shiftConst.SchoolWeek; j ++ {
			// 遍历单周所有天 k
			for k := 0; k < shiftConst.DayOfWeek; k ++ {
				// 遍历一天中所有时间段 l
				for l := 0; l < shiftConst.DayTimePeriod; l ++ {
					OT = append(OT, model.Occupation{
						UserID:     users[i].ID,
						Week: 		j,
						Weekday:    k,
						TimePeriod: l,
						Status:     false,
					})
				}
			}
		}
	}
	db.Create(&OT)
}
