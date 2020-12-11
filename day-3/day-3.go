package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	X int
	Y int
}

func main() {
	//Read Input File
	mapFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer mapFile.Close()
	scanner := bufio.NewScanner(mapFile)
	mapArray := []string{}
	for scanner.Scan() {
		mapArray = append(mapArray, scanner.Text())
	}

	slopes := []Slope{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	runningProduct := 1
	for _, currentSlope := range slopes {
		tmpTreeCount := checkSlopeTreeCount(mapArray, currentSlope.X, currentSlope.Y)
		fmt.Printf("currSlopeX: %v, currSlopeY: %v , TreeCount: %v \n", currentSlope.X, currentSlope.Y, tmpTreeCount)
		runningProduct *= tmpTreeCount
	}

	fmt.Printf("Product Of All Trees From all Slopes: %v", runningProduct)
}

func safeWrap(index int, length int) int {
	return (index + length) % length
}

func isTree(encountedTile string) bool {
	tree := "#"
	return encountedTile == tree
}

func checkSlopeTreeCount(mapArray []string, slopeX int, slopeY int) int {
	treeCount := 0
	for x, y := slopeX, slopeY; x < len(mapArray); x+=slopeX {
		if isTree(string(mapArray[x][safeWrap(y, len(mapArray[x]))])) {
			treeCount++
		}
		y+=slopeY
	}
	return treeCount
}