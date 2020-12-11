package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear string
	IssuedYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportID string
	CountryId string
}

func main() {

	//partOne()
	partTwo()
}

func partTwo() {
	//Read Input File
	passwordsFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer passwordsFile.Close()
	scanner := bufio.NewScanner(passwordsFile)
	scanner.Split(ScanDoubleNewLine)

	validPassports := 0
	for scanner.Scan() {
		passportScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		passportScanner.Split(bufio.ScanWords)
		currentPassport := Passport{}
		for passportScanner.Scan() {
			currentPassport = parsePassportKey(passportScanner.Text(), currentPassport)
		}
		if isValidPassport(currentPassport)  {
			validPassports += 1
		}
	}
	fmt.Printf("Count of valid passports is %v \n", validPassports)
}

func partOne() {
	//Read Input File
	passwordsFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer passwordsFile.Close()
	scanner := bufio.NewScanner(passwordsFile)
	scanner.Split(ScanDoubleNewLine)

	validPassports := 0
	for scanner.Scan() {
		passportScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		passportScanner.Split(bufio.ScanWords)
		currentPassport := Passport{}
		for passportScanner.Scan() {
			currentPassport = parsePassportKey(passportScanner.Text(), currentPassport)
		}
		if isValidPassport(currentPassport)  {
			validPassports += 1
		}
	}
	fmt.Printf("Count of valid passports is %v \n", validPassports)
}

func isValidPassport(passport Passport) bool {
	fields := reflect.TypeOf(passport)
	values := reflect.ValueOf(passport)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		v := value.String()
		if !isPassPortFieldValid(field.Name, v) {
			if v != "" {
				//fmt.Printf("Invalid Field: %s, value: %s \n", field.Name, v)
			}
			return false
		}
		//fmt.Printf("Valid Field: %s, value: %s \n", field.Name, v)
	}
	return true
}

func isPassPortFieldValid(fieldName string, value string) bool {
	switch fieldName {
	case "BirthYear":
		if value == "" {
			return false
		}
		birthYear, _ := strconv.Atoi(value)
		return between(birthYear, 1920, 2002)
	case "IssuedYear":
		if value == "" {
			return false
		}
		issuedYear, _ := strconv.Atoi(value)
		return between(issuedYear, 2010, 2020)
	case "ExpirationYear":
		if value == "" {
			return false
		}
		ExpirationYear, _ := strconv.Atoi(value)
		return between(ExpirationYear, 2020, 2030)
	case "Height":
		re := regexp.MustCompile(`(?P<num>\d+)(?P<metric>cm|in)`)
		if re.MatchString(value) {
			matches := re.FindStringSubmatch(value)
			height, _ := strconv.Atoi(matches[re.SubexpIndex("num")])
			if matches[re.SubexpIndex("metric")] == "cm" {
				return between(height, 150, 193)
			} else if matches[re.SubexpIndex("metric")] == "in" {
				return between(height, 59, 76)
			} else {
				return false
			}
		}
		return false
	case "HairColor":
		if value == "" {
			return false
		}
		re := regexp.MustCompile(`#[0-9a-f]{6}`)
		return re.MatchString(value)
	case "EyeColor":
		if value == "" {
			return false
		}
		return isValidEyeColor(value)
	case "PassportID":
		re := regexp.MustCompile(`^\d{9}$`)
		return re.MatchString(value)
	case "CountryId":
		return true
	default:
		return false
	}
}

func parsePassportKey(passportKeyValue string, passport Passport) Passport{
	keyValueSplit := strings.Split(passportKeyValue, ":")
	switch keyValueSplit[0] {
	case "byr":
		passport.BirthYear = keyValueSplit[1]
		return passport
	case "iyr":
		passport.IssuedYear = keyValueSplit[1]
		return passport
	case "eyr":
		passport.ExpirationYear = keyValueSplit[1]
		return passport
	case "hgt":
		passport.Height = keyValueSplit[1]
		return passport
	case "hcl":
		passport.HairColor = keyValueSplit[1]
		return passport
	case "ecl":
		passport.EyeColor = keyValueSplit[1]
		return passport
	case "pid":
		passport.PassportID = keyValueSplit[1]
		return passport
	case "cid":
		passport.CountryId = keyValueSplit[1]
		return passport
	default:
		fmt.Printf("Invalid Key while parsing kv: %v\n" , keyValueSplit[0])
		return passport
	}
}

func isValidEyeColor(eyeColor string) bool {
	if len(eyeColor) != 3 {
		return false
	}
	validColors := "amb,blu,brn,gry,grn,hzl,oth"
	return strings.Contains(validColors, eyeColor)
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

func between (x int, min int, max int) bool {
	if (x >= min) && (x <= max) {
		return true
	}
	return false
}