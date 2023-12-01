package day01

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	partFlag := flag.String("part", "p",
		"Specify 'a' or 'b' for the part. If none print both answers")
	// Parse the command-line flags
	flag.Parse()
	// Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer input.Close()

	switch *partFlag {
	case "a":
		fmt.Printf("Answer A: %s", getAnswerA(input))
	case "b":
		fmt.Printf("Answer B: %s", getAnswerB(input))
	default:
		fmt.Printf("Answer A: %s\n Answer B: %s", getAnswerA(input), getAnswerB(input))
	}
}

func getAnswerA(input *os.File) string {
	return "answer A"
}

func getAnswerB(input *os.File) string {
	return "answer B"
}
