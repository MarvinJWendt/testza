package testza

// FuzzBoolFull returns true and false in a boolean slice.
func FuzzBoolFull() []bool {
	return []bool{true, false}
}
