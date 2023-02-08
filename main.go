package test

import (
	"strconv"
)

func GetSomeTestValueFromModule() string {
	return "Some Test Value From Module"
}

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}
