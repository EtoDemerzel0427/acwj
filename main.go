package main

import (
	"fmt"
	token "github.com/EtoDemerzel0427/acwj/Token"
	"os"
)

func main() {
	ts := token.NewScanner(os.Stdin)  // currently use gocc < filename.c to avoid file manipulating.
	for {
		num, err := ts.Scan()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		if num != 1 || err != nil {
			break
		}

		fmt.Print(ts.Tok)
	}
	fmt.Print("\n")
}