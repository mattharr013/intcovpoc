package strcheck

func IsCool(s string) bool {
	return s == "cool"
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	} else {
		return false
	}
}
