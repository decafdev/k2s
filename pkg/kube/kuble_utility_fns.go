package kube

// Int32Ptr - a pointer to an int32 for use with kubernetes classes
func Int32Ptr(i int32) *int32 { return &i }

// Int64Ptr - a pointer to an int64 for use with kubernetes classes
func Int64Ptr(i int64) *int64 { return &i }

// BoolPtr - a pointer to an int64 for use with kubernetes classes
func BoolPtr(i bool) *bool { return &i }

// StringDefault - returns the first string with a non-empty value
func StringDefault(a ...string) string {
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			return a[i]
		}
	}
	return ""
}
