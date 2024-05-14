package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const MEM_SIZE int = 128
const MAX_STEPS int = 10000

func main() {
	flag.Parse()
	args := flag.Args()

	var script string
	if len(args) != 0 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			log.Print("Failed to open file")
			log.Fatal(err)
		}
		script = string(file)
	} else {
		fmt.Print("script > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		script = scanner.Text()
	}

	fmt.Print("input  > ")
	brainfxxk(script, os.Stdin, os.Stdout)
}

func brainfxxk(script string, stdin *os.File, stdout *os.File) {
	var memory [MEM_SIZE]byte
	var cursor int = 0
	var pc int = 0
	var stepsCounter int = 0

	scanner := bufio.NewScanner(stdin)
	if !scanner.Scan() {
		log.Fatalf("Failed to read: %v", scanner.Err())
	} 
	input := scanner.Bytes()
	input = append(input, byte(26))

	defer func(){
		stdout.Write([]byte{10})
	}()

	fmt.Print("output > ")
	for {
		r := script[pc]
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
				if input[0] == byte(26) {
					return
				}
				memory[cursor] = input[0]
				input = input[1:]
			case '.':
				b := []byte{memory[cursor]}
				stdout.Write(b)
			case '[':
				if memory[cursor] == 0 {
					loopsCounter := 0
					idx := pc + 1
					for {
						if script[idx] == '['{
							loopsCounter++
						}
						if script[idx] == ']'{
							if loopsCounter == 0 {
								pc = idx
								break
							}
							loopsCounter--
						}
						idx++
					}
				}
			case ']':
				if memory[cursor] != 0 {
					loopsCounter := 0
					idx := pc - 1
					for {
						if script[idx] == ']'{
							loopsCounter++
						}
						if script[idx] == '['{
							if loopsCounter == 0 {
								pc = idx
								break
							}
							loopsCounter--
						}
						idx--
					}
				}
				
			default:
		}

		pc++
		stepsCounter++
		if cursor < 0 || cursor > MEM_SIZE {
			log.Fatal("cursor out of range!")
		}
		if stepsCounter > MAX_STEPS {
			log.Fatal("reaching maximum step!: ", MAX_STEPS)
		}
		if pc >= len(script) {
			break
		}
	}
}
