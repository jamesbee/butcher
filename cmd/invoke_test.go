package cmd

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_FileParse(t *testing.T) {
	file, err := os.Open("invoke.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
