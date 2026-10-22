// Created: 2026-10-22 12:05 -05:00 · Edited: 2026-10-22 12:05 -05:00
// Command loomhall is the single binary a guild runs (FR2). For now it only says
// what it is; the version string arrives with C1-T3 and the server with F1.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "loomhall (development build)")
}
