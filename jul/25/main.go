package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

     "github.com/Gealber/adventofcode/utils"
)

const (
	filename = "input.txt"
)

func main() {
    acc := 0
	for _, line := range retrieveInput() {
	    acc += utils.FromSNAFU(line)
	}

    fmt.Println(acc)
    snafu := utils.ToSNAFU(acc)
    fmt.Println(snafu)
}

func retrieveInput() []string {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil
	}

	result := make([]string, 0)
	for _, lineb := range bytes.Split(b, []byte{'\n'}) {
		result = append(result, string(lineb))
	}

	return result
}
