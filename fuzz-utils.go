package testza

// FuzzUtilModifySet returns a modified version of a test set.
//
// Example:
//  modifiedSet := testza.FuzzUtilModifySet(testza.FuzzIntFull(), func(i int, value int) int {
//		return i * 2 // double every value in the test set
//	})
func FuzzUtilModifySet[setType any](inputSet []setType, modifier func(index int, value setType) setType) (floats []setType) {
	for i, input := range inputSet {
		floats = append(floats, modifier(i, input))
	}

	return
}
