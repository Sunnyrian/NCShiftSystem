package shiftConst

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ModifyEnv(key string, value string)(status bool) {
	err := os.Setenv(key, value)
	if err != nil {
		log.Fatal(err)
	}

	var Env map[string]string
	Env, err = godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return false
	}
	Env[key]=value
	err = godotenv.Write(Env, "./.env")
	return true
}