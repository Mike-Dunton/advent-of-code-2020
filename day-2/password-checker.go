package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	Min string
	Max string
	Required string
	Password string
}

func main() {

	partOne()
	//partTwo()
}

func partOne() {
	//Read Input File
	passwordsFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer passwordsFile.Close()
	scanner := bufio.NewScanner(passwordsFile)

	//var passwords []Password
	validPasswordCount := 0
	//\d+-\d+ \w: \w+
	re := regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+) (?P<requiredChar>\w): (?P<password>\w+)$`)
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			matches := re.FindStringSubmatch(scanner.Text())
			//passwords = append(passwords, Password{
			//	matches[re.SubexpIndex("min")],
			//	matches[re.SubexpIndex("max")],
			//	matches[re.SubexpIndex("requiredChar")],
			//	matches[re.SubexpIndex("password")],
			//})
			if isValid := isTobogganPasswordValid(Password{
					matches[re.SubexpIndex("min")],
					matches[re.SubexpIndex("max")],
					matches[re.SubexpIndex("requiredChar")],
					matches[re.SubexpIndex("password")],
				}); isValid {
				validPasswordCount++
			}
		}
	}
	fmt.Printf("There is %v valid passwords\n", validPasswordCount)

}

func partTwo() {
}

func isSledPasswordValid(password Password) bool {
	intMin, _ := strconv.Atoi(password.Min)
	intMax, _ := strconv.Atoi(password.Max)
	requiredCount := strings.Count(password.Password, password.Required)
	if (requiredCount >= intMin) && (requiredCount <= intMax) {
		return true
	} else {
		return false
	}
}

func isTobogganPasswordValid(password Password) bool {
	positionA, _ := strconv.Atoi(password.Min)
	positionB, _ := strconv.Atoi(password.Max)
	positionACharacter := string(password.Password[positionA-1])
	positionBCharacter := string(password.Password[positionB-1])
	if positionBCharacter != password.Required && positionACharacter != password.Required {
		return false
	}
	if positionACharacter == positionBCharacter {
		return false
	}
	return true
}