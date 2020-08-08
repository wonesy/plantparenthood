package member

// AuthenticationError error type fo login errors
type AuthenticationError struct{}

// Error returns error message
func (a *AuthenticationError) Error() string {
	return "Incorrect credentials"
}
