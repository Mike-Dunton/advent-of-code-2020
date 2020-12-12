package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	Operation string
	Argument int
}

func main() {
	//Read Input File
	instructionSetRaw, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer instructionSetRaw.Close()
	partOne(instructionSetRaw)
	//partTwo()
}
func partOne(instructionSetRaw *os.File) {

	scanner := bufio.NewScanner(instructionSetRaw)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		re := regexp.MustCompile(`^(?P<operation>[a-z]{3}) (?P<argument>[-|\+]\d+)$`)
		if re.MatchString(scanner.Text()) {
			matches := re.FindStringSubmatch(scanner.Text())
			argumentInt, _ := strconv.Atoi(matches[re.SubexpIndex("argument")], )
			instructions = append(instructions, Instruction{
				matches[re.SubexpIndex("operation")],
				argumentInt,
			})
		}
	}
	for i, instruction := range instructions {
		fmt.Printf("index: %v instruction: %v %v\n", i, instruction.Operation ,instruction.Argument)
	}
	isLoop, loopedInstructionIndex, currI, acc, lastTen := findLoopAndAcc(instructions)
	if isLoop {
		fmt.Printf("Acc detected at loop lastI: %v currI: %v acc: %v lastTen %v\n",loopedInstructionIndex, currI, acc, lastTen)
		for _, changeIndex := range lastTen {
			tempInstructions := instructions
			fixedInstruction := swapInstruction(tempInstructions[changeIndex])
			tempInstructions[changeIndex] = fixedInstruction
			isLoop, loopedInstructionIndex, currI, acc, lastTen = findLoopAndAcc(instructions)
			if !isLoop {
				fmt.Printf("Broken Instruction was %v \n", changeIndex)
				break
			}
		}
	}

	fmt.Printf("Loop detected: %v acc: %v \n",isLoop,acc)

}

func findLoopAndAcc(instructions []Instruction) (bool, int, int, int, []int) {
	lastTen := make([]int, 0)
	visited := make([]bool, len(instructions))
	acc := 0
	lastI := 0
	currI := lastI
	for currI < len(instructions) {
		if visited[currI] == true {
			return true, lastI, currI, acc, lastTen
		}
		switch instructions[currI].Operation {
		case "nop":
			visited[currI] = true
			lastTen = append([]int{currI}, lastTen...)
			lastI = currI
			currI++
			continue
		case "acc":
			acc += instructions[currI].Argument
			visited[currI] = true
			lastTen = append([]int{currI}, lastTen...)
			lastI = currI
			currI++
			continue
		case "jmp":
			visited[currI] = true
			lastTen = append([]int{currI}, lastTen...)
			lastI = currI
			currI += instructions[currI].Argument
			continue
		default:
			fmt.Println("Unknown instruction")
			os.Exit(1)
		}
	}
	return false, lastI, currI, acc, lastTen
}

func swapInstruction(instruction Instruction) Instruction {
	switch instruction.Operation {
	case "nop":
		//fmt.Println("Switched instruction to jmp")
		instruction.Operation = "jmp"
		return instruction
	case "jmp":
		//fmt.Println("Switched instruction to nop")
		instruction.Operation = "nop"
		return instruction
	}
	return instruction
}