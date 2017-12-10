package common

// GetWrappedIndex will return the next index from an array as it it was a loop
func GetWrappedIndex(currentPos int, loopsize int, modifier int) int {
	rawIndex := currentPos + modifier
	return rawIndex % loopsize
}
