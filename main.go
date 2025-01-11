package test

import (
	"fmt"
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

func GetLogMessage(m string) string {
	return fmt.Sprintf("LOG: %s", m)
}
