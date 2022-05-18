package readXls

import (
	"NCShiftSystem/tool"
	"gorm.io/gorm"
)

//处理所有 Course 的周切片产生 OT 表

// ReadSingleCell 循环遍历这个 Cell 中的每一门课程
func ReadSingleCell(userID int, timePeriod int, weekday int, s string, db *gorm.DB) {
	count, cellSlice := tool.SplitAndCount(s)
	//遍历 Slice 取出其中的周
	weekSlice := tool.PickupWeek(cellSlice, count)
	for _, v := range weekSlice {
		//处理之后，得到了这个Cell对应的 OT 表
		// 要初始化大家的空闲表先 (在另外一个方法中初始化)
		// 这里应该是更新空闲表
		dealSingleCourse(userID, timePeriod, weekday, v, db)
	}
}
