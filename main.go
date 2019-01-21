package main

import "fmt"
import "texlive/tlpsrc"

func main() {
	// a := tlpsrc.TLPSRC{Name: "12many", Category: "TLCore"}
	// fmt.Println(a)
	// fmt.Println(a.Executes)
	b := tlpsrc.FromFile("zapfding.tlpsrc")
	fmt.Println(b)
}
