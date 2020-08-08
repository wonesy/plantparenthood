package datetime

import "time"

// Validate returns whether a timestamp string is valid or not
func Validate(date string) error {
	_, err := time.Parse(time.RFC3339, date)
	return err
}

// Now returns a string format of time Now
func Now() string {
	return time.Now().Format(time.RFC3339)
}
