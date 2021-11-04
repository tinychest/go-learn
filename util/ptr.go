package util

// from github.com/xuri/excelize/v2@v2.4.1/styles.go:2193

// BoolPtr returns a pointer to a bool with the given value.
func BoolPtr(b bool) *bool { return &b }

// IntPtr returns a pointer to a int with the given value.
func IntPtr(i int) *int { return &i }

// Float64Ptr returns a pofloat64er to a float64 with the given value.
func Float64Ptr(f float64) *float64 { return &f }

// StringPtr returns a pointer to a string with the given value.
func StringPtr(s string) *string { return &s }
