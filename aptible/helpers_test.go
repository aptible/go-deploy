package aptible

import (
	"testing"
)

func TestGetIDFromHref(t *testing.T) {
	var tests = []struct {
		name     string
		href     string
		expected int64
		errored  bool
	}{
		{"test_vanilla", "https://api.aptible.com/disks/46", 46, false},
		{"test_conversion_err", "https://api.aptible.com/disks/str", 0, true},
		{"test_too_short", "https://api.aptible.com/2", 0, true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testName := tc.name
		t.Run(testName, func(t *testing.T) {

			id, err := GetIDFromHref(tc.href)

			if (err != nil) != tc.errored {
				t.Errorf("Input: %s caused error: %s", tc.href, err)
			}

			if id != tc.expected {
				t.Errorf("Input: %s should have resulted in id = %d. It was %d instead.", tc.href, tc.expected, id)
			}
		})
	}
}

func TestMakeStringSlice(t *testing.T) {
	var tests = []struct {
		name           string
		interfaceSlice []interface{}
		expected       []string
		errored        bool
	}{
		{"test_vanilla",
			[]interface{}{"chocolate", "vanilla", "strawberry"},
			[]string{"chocolate", "vanilla", "strawberry"},
			false},
		{"test_invalid_slice",
			[]interface{}{"chocolate", 24, ""},
			[]string{},
			true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testName := tc.name
		t.Run(testName, func(t *testing.T) {

			slice, err := MakeStringSlice(tc.interfaceSlice)

			if (err != nil) != tc.errored {
				t.Errorf("Input: %v caused error: %v", tc.interfaceSlice, err)
			}

			if !isEqual(slice, tc.expected) {
				t.Errorf("Input: %v should have resulted in slice = %v. It was %v instead.", tc.interfaceSlice, tc.expected, slice)
			}
		})
	}
}

// Tests if two string slices are equal
func isEqual(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
