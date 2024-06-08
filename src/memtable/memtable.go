package memtable

import (
	"fmt"
	"math/rand"
)

type Memtable interface {
	Put(k string, v string)
	Get(k string) *Node
}

const NEG_INF = "-INF"
const POS_INF = "+INF"
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
	curNode := s.head

	newNode := &Node{
		key: k,
		val: v,
	}

	for curNode.next != tail {
		if curNode.next != nil && curNode.next.key < newNode.key {
			curNode = curNode.next
		} else if newNode.key < curNode.next.key && curNode.down != nil {
			curNode = curNode.down
			tail = tail.down
			head = head.down
		} else {
			// found it
			break
		}
	}

	newNode.next = curNode.next
	curNode.next.prev = newNode
	curNode.next = newNode
	newNode.prev = curNode

	nearestLeftNode := newNode.prev
	nearestRightNode := newNode.next
	curNode = newNode

	var i uint8 = 0
	for s.coinflip() == HEADS {
		if i >= s.height-1 {
			colHead := &Node{key: NEG_INF}
			colTail := &Node{key: POS_INF}

			head.up = colHead
			tail.up = colTail

			s.height++
		}

		for nearestLeftNode.up == nil {
			nearestLeftNode = nearestLeftNode.prev
		}
		nearestLeftNode = nearestLeftNode.up

		for nearestRightNode.up == nil {
			nearestRightNode = nearestRightNode.next
		}
		nearestRightNode = nearestRightNode.up

		newColumnNode := &Node{
			key: k,
			val: v,
		}

		curNode.up = newColumnNode
		newColumnNode.down = curNode

		newColumnNode.prev = nearestLeftNode
		newColumnNode.next = nearestRightNode
		nearestLeftNode.next = newColumnNode
		nearestRightNode.prev = newColumnNode

		curNode = newColumnNode
		head = head.up
		tail = tail.up

		i++
	}
}

func (s *SkipList) Get(k string) *Node {
	// todo
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

	for head != nil {

		colNode := head
		column := []string{}

		for colNode != nil {
			column = append(column, colNode.key)
			colNode = colNode.up
		}

		fmt.Println(column)
		head = head.next
	}
}

type Node struct {
	key  string
	val  string
	up   *Node
	down *Node
	next *Node
	prev *Node
}
