package main

import (
	"bin-tree/src/nodes"
)

func main() {

	x := nodes.GetData()

	println("-----------------------")
	println(" *Stopping getting data --")
	println("-----------------------")

	rootNode := nodes.CreateNode(x[0])
	nodes.MountTree(x, rootNode, 0)
	nodes.PrintTree(rootNode)
	height := nodes.GetHeight(rootNode, 1)

	println("-----------------------")
	println(" *Tree height: ", height)
	println("-----------------------")

	// nodes.checkIfIsBallenced
	// nodes.ballanceTree()
}
