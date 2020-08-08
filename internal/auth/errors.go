package auth

// UnauthenticatedError when user is not authenticated
type UnauthenticatedError struct{}

// Error error message
func (u *UnauthenticatedError) Error() string {
	return "Unauthenticated"
}
