package leap

// IsLeapYear returns if the given year is a leap year.
func IsLeapYear(yr int) bool {
	if yr%400 == 0 {
		return true
	}
	if yr%100 == 0 {
		return false
	}
	if yr%4 == 0 {
		return true
	}
	return false
}
