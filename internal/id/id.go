// Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
// Package id makes globally unique, time-ordered identifiers (NFR13).
//
// Why hand-rolled: a UUID version 7 is sixteen bytes, six of them the time, and the
// standard library has everything needed. Pulling a module in for twenty lines would
// be a dependency the spec's Stack does not name.
package id

import (
	"crypto/rand"
	"fmt"
	"time"
)

// New returns a UUID version 7 as the usual 36-character string. The first 48 bits
// are Unix milliseconds, so sorting IDs as strings sorts records by creation time,
// which is what a forum wants, and no two installs can collide because the rest is
// random (74 bits of it).
func New() string {
	var b [16]byte
	ms := uint64(time.Now().UnixMilli())
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	b[5] = byte(ms)
	if _, err := rand.Read(b[6:]); err != nil {
		// crypto/rand failing means the OS has no entropy source at all; there is
		// nothing sensible to do but stop, and Go's own docs say the same.
		panic("id: crypto/rand failed: " + err.Error())
	}
	b[6] = (b[6] & 0x0f) | 0x70 // version 7 in the top nibble
	b[8] = (b[8] & 0x3f) | 0x80 // RFC 9562 variant in the top two bits
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
