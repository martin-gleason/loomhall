// Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
// Command loomhall is the single binary a guild runs (FR2). For now it only says
// which version it is; the server arrives with F1.
package main

import (
	"fmt"
	"os"

	"github.com/martin-gleason/loomhall/internal/version"
)

func main() {
	fmt.Fprintln(os.Stdout, "loomhall "+version.String())
}
