package main

import (
	"bin-tree/src/nodes"
)

func main() {

	x := nodes.GetData()

	println("--Stopping getting data--")

	rootNode := nodes.CreateNode(x[0])
	nodes.MountTree(x, rootNode, 0)
	nodes.PrintTree(rootNode)
	// nodes.getHeight()
	// nodes.checkIfIsBallenced
	// nodes.ballanceTree()
}
