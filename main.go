package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var memory [128]int8
var cursor int8

func main(){
	flag.Parse()
	args := flag.Args()

	var script string
	if len(args) != 0 {
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			script += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	
}