package readXls

import (
	"NCShiftSystem/const"
	"NCShiftSystem/model"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

//对单个字符串切片判断周 POT Processed OT
func dealSingleCourse(userID int, timePeriod int, weekday int, s string, db *gorm.DB){
	//多种情况讨论  2-4[周] 1-17[单周] 1-17[双周] 10,12,14,16[周] 13[周] 1-10,12-17[周]
	//如果只有一个"-"，且没有逗号
	// 0 表示[周] 1 表示[单周] 2 表示[双周]

	//使用切片的方案
	//先用逗号切开字符串
	var SOT [shiftConst.SchoolWeek]bool
	weekSlice := strings.Split(s, ",")
	for _, v := range weekSlice {
		//如果是情况 3 或者 4
		if !strings.Contains(v, "-") {
			s1 := strings.Trim(v, "[周]")
			week, err := strconv.Atoi(s1)
			if err != nil {
				fmt.Println("can't convert to int")
			} else {
				//fmt.Println("week:", week)
				SOT[week-1] = true
			}
			//dealing
		} else {
			var xstartWeek = -1
			var xendWeek = -1
			var flag = -1
			weekSliceSlice := strings.Split(v, "-")
			for i, v := range weekSliceSlice {
				if i == 0 {
					s1 := v
					startWeek, err := strconv.Atoi(s1)
					if err != nil {
						fmt.Println("can't convert to int")
					} else {
						xstartWeek = startWeek
					}
				} else if i == 1 && strings.Contains(v, "[周]") {
					s2 := strings.Trim(v, "[周]")
					endWeek, err := strconv.Atoi(s2)
					if err != nil {
						fmt.Println("can't convert to int")
					} else {
						xendWeek = endWeek
						flag = 0
					}
				} else if i == 1 && strings.Contains(v, "[单周]") {
					s3 := strings.Trim(v, "[单周]")
					endWeek, err := strconv.Atoi(s3)
					if err != nil {
						fmt.Println("can't convert to int")
					} else {
						xendWeek = endWeek
						flag = 1
					}
				} else if i == 1 && strings.Contains(v, "[双周]") {
					s4 := strings.Trim(v, "[双周]")
					endWeek, err := strconv.Atoi(s4)
					if err != nil {
						fmt.Println("can't convert to int")
					} else {
						xendWeek = endWeek
						flag = 2
					}
				} else if i == 1 && !strings.Contains(v, "[周]") {
					s5 := v
					endWeek, err := strconv.Atoi(s5)
					if err != nil {
						fmt.Println("can't convert to int")
					} else {
						xendWeek = endWeek
						flag = 0
					}
				}
			}
			//fmt.Println("xstartWeek:", xstartWeek)
			//fmt.Println("xendWeek:", xendWeek)
			if flag == 0 {
				for i := xstartWeek; i <= xendWeek; i++ {
					SOT[i-1] = true
				}
			} else if flag == 1 {
				for i := xstartWeek; i <= xendWeek; i++ {
					if i%2 == 1 {
						SOT[i-1] = true
					}
				}
			} else if flag == 2 {
				//fmt.Println("我是双周")
				for i := xstartWeek; i <= xendWeek; i++ {
					//fmt.Println("i=",i,"i/2=",i/2)
					if i%2 == 0 {
						SOT[i-1] = true
					}
				}
			}
		}
	}
	/* 遍历SOT */
	for i := 0; i < 17; i++ {
		if SOT[i] {
			db.Model(&model.Occupation{}).Where("user_id = ? AND Week = ? AND Weekday = ? AND Time_Period = ?", userID, i, weekday, timePeriod).Update("Status", true)
		}
	}
}
