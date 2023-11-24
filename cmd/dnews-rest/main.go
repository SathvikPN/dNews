package main

import (
	dnewsREST "github.com/SathvikPN/dNews/internal/dnews-rest"
)

// filled in by ldflags during build
var (
	Version string
)

func main() {
	dnewsREST.StartApplication(Version)
}
