package huffmango

import "slices"

type node struct {
	char        rune
	freq        int
	left, right *node
}

func makeNodes(s string) []node {
	var nodes []node
	m := make(map[rune]int)

	for _, v := range s {
		m[v] += 1
	}

	for char, freq := range m {
		nodes = append(nodes, node{char: char, freq: freq})
	}

	return nodes
}

func (n *node) isLeaf() bool {
	return (n.left == nil) && (n.right == nil)
}

func newInternalNode(node1, node2 node) node {
	return node{
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

func removeAtIndex(nodes []node, index int) []node {
	return slices.Delete(nodes, index, index+1)
}

func makeTreeFromCodeMap(codeMap map[rune]string) *node {
	root := &node{}

	for k, v := range codeMap {
		reference := root
		for i, rune := range v {
			if i == len(v)-1 {
				if rune == '0' {
					reference.left = &node{char: k}
				}

				if rune == '1' {
					reference.right = &node{char: k}
				}
			}

			if rune == '0' {
				if reference.left == nil {
					reference.left = &node{}
				}
				reference = reference.left
				continue
			}

			if rune == '1' {
				if reference.right == nil {
					reference.right = &node{}
				}
				reference = reference.right
				continue
			}
		}
	}

	return root
}
