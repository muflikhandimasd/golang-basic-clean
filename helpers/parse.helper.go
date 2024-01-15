package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt(v interface{}) int {
	val, err := strconv.Atoi(fmt.Sprintf("%v", v))
	if err != nil {
		return 0
	}

	return val
}

func ParseInt64(v interface{}) int64 {
	val, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
	if err != nil {
		return 0
	}

	return val
}

func ParseFloat64(v interface{}) float64 {
	val, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	if err != nil {
		return 0
	}

	return val
}
