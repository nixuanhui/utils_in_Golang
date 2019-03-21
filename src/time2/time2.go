package time2

func IsGapYear(year int) bool {

	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	}

	return false
}
