package main

import (
	"fmt"

	"github.com/yggdrasildev/seer/third_party/swephgo"
)

func main() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", sweVer)
}
