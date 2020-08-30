package main

import (
	"errors"
	"fmt"
	"github.com/EtoDemerzel0427/acwj/Parser"
	token "github.com/EtoDemerzel0427/acwj/Token"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		err := errors.New("should specify the C filename")
		log.Fatal(err)
	} else {
		Infile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer Infile.Close()

		ts := token.NewScanner(Infile)
		p := Parser.NewParser(ts)

		n := p.BinExpr()
		v, err := Parser.InterpretTree(n)
		if err == nil {
			fmt.Printf("%d\n", v)
		}

	}
}