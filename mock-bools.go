package testza

type MockInputsBoolsHelper struct{}

// Full returns true and false in a boolean slice.
func (MockInputsBoolsHelper) Full() []bool {
	return []bool{true, false}
}
