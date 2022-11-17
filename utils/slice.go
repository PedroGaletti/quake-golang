package utils

// Contains: This function it is to check if contains a string inside a string array
func Contains(array []string, str string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}

	return false
}
