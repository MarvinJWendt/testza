package testza

type BoolsHelper struct{}

// Full returns true and false in a boolean slice.
func (BoolsHelper) Full() []bool {
	return []bool{true, false}
}
