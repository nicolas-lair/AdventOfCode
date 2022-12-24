package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadDataLineByLine(filepath string) *bufio.Scanner {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner
}

func ReadEntireFile(filepath string) string {
	f, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}
