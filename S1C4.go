package main

import (
	"bufio"
	"fmt"
	"github.com/scripterio/go-cryptopals/bitwise"
	"github.com/scripterio/go-cryptopals/hexstring"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {

	file, _ := os.Open(path)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {

	lines := readLines("./data/4.txt")
	topscore := 0

	var topoutput []byte

	for _, line := range lines {

		raw := hexstring.Decode(line)
		bruteout := bitwise.BRUTEXOR(raw)
		score, key := hexstring.ScoreEnglish(bruteout)
		if score >= topscore {
			topscore = score
			topoutput = bruteout[key]
		}

	}

	output := string(topoutput)

	//Test
	fmt.Printf("Set 1   Challenge 4 \n")

	if output == "one day this nation will rise\n" {

		fmt.Printf("PASS!!\n\n")

	} else {
		fmt.Printf("##FAIL##\n\n")
	}

}
