package configs

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_ConfigureEnv_Failed(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(
				"ConfigureEnv는 에러를 발생시켜야 했습니다.",
			)
		}
	}()

	ConfigureEnv("wrong_sample.env")
}

func Test_ConfigureEnv_Success(t *testing.T) {
	filename:="sample.env"
	
	ConfigureEnv(filename)

	result, _ := godotenv.Read(filename)
	for key, value := range result {
		if configValue := os.Getenv(key); configValue != value {
			t.Error("ConfigureEnv 주입이 정상적으로 진행되지 않았습니다.")
		}
	}
}