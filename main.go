package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

const MEM_SIZE int = 128

func main(){
	flag.Parse()
	args := flag.Args()

	var script string
	var scanner *bufio.Scanner
	if len(args) != 0 {
		file, err := os.Open(args[0])
		if err != nil {
			log.Print("Failed to open file")
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		script += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Print("Failed to scan file")
		log.Fatal(err)
	}

	brainfxxk(script, os.Stdin, os.Stdout)
}

func brainfxxk(script string, stdin *os.File, stdout *os.File) {
	var memory [MEM_SIZE]byte
	var cursor int = 0
	for _, r := range script {
		if cursor < 0 || cursor > MEM_SIZE {
			log.Fatal("cursor out of range!")
		}
		switch r {
			case '>':
				cursor++
			case '<':
				cursor--
			case '+':
				memory[cursor]++
			case '-':
				memory[cursor]--
			case ',':
				b := []byte{0}
				_, err := stdin.Read(b)
				if err != nil {
					log.Fatal(err)
				}
				memory[cursor] = b[0]
			case '.':
				b := []byte{memory[cursor]}
				stdout.Write(b)
			default:
				continue
		}
	}
	log.Println(memory)
	log.Println(cursor)
}