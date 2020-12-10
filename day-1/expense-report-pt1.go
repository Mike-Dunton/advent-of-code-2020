package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	partOne()
	partTwo()
}

func partOne() {
	//Read Input File
	expenseReportRaw, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer expenseReportRaw.Close()
	scanner := bufio.NewScanner(expenseReportRaw)

	// Find Pair that sums to 2020
	sum := 2020
	expenses := make(map[int]int)


	// As we iterate of the list of expenses check if the sum minus current number has been seen yet.
	// If the sum - current number exists in the expenses hashMap then the currentNumber and found number sum to the sum
	// Associative Property For Addition
	//The sum of two or more real numbers is always the same regardless of how you group them. When you add real numbers, any change in their grouping does not affect the sum.
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error converting scanned report item: %s  err: %v \n", scanner.Text(), err)
		}
		tempSum := sum - expense
		if _, ok := expenses[tempSum]; ok {
			fmt.Printf("Found Matching Pair %v + %v = %v \n", expense, tempSum, sum)
			fmt.Printf("Matching Pair mulitpled %v * %v = %v \n", expense, tempSum, expense*tempSum)
		}
		expenses[expense]=expense
	}
}

func partTwo() {
	//Read Input File
	expenseReportRaw, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer expenseReportRaw.Close()
	scanner := bufio.NewScanner(expenseReportRaw)

	// Find Pair that sums to 2020
	sum := 2020
	expenses := make(map[int]int)


	// As we iterate of the list of expenses check if the sum minus current number has been seen yet.
	// If the sum - current number exists in the expenses hashMap then the currentNumber and found number sum to the sum
	// Associative Property For Addition
	//The sum of two or more real numbers is always the same regardless of how you group them. When you add real numbers, any change in their grouping does not affect the sum.
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error converting scanned report item: %s  err: %v \n", scanner.Text(), err)
		}
		expenses[expense]=expense
	}
	found := false
	for keyA, expenseA := range expenses {
		if found {
			break
		}
		for keyB, expenseB := range expenses {
			if keyA == keyB{
				continue
			}
			tempSum := sum - (expenseA + expenseB)
			if _, ok := expenses[tempSum]; ok {
				fmt.Printf("Found Matching Triplet %v + %v + %v = %v \n", expenseA, expenseB, tempSum, sum)
				fmt.Printf("Matching Pair mulitpled %v * %v * %v = %v \n", expenseA, expenseB, tempSum, (expenseA*expenseB)*tempSum)
				found = true
				break
			}
		}
	}
}