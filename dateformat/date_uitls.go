package dateformat

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow returns current UTC time format
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns time format in this "2006-01-02T15:04:05Z"
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns time format in this "2006-01-02 15:04:05"
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
