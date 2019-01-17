package tools

func IsValueInListString(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func IsValueInListInt(value int64, list []int64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
