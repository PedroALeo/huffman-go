package main

import (
	"fmt"
	"strconv"
)

type node struct {
	char        string
	freq        int
	left, right *node
}

func makeNodes(s string) []node {
	var nodes []node
	m := make(map[string]int)

	for _, v := range s {
		m[string(v)] += 1
	}

	fmt.Printf("%v\n", m)

	for char, freq := range m {
		nodes = append(nodes, node{char: char, freq: freq})
	}

	return nodes
}

func (n *node) isLeaf() bool {
	return (n.left == nil) && (n.right == nil)
}

func printCode(root node, steps []int) {
	if root.isLeaf() {
		var code string
		for _, s := range steps {
			code += strconv.Itoa(s)
		}
		fmt.Printf("%v %v\n", root.char, code)
		return
	}

	if root.left != nil {
		steps = append(steps, 0)
		printCode(*root.left, steps)
	}

	if root.right != nil {
		steps = append(steps[:len(steps)-1], 1)
		printCode(*root.right, steps)
	}
}

func removeAtIndex(nodes []node, index int) []node {
	return append(nodes[:index], nodes[index+1:]...)
}

func newInternalNode(node1, node2 node) node {
	return node{
		char:  "Internal node",
		freq:  node1.freq + node2.freq,
		left:  &node1,
		right: &node2,
	}
}

func findLowestFreq(nodes []node) int {
	lf := nodes[0].freq
	ci := 0

	for i, node := range nodes {
		if node.freq < lf {
			lf = node.freq
			ci = i
		}
	}

	return ci
}

func iteration(nodes []node) []node {
	lfi1 := findLowestFreq(nodes)
	lfn1 := nodes[lfi1]

	nodes = removeAtIndex(nodes, lfi1)

	lfi2 := findLowestFreq(nodes)
	lfn2 := nodes[lfi2]

	nodes = removeAtIndex(nodes, lfi2)

	iNode := newInternalNode(lfn1, lfn2)

	nodes = append(nodes, iNode)

	return nodes
}

func CreateHuffmanCodeFromString(s string) {
	nodes := makeNodes(s)

	for len(nodes) > 1 {
		nodes = iteration(nodes)
	}

	printCode(nodes[0], []int{})
}

func main() {
	CreateHuffmanCodeFromString("fljkashfkjlshagdfjkhsdkjfghsdaghfksdghfkhsdfjkhsadjkfhskdjfhjkasdhfjklasdfghjlaskdfgaskldjhfgasdjkhlfgasdjklhfgsadjklhf")
}
