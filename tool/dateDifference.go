package tool

import (
	shiftConst "NCShiftSystem/const"
	"fmt"
	"math"
	"time"
)

func dateDifference (start string, end string) int{

	//转换开始时间为 Go 的时间  加上时区
	StartDate, err := time.ParseInLocation(shiftConst.TimeForm, start, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	//转换 s
	EndDate, err2 := time.ParseInLocation(shiftConst.TimeForm, end, time.Local)
	if err2 != nil {
		fmt.Println(err2)
	}

	//计算时间差多少天
	difference := EndDate.Sub(StartDate)
	floatDay := difference.Hours()/24
	day := int(math.Ceil(floatDay))
	return day
}