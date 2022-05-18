package readXls

import (
	xlsConst "NCShiftSystem/const"
	"github.com/shakinm/xlsReader/xls"
	"gorm.io/gorm"
	"log"
)

// readSingleXls 读取单个Xls文件
func readSingleXls(userID int, na string, db *gorm.DB) {
	var filename = "./2022现任网管课表/" + na + ".xls"
	workbook, err := xls.OpenFile(filename)

	if err != nil {
		log.Panic(err.Error())
	}

	sheet, err := workbook.GetSheet(0)

	if err != nil {
		log.Panic(err.Error())
	}

	//最后一行不要
	for i := xlsConst.XlsStartRow; i < sheet.GetNumberRows()-1; i++ {
		if row, err := sheet.GetRow(i); err == nil {
			for j := 1; j <= xlsConst.DayOfWeek; j++ {
				if cell, err := row.GetCol(j); err == nil {
					// Cell value, string type
					//fmt.Println("\n\n============")
					//直接将 (i-3)为timePeriod (j-1)为星期几 week 传入数据库
					//fmt.Printf("这是第%d行,第%d列的数据", i - 3, j - 1)
					//fmt.Println(cell.GetString())
					//fmt.Println("============")
					var timePeriod = i - 3
					var weekday = j - 1
					ReadSingleCell(userID, timePeriod, weekday, cell.GetString(), db)

				}
			}
		}
	}
}
