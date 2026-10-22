// Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
package id

import (
	"regexp"
	"sort"
	"testing"
	"time"
)

// NFR13: every record carries a globally unique ID.
var shape = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestNew_IsAVersion7UUID(t *testing.T) {
	got := New()
	if !shape.MatchString(got) {
		t.Fatalf("New() = %q, not a version-7 UUID", got)
	}
}

func TestNew_NeverRepeats(t *testing.T) {
	seen := make(map[string]struct{}, 20000)
	for i := 0; i < 20000; i++ {
		s := New()
		if _, dup := seen[s]; dup {
			t.Fatalf("duplicate id after %d calls: %s", i, s)
		}
		seen[s] = struct{}{}
	}
}

func TestNew_SortsByCreationTime(t *testing.T) {
	first := New()
	time.Sleep(2 * time.Millisecond)
	second := New()
	ids := []string{second, first}
	sort.Strings(ids)
	if ids[0] != first {
		t.Fatalf("ids do not sort by time: %q should precede %q", first, second)
	}
}
