package slice

// check is slice of string have one string with exact value
// linear search, should only be used for small slice
func InStrings(target string, elements []string) bool {
	for _, element := range elements {
		if element == target {
			return true
		}
	}
	return false
}
