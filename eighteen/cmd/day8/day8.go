package main

import (
	"eighteen/internal/parser"
	"fmt"
	"os"
)

func main() {
	intarr := parser.ParseSpacesdNums("day8.txt")
	fmt.Println("sam", 0+10+11+12+99+2+1+1+2)
	fmt.Println("Day 8 text", intarr)
	//res := partOne(intarr)
	//if res != 138 {
	//	fmt.Println("Part 1 mismatch", 138, res)
	//}
	res := partTwo(intarr)
	fmt.Println("result", res)
}

func part1(intarr []int64) {

}

type NodeHeader struct {
	childNodeCount     int
	metaDataEntryCount int
}
type Node struct {
	ID         string
	header     NodeHeader
	childNodes []Node
	metadata   []int
}

type NodeIndexer struct {
	CurrentIndex int
	NodeValues   []int
}

func valueOfNode(node Node) (value int) {
	if node.header.childNodeCount == 0 {
		for _, m := range node.metadata {
			value += m
		}
	} else {
		children := node.childNodes
		ccound := len(children)
		for _, m := range node.metadata {
			m--
			if m >= ccound {
				continue
			}
			mChild := children[m]
			value += valueOfNode(mChild)
		}
	}
	fmt.Println("id", node.ID, "value", value)
	return
}

func (n *NodeIndexer) GetNext() (returnIndex int) {

	if n.CurrentIndex >= len(n.NodeValues) {
		fmt.Printf("Error")
		os.Exit(0)
	}
	returnIndex = n.NodeValues[n.CurrentIndex]
	n.CurrentIndex++
	return
}
func parseNums(nums []int) (exitStaus int) {

	nodeIndexer := NodeIndexer{CurrentIndex: 0, NodeValues: nums}
	var rootNode Node
	rootNode.header.childNodeCount = nodeIndexer.GetNext()
	rootNode.header.metaDataEntryCount = nodeIndexer.GetNext()

	if rootNode.header.childNodeCount != 2 {
		return -1
	}
	var bChild Node
	bChild.header.childNodeCount = nodeIndexer.GetNext()
	bChild.header.metaDataEntryCount = nodeIndexer.GetNext()
	for i := 0; i < bChild.header.metaDataEntryCount; i++ {
		bChild.metadata = append(bChild.metadata, nodeIndexer.GetNext())
	}
	rootNode.childNodes = append(rootNode.childNodes, bChild)

	var cChild Node
	cChild.header.childNodeCount = nodeIndexer.GetNext()
	cChild.header.metaDataEntryCount = nodeIndexer.GetNext()
	rootNode.childNodes = append(rootNode.childNodes, cChild)
	if len(rootNode.childNodes) != rootNode.header.childNodeCount {
		return -1
	}
	fmt.Println(rootNode)
	return 0
}

func getTopNode(nums []int) (node Node) {
	nodeIndexer := NodeIndexer{CurrentIndex: 0, NodeValues: nums}

	idIndex := 0
	node = doOneNode(&nodeIndexer, &idIndex)
	return
}

func partOne(nums []int) (value int) {
	//nodeIndexer := NodeIndexer{CurrentIndex: 0, NodeValues: nums}

	node := getTopNode(nums)
	fmt.Println("node", node)

	return getMetadataSum(node)
}

func partTwo(nums []int) (value int) {
	//nodeIndexer := NodeIndexer{CurrentIndex: 0, NodeValues: nums}

	node := getTopNode(nums)
	//fmt.Println("node", node)

	return valueOfNode(node)
}

//func parseTwo(nums []int) (exitStatus int) {
//	nodeIndexer := NodeIndexer{CurrentIndex: 0, NodeValues: nums}
//
//	node := doOneNode(&nodeIndexer)
//	fmt.Println("node", node)
//
//	getMetadataSum(node)
//	return
//}

func getMetadataSum(node Node) (sum int) {

	for _, c := range node.childNodes {
		sum += getMetadataSum(c)
	}

	for _, s := range node.metadata {
		sum += s
	}

	fmt.Println("sum", sum)
	return
}

func doOneNode(nodeIndexer *NodeIndexer, idIndex *int) (node Node) {

	node.ID = string('A' + *idIndex)
	*idIndex++

	node.header.childNodeCount = nodeIndexer.GetNext()
	node.header.metaDataEntryCount = nodeIndexer.GetNext()

	for i := 0; i < node.header.childNodeCount; i++ {
		childNode := doOneNode(nodeIndexer, idIndex)
		node.childNodes = append(node.childNodes, childNode)
	}
	for i := 0; i < node.header.metaDataEntryCount; i++ {
		node.metadata = append(node.metadata, nodeIndexer.GetNext())
	}
	fmt.Println("oneNode", node)

	return
}
