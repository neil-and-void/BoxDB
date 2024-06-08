package memtable

import (
	"fmt"
	"math/rand"
)

type Memtable interface {
	Put(k string, v string)
	Get(k string) *Node
}

const POS_INF = "+INF"
const NEG_INF = "-INF"
const HEADS = 1

type SkipList struct {
	head      *Node
	tail      *Node
	height    uint8
	maxHeight uint8
}

func New(maxHeight uint8) *SkipList {
	head := &Node{key: NEG_INF}
	tail := &Node{key: POS_INF}

	head.next = tail
	tail.prev = head

	return &SkipList{head, tail, 1, maxHeight}
}

func (s *SkipList) Put(k string, v string) {
	head := s.head
	tail := s.tail

	newNode := &Node{
		key: k,
		val: v,
	}

	for head.next != tail {
		if head.next != nil && head.next.key < newNode.key {
			head = head.next
		} else if newNode.key < head.next.key && head.down != nil {
			head = head.down
			tail = tail.down
		} else {
			break
		}
	}

	newNode.next = head.next
	head.next.prev = newNode
	head.next = newNode
	newNode.prev = head

	i := uint8(0)
	for i < s.maxHeight && s.coinflip() == HEADS {
		columnNode := &Node{
			key: k,
			val: v,
		}

		newNode.up = columnNode
		columnNode.down = newNode

		// todo: left and right

		columnNode = newNode

		i++
	}
}

func (s *SkipList) Get(k string) *Node {
	return &Node{}
}

func (s *SkipList) coinflip() int {
	return rand.Intn(2)
}

func (s *SkipList) Print() {
	head := s.head

	for head.down != nil {
		head = head.down
	}

	columns := [][]string{}

	for head != nil {

		colNode := head
		column := []string{}

		for colNode != nil {
			column = append(column, colNode.key)
			colNode = colNode.up
		}

		columns = append(columns, column)

		head = head.next
	}

	fmt.Println(columns)

	fmt.Printf("nil\n")
}

type Node struct {
	key  string
	val  string
	up   *Node
	down *Node
	next *Node
	prev *Node
}
