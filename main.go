package test

import (
	"log"
	"strconv"
)

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
