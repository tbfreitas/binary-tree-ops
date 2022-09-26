package nodes

import (
	"bin-tree/src/structs"
	"fmt"
	"strconv"
)

const (
	error      = "Not a int value, try again..."
	size_array = 7
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

	if 2*index+1 >= len(values) {
		return
	}

	node.LeftNode = CreateNode(values[leftIndex])
	node.RightNode = CreateNode(values[rightIndex])

	MountTree(values, node.LeftNode, leftIndex)   //left node
	MountTree(values, node.RightNode, rightIndex) //right node

	return
}

func PrintTree(node *structs.Node) {

	if node.LeftNode != nil {
		fmt.Println("valor:", node.Value)
		fmt.Println("L:", node.LeftNode.Value)
		fmt.Println("R:", node.RightNode.Value)
		PrintTree(node.LeftNode)
		PrintTree(node.RightNode)

	}
}
