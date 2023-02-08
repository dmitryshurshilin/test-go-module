package test

import (
	"log"
	"strconv"
)

func GetSomeTestValueFromModule() string {
	return "Some Test Value From Module"
}

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}

func ConvertStringToInt(value string) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
