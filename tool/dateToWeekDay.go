package tool

import (
	shiftConst "NCShiftSystem/const"
)

// DateToWeekDay 输入时间 "2006-05-14"，转换为某个学期的第几周星期几
func DateToWeekDay(s string) (int, int) {
	s = s + " 00:00:00"
	day := dateDifference(shiftConst.TermStartDate, s)
	week := day / shiftConst.DayOfWeek
	weekDay := day % shiftConst.DayOfWeek
	return week, weekDay
}
