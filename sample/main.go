package main

import (
	"fmt"

	"yggdrasildev.io/third_party/swephgo"
)

func main() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", sweVer)
}
