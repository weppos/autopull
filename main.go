package main

import (
	"github.com/weppos/publicsuffix-go/publicsuffix/generator"
)

func main() {
	g := generator.NewGenerator()
	g.Verbose = false
	g.Print()
}
