package drivers

type jwtAuthManager struct {
}

func (jwt jwtAuthManager) check() bool  {
	// read token
	return false
}

func (jwt jwtAuthManager) user() bool  {
	// get model user
	return false
}

func (jwt jwtAuthManager) login() bool  {
	// write token into headers
	return false
}

func (jwt jwtAuthManager) logout() bool  {
	// del token
	return false
}