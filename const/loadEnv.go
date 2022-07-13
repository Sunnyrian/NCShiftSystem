package shiftConst

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv 每次程序运行的时候只有第一次加载环境变量有效,更改环境变量时先 os.setEnv,设置当前程序的环境变量，然后更新.env文件中的环境变量，
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version:", os.Getenv("VERSION"))
	fmt.Println("Database:", os.Getenv("DATABASE"))
}
