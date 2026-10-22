// Created: 2026-10-22 11:38 -05:00 · Edited: 2026-10-22 11:38 -05:00
package version

import "testing"

// C1-T3, owner: human. This test was written first; version.go is Marty's to write.
// It needs two things in package version:
//
//	const Version = "0.0.1"      // any non-empty string in the form major.minor.patch
//	func String() string         // returns Version
func TestString_ReturnsTheVersion(t *testing.T) {
	if String() == "" {
		t.Fatal("String() returned an empty string; it should return Version")
	}
	if String() != Version {
		t.Fatalf("String() = %q, want Version = %q", String(), Version)
	}
}

func TestVersion_LooksLikeMajorMinorPatch(t *testing.T) {
	dots := 0
	for _, c := range Version {
		switch {
		case c == '.':
			dots++
		case c < '0' || c > '9':
			t.Fatalf("Version = %q contains %q; expected digits and dots only", Version, c)
		}
	}
	if dots != 2 {
		t.Fatalf("Version = %q has %d dots, want 2 (major.minor.patch)", Version, dots)
	}
}
