package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/scanner"
)

var TokenScanner scanner.Scanner
var Line = 1
var TokenStr = [...]string{T_PLUS: "+", T_MINUS: "-", T_STAR: "*", T_SLASH: "/", T_INTLIT: "intlit"}

func scanfile() {
	t := Token{}

	for t.Scan() == 1 {
		fmt.Printf("Token %s", TokenStr[t.token])
		if t.token == T_INTLIT {
			fmt.Printf(", value %d", t.intvalue)
		}
		fmt.Printf("\n")
	}
}

func main()  {
	if len(os.Args) != 2 {
		err := errors.New("should specify the C filename")
		log.Fatal(err)
	} else {
		Infile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer Infile.Close()

		TokenScanner.Init(Infile)

		scanfile()
		//data := bufio.NewScanner(Infile)
		//data.Split(bufio.ScanRunes)
		//
		//for data.Scan() {
		//	fmt.Println(data.Text())
		//}
	}
}
