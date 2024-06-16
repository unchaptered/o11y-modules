package configs

import "github.com/joho/godotenv"

func ConfigureEnv(filename string) {
	envError := godotenv.Load(filename)
	if envError != nil {
		panic("ConfigureEnv is failed : "+envError.Error())
	}
}