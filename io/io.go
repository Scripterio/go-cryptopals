package io

import (
	"bufio"
	"bytes"
	"os"
)

func Read(path string) string {

	file, _ := os.Open(path)

	defer file.Close()

	var buffer bytes.Buffer

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		buffer.WriteString(scanner.Text())
	}

	return buffer.String()
}
