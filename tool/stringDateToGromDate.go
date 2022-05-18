package tool

import (
	shiftConst "NCShiftSystem/const"
	"fmt"
	"gorm.io/datatypes"
	"time"
)

// StringDateToGromDate  输入字符串"2022-05-16" 转换为 Gorm 的datatype date 类型
func StringDateToGromDate(s string) (sDate datatypes.Date) {
	s += " 00:00:00"
	sTime, err := time.ParseInLocation(shiftConst.TimeForm, s, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	sDate = datatypes.Date(sTime)
	return sDate
}
