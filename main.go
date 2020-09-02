package main

import (
	"fmt"
	"github.com/EtoDemerzel0427/acwj/Parser"
	token "github.com/EtoDemerzel0427/acwj/Token"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"strings"
)

var opt struct{
	Output string `short:"o" long:"output" default:""`
}

func main() {
	args, err := flags.Parse(&opt)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) == 0 {
		log.Fatal("No C filename specified.")
	}
	Infile := args[0]  // currently we only have one single infile.
	if opt.Output == "" {
		opt.Output = strings.Split(Infile, ".")[0] + ".s"  // currently output the assembly
	}

	fmt.Print(Infile + " " +  opt.Output + "\n")

	in, err := os.Open(Infile)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	ts := token.NewScanner(in)
	p := Parser.NewParser(ts)

	n := p.BinExpr(0)
	v, err := Parser.InterpretTree(n)

	if err == nil {
		fmt.Printf("%d\n", v)
	} else {
		log.Fatal(err)
	}



}