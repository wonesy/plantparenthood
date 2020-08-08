package auth

// AuthenticationError when user is not authenticated
type AuthenticationError struct{}

// Error error message
func (u *AuthenticationError) Error() string {
	return "Unauthenticated"
}
