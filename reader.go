package onion

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s is does not exist\n", filename)
	}

	reader := bufio.NewReaderSize(file, 4096)
	var result []string
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	return result
}
