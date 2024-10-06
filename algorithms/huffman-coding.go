package algorithms

import (
	"bytes"
	"container/heap"
	"io"
)

// A node in the huffman tree
type Node struct {
	Char  rune
	Freq  int
	Left  *Node
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

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

type BitWriter struct {
	buffer byte
	count  uint8
	writer io.Writer
}

func NewBitWriter(w io.Writer) *BitWriter {
	return &BitWriter{writer: w}
}

func (bw *BitWriter) WriteBit(bit bool) error {
	if bit {
		bw.buffer |= 1 << (7 - bw.count)
	}
	bw.count++

	if bw.count == 8 {
		if err := bw.Flush(); err != nil {
			return err
		}
	}

	return nil
}

func (bw *BitWriter) Flush() error {
	if bw.count > 0 {
		_, err := bw.writer.Write([]byte{bw.buffer})
		bw.buffer = 0
		bw.count = 0
		return err
	}
	return nil
}

func Huffman_encoding(input []byte) ([]byte, map[byte]string) {
	freqMap := make(map[byte]int)
	for _, char := range input {
		freqMap[char]++
	}

	pq := make(PriorityQueue, 0)
	for char, freq := range freqMap {
		heap.Push(&pq, &Node{Char: rune(char), Freq: freq})
	}

	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		parent := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, parent)
	}

	root := heap.Pop(&pq).(*Node)

	codes := make(map[byte]string)
	generateCodes(root, "", codes)

	var encoded_bytes []byte
	bit_writer := NewBitWriter(bytes.NewBuffer((encoded_bytes)))

	for _, char := range input {
		code := codes[char]
		for _, bit := range code {
			if bit == '1' {
				bit_writer.WriteBit(true)
			} else {
				bit_writer.WriteBit(false)
			}
		}
	}

	bit_writer.Flush()

	return bit_writer.writer.(*bytes.Buffer).Bytes(), codes
}

func generateCodes(node *Node, code string, codes map[byte]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		codes[byte(node.Char)] = code
	}
	generateCodes(node.Left, code+"0", codes)
	generateCodes(node.Right, code+"1", codes)
}
