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
