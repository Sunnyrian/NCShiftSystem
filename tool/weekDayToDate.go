package tool

import (
	shiftConst "NCShiftSystem/const"
	"fmt"
	"time"
)

// WeekDayToDate 将周数 星期 转换为日期
func WeekDayToDate(week int, weekDay int) (d time.Time) {
	d, err := time.ParseInLocation(shiftConst.TimeForm, shiftConst.TermStartDate, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	d = d.Add(time.Hour * time.Duration(24*7*week))
	d = d.Add(time.Hour * time.Duration(24*weekDay))
	return d
}
