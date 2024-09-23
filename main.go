package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: hexylon <hexy file> [hexy file...] <out file|->")
		return
	}

	inputFiles := os.Args[1 : len(os.Args)-1]
	outputFile := os.Args[len(os.Args)-1]

	out := os.Stdout

	if outputFile != "-" {
		var err error
		out, err = os.Create(outputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()
	}

	for _, inputFile := range inputFiles {
		in, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer in.Close()

		var s scanner.Scanner
		s.Init(in)
		s.Filename = inputFile

		for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
			if s.ErrorCount > 0 {
				return
			}

			if s.TokenText() == "," {
				continue
			}

			value := strings.TrimPrefix(strings.ToLower(s.TokenText()), "0x")
			if len(value) != 2 {
				fmt.Printf("%s: unexpected nibble length %d, expected 2: %s\n", s.Pos(), len(value), value)
				return
			}

			// validate values
			if !ValidateNibble(value[0]) {
				fmt.Printf("%s: first nibble is out of bounds: %d\n", s.Pos(), value[0])
				return
			}
			if !ValidateNibble(value[1]) {
				fmt.Printf("%s: second nibble is out of bounds: %d\n", s.Pos(), value[1])
				return
			}

			if _, err := out.Write([]byte{HexToBinary(value)}); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func ValidateNibble(value byte) bool {
	return (value >= '0' && value <= '9') || (value >= 'a' && value <= 'f')
}

func HexToBinary(value string) byte {
	var byt byte

	for i := 0; i < 2; i++ {
		if value[i] >= '0' && value[i] <= '9' {
			byt |= (value[i] - '0') << (4 * (1 - i))
		} else if value[i] >= 'a' && value[i] <= 'f' {
			byt |= (value[i] - 'a' + 10) << (4 * (1 - i))
		}
	}

	return byt
}
