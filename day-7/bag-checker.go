package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Contains []BagRule
	Color string
}

type BagRule struct {
	Count int
	Color string
}

func main() {
	//Read Input File
	bagRulesFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer bagRulesFile.Close()
	partOne(bagRulesFile)
	//partTwo()
}

func partOne(bagRulesFile *os.File) {
	//Read Input File
	scanner := bufio.NewScanner(bagRulesFile)
	bags := make(map[string]Bag)
	isContainedBy := make(map[string]map[string]struct{})
	for scanner.Scan() {
		bagRules := make([]BagRule, 0)
		bagRuleSplit := strings.Split(scanner.Text(), " bags contain ")
		for _, currRule := range strings.Split(bagRuleSplit[1], ",") {
			re := regexp.MustCompile(`^(?P<count>\d+) (?P<color>[\w\s]+) bag`)
			if re.MatchString(strings.Trim(currRule, " ")) {
				matches := re.FindStringSubmatch(strings.Trim(currRule, " "))
				countInt, _:= strconv.Atoi(matches[re.SubexpIndex("count")])
				bagRules = append(bagRules, BagRule{countInt, matches[re.SubexpIndex("color")]})
				if isContainedBy[matches[re.SubexpIndex("color")]] == nil {
					isContainedBy[matches[re.SubexpIndex("color")]] = make(map[string]struct{})
				}
				isContainedBy[matches[re.SubexpIndex("color")]][bagRuleSplit[0]] = struct{}{}
			} else {
				fmt.Printf("Rule did not match %v\n", strings.Trim(currRule, " "))
			}
		}
		bags[bagRuleSplit[0]] = Bag{bagRules, bagRuleSplit[0]}
	}
	
	fmt.Println(len(findParentBags("shiny gold", isContainedBy)))
	fmt.Println(findChildBags("shiny gold", bags))

}

func findParentBags(color string, Bags map[string]map[string]struct{}) map[string]struct{} {
	containerBags := Bags[color]
	for parentColor, _ := range Bags[color] {
		for parentsContainers, _ := range findParentBags(parentColor, Bags) {
			containerBags[parentsContainers] = struct{}{}
		}
	}
	//fmt.Printf("Building ContainerBags %v\n",containerBags )
	return containerBags
}

func findChildBags(color string, bags map[string]Bag) int{
	runningTotal := 0
	for _, r := range bags[color].Contains {
		runningTotal += r.Count + r.Count*findChildBags(r.Color, bags)
	}
	return runningTotal
}