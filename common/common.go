package common

import (
	"strings"
	"strconv"
	"fmt"
)

func CheckParams(params map[string]string, keys []string) bool {
	for _, k := range keys {
		if val, ok := params[k]; !ok || strings.TrimSpace(val) == "" {
			return false
		}
	}

	return true
}

func Str2Int(s string) int {
	s1, _ := strconv.ParseInt(s, 10, 64)

	return int(s1)
}

func Str2Float(s string) float64 {
	s1, _ := strconv.ParseFloat(s, 10)

	return s1
}

func Float2Money(v float64) int {
	return int(v * 100)
}

func Money2Float(v int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v) / 100), 64)
	return value
}
