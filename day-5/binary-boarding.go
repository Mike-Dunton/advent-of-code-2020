package main

import (
	"bufio"
	"fmt"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  BoardingPass
}

type BinaryTree struct {
	root *BinaryNode
}

type BoardingPass struct {
	row int
	col int
	ID int
}

func main() {
	//Read Input File
	boardingFile, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file. err: %v \n", err)
	}
	defer boardingFile.Close()
	partOne(boardingFile)
	//partTwo()
}

func partOne(passwordsFile *os.File) {
	tree := &BinaryTree{}
	scanner := bufio.NewScanner(passwordsFile)
	for scanner.Scan() {
		if isValidBoardingPass(scanner.Text()) {
			var seatRow, seatCol int
			currMinRow := 0
			currMaxRow := 127
			currMinCol := 0
			currMaxCol := 7
			completeBoardingPass := scanner.Text()
			rowInfo := completeBoardingPass[:len(completeBoardingPass)-3]
			seatInfo := completeBoardingPass[7:]
			//fmt.Printf("row info: %s len: %s \n", rowInfo, len(rowInfo))
			//fmt.Printf("seatInfo: %s len: %s \n", seatInfo, len(seatInfo))


			for i, currSeatEncoding := range  rowInfo {
				if string(currSeatEncoding) == "F" {
					if i+1 == len(rowInfo) {
						seatRow = currMinRow
						break
					}
					currMinRow, currMaxRow = lowerRange(currMinRow, currMaxRow)
				} else if string(currSeatEncoding) == "B" {
					if i+1 == len(rowInfo) {
						seatRow = currMaxRow
						break
					}
					currMinRow, currMaxRow = upperRange(currMinRow, currMaxRow)
				}
			}


			for i, currSeatEncoding := range  seatInfo {
				if string(currSeatEncoding) == "R" {
					if i+1 == len(seatInfo) {
						seatCol = currMaxCol
						break
					}
					currMinCol, currMaxCol = upperRange(currMinCol, currMaxCol)
				} else if string(currSeatEncoding) == "L" {
					if i+1 == len(seatInfo) {
						seatCol = currMinCol
						break
					}
					currMinCol, currMaxCol = lowerRange(currMinCol, currMaxCol)
				}
			}
			//fmt.Printf("seatRow: %v, seatCol: %v , (seatRow*8)+seatCol: %v\n", seatRow, seatCol, (seatRow*8)+seatCol)
			tree.insert(BoardingPass{seatRow, seatCol, (seatRow*8)+seatCol})
		}
	}
	fmt.Printf("Max Boarding Pass: %v", tree.Max())
	for i := tree.Min().ID; i <= tree.Max().ID; i++ {
		if !tree.Search(i) {
			fmt.Printf("Your Seat ID : %v", i)
		}
	}
}

func lowerRange(min int , max int) (int, int) {
	return min, (min+max-1)/2
}

func upperRange(min int , max int) (int, int) {
	return (min+max+1)/2, max
}

func isValidBoardingPass(boardingPass string) bool {
	return len(boardingPass) == 10
}

func (t *BinaryTree) insert(data BoardingPass) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data BoardingPass) {
	if n == nil {
		return
	} else if data.ID <= n.data.ID {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

// Max returns the BoardingPass with max ID value stored in the tree
func (t *BinaryTree) Max() *BoardingPass {
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.data
		}
		n = n.right
	}
}

// Min returns the Item with min value stored in the tree
func (t *BinaryTree) Min() *BoardingPass {
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.data
		}
		n = n.left
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (t *BinaryTree) InOrderTraverse(f func(BoardingPass)) {
	inOrderTraverse(t.root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *BinaryNode, f func(BoardingPass)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.data)
		inOrderTraverse(n.right, f)
	}
}

// Search returns true if the BoardingPass ID exists in the tree
func (t *BinaryTree) Search(ID int) bool {
	return search(t.root, ID)
}

// internal recursive function to search an item in the tree
func search(n *BinaryNode, ID int) bool {
	if n == nil {
		return false
	}
	if ID < n.data.ID {
		return search(n.left, ID)
	}
	if ID > n.data.ID {
		return search(n.right, ID)
	}
	return true
}