package common

import "strings"

func CheckParams(params map[string]string, keys []string) bool {
	for _, k := range keys {
		if val, ok := params[k]; !ok || strings.TrimSpace(val) == "" {
			return false
		}
	}

	return true
}