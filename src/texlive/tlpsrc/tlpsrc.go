package tlpsrc

import (
	"bufio"
	"fmt"
	"os"
)

type TLPSRC struct {
	Name        string
	Category    string
	Shortdesc   string
	Longdesc    string
	Catalogue   string
	Runpatterns []string
	Srcpatterns []string
	Docpatterns []string
	Binpatterns []string
	Postactions []string
	Executes    []string
	Depends     []string
}

func FromFile(fn string) TLPSRC {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		fmt.Println(s)
		s, e = Readln(r)
	}

	return TLPSRC{}
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
