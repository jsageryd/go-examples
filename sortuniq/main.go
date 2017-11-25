package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(scanLinesIncludingCR)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range sortUniq(lines) {
		fmt.Println(line)
	}
}

// scanLinesIncludingCR is bufio.ScanLines modified to not strip carriage
// returns.
func scanLinesIncludingCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
