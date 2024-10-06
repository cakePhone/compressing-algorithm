package huffman

import {
	"container/heap"
	"sort"
}

// A node in the huffman tree
type Node struct {
	Char rune
	Freq int
	Left *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interfce{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func huffman_encoding(text string) (string, map[rune]string) {
	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	pq := make(PriorityQueue, 0)
	for char, freq := range freqMap {
		heap.Push(&pq, &Node{Char: char, Freq: freq})
	}

	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		parent := &Node{
			Freq: left.Freq + right.Freq,
			Left: left,
			Right: right
		}
		heap.Push(&pq, parent)
	}

	root := heap.Pop(&pq).(*Node)

	codes := make(map[rune]string)
	generateCodes(root, "", codes)

	var encoded string
	for _, char := range text {
		encoded += codes[char]
	}

	return encoded, codes
}

func generateCodes(node *Node, code string, codes map[rune]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		codes[node.Char] = code
	}
	generateCodes(node.Left, code+"0", codes)
	generateCodes(node.Right, code+"1", codes)
}
