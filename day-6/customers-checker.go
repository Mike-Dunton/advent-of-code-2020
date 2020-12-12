package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	partOne()
	partTwo()
}

func partOne() {
	//Read Input File
	customsAnswers, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer customsAnswers.Close()
	scanner := bufio.NewScanner(customsAnswers)
	scanner.Split(ScanDoubleNewLine)
	allGroupsAnswers := make([]map[string]bool, 0)
	for scanner.Scan() {
		answerScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		answerScanner.Split(bufio.ScanBytes)
		currGroup := make(map[string]bool)
		for answerScanner.Scan() {
			re := regexp.MustCompile(`[a-z]{1}`)
			if re.MatchString(answerScanner.Text()) {
				currGroup[answerScanner.Text()] = true
			}
		}
		allGroupsAnswers = append(allGroupsAnswers, currGroup)
	}
	sum := 0
	for _, currGroup := range allGroupsAnswers {
		sum += len(currGroup)
	}
	fmt.Printf("sum %v\n", sum)
}

func partTwo() {
	//Read Input File
	customsAnswers, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer customsAnswers.Close()
	scanner := bufio.NewScanner(customsAnswers)
	scanner.Split(ScanDoubleNewLine)
	sum := 0
	for scanner.Scan() {
		answerScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		answerScanner.Split(bufio.ScanLines)
		currGroup := make(map[string]int)
		currentGroupSize := 0
		for answerScanner.Scan() {
			re:= regexp.MustCompile(`[a-z]`)
			if re.MatchString(answerScanner.Text()) {
				currentGroupSize++
			}
			indivScanner := bufio.NewScanner(strings.NewReader(answerScanner.Text()))
			indivScanner.Split(bufio.ScanBytes)
			for indivScanner.Scan() {
				re := regexp.MustCompile(`[a-z]{1}`)
				if re.MatchString(indivScanner.Text()) {
					currGroup[indivScanner.Text()]+=1
				}
			}
		}
		//fmt.Printf("currGroup: %v currentGroupSize: %v\n", currGroup, currentGroupSize)
		for _, yesAnswers := range currGroup {
			if yesAnswers == currentGroupSize {
				sum +=1
			}
		}
	}
	fmt.Printf("sum pt2 %v \n", sum)
}

func ScanDoubleNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n', '\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}