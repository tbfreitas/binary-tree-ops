package nodes

import (
	"bin-tree/src/structs"
	"fmt"
	"strconv"
)

const (
	error      = "Not a int value, try again..."
	size_array = 16
)

func GetData() []int {
	var values []int

	for len(values) < size_array {
		temp := ""
		fmt.Scanln(&temp)
		ascii := temp
		asciiInt := string(ascii)

		i, err := strconv.Atoi(asciiInt)
		if err != nil {
			fmt.Println("Not a int value, try again...")
		} else {
			values = append(values, i)
		}
	}

	return values
}

func CreateNode(value int) *structs.Node {

	rootNode := structs.Node{
		RightNode: nil,
		LeftNode:  nil,
		Value:     value,
	}

	return &rootNode
}

func MountTree(values []int, node *structs.Node, index int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	if leftIndex < len(values) {
		node.LeftNode = CreateNode(values[leftIndex])
		MountTree(values, node.LeftNode, leftIndex) //left node
	}

	if rightIndex < len(values) {
		node.RightNode = CreateNode(values[rightIndex])
		MountTree(values, node.RightNode, rightIndex) //right node
	}

	return
}

func GetHeight(node *structs.Node, currentHeight int) int {

	tempLeft := 0
	tempRight := 0

	if node.LeftNode == nil && node.RightNode == nil {
		return currentHeight
	}

	if node.LeftNode != nil {
		tempLeft = GetHeight(node.LeftNode, currentHeight+1)
	} else {
		tempLeft = currentHeight
	}

	if node.RightNode != nil {
		tempRight = GetHeight(node.RightNode, currentHeight+1)
	} else {
		tempRight = currentHeight
	}

	if tempLeft >= tempRight {
		return tempLeft
	}

	return tempRight
}

func PrintTree(node *structs.Node) {
	fmt.Println("valor:", node.Value)

	if node.LeftNode != nil && node.RightNode != nil {
		fmt.Println("L:", node.LeftNode.Value)
		fmt.Println("R:", node.RightNode.Value)
		PrintTree(node.LeftNode)
		PrintTree(node.RightNode)
	}

	if node.LeftNode != nil && node.RightNode == nil {
		fmt.Println("L:", node.LeftNode.Value)
		PrintTree(node.LeftNode)
	}

	if node.RightNode != nil && node.LeftNode == nil {
		fmt.Println("R:", node.RightNode.Value)
		PrintTree(node.RightNode)
	}
}
