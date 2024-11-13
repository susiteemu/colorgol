package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	// Define some simple patterns and their associated colors
	errorColor := color.New(color.FgRed).SprintFunc()
	infoColor := color.New(color.FgWhite).SprintFunc()
	warnColor := color.New(color.FgYellow).SprintFunc()

	// Flag to track if we are in an error stack trace
	inErrorStackTrace := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			fmt.Println(errorColor(line))
			inErrorStackTrace = true // Start coloring subsequent lines red
		} else if strings.Contains(line, "INFO") {
			fmt.Println(infoColor(line))
			inErrorStackTrace = false // Reset flag
		} else if strings.Contains(line, "WARN") {
			fmt.Println(warnColor(line))
			inErrorStackTrace = false // Reset flag
		} else if inErrorStackTrace {
			// Color stack trace lines red if inside an error stack trace
			fmt.Println(errorColor(line))
		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}
