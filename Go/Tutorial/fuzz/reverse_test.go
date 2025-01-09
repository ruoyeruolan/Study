package main

import (
	// "fmt"

	"testing"

	"unicode/utf8"
)


// func TestReverse(t *testing.T) {
// 	fmt.Println("Running unit tests")
// 	testcases := []struct {
// 		in, want string
// 	} {
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
//         {"!12345", "54321!"},
// 	}

// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }


func FuzzReverse(f *testing.F) {
	// fmt.Println("Running fuzz tests")
	testcases := []string{"Hello, world", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return 
		}

		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", 
		utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Invalid UTF-8: %q", rev)
		}
	})
}