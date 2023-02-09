package utils

func CheckInStringArray(needle string, haystack *[]string) bool {
	for _, row := range *haystack {
		if needle == row {
			return true
		}
	}

	return false
}
