package env

import (
	"log"
	"os"
	"strconv"
)

func GetPassword() string {
	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatalln("PASSWORD is required")
		os.Exit(1)
	}
	return password
}

func GetLoginTime() int {
	rawTime := os.Getenv("LOGIN_TIME")
	time, err := strconv.Atoi(rawTime)
	if err != nil {
		log.Fatalln("LOGIN_TIME is in wrong format")
		os.Exit(1)
	}
	return time
}

func GetServiceAddress() string {
	serviceAddress := os.Getenv("SERVICE_ADDRESS")
	if serviceAddress == "" {
		log.Fatalln("SERVICE_ADDRESS is required")
		os.Exit(1)
	}
	return serviceAddress
}
