package memtable

import (
	"fmt"
	"math/rand"
)

type Memtable interface {
	Put(k string, v string)
	Get(k string) *string
}

const NEG_INF = "__-INF__"
const POS_INF = "__+INF__"
const HEADS = 1

type node struct {
	key  string
	val  string
	up   *node
	down *node
	next *node
	prev *node
}

type SkipList struct {
	head      *node
	tail      *node
	height    uint8
	maxHeight uint8
	comparer  Comparer
}

func NewMemTable(comparer Comparer, maxHeight uint8) *SkipList {
	head := &node{key: NEG_INF}
	tail := &node{key: POS_INF}

	head.next = tail
	tail.prev = head

	return &SkipList{head, tail, 1, maxHeight, comparer}
}

func (s *SkipList) Put(k string, v string) {
	if k == POS_INF || k == NEG_INF {
		panic("Cannot use reserved key")
	}

	head := s.head
	tail := s.tail
	curNode := s.head

	for curNode.next != nil {
		if curNode.key == k {
			for curNode != nil {
				curNode.val = v
				curNode = curNode.down
			}
			return
		} else if s.comparer.isLessThanOrEqual(curNode.next.key, k) {
			curNode = curNode.next
		} else if s.comparer.isLessThan(k, curNode.next.key) && curNode.down != nil {
			curNode = curNode.down
			tail = tail.down
			head = head.down
		} else {
			break
		}
	}

	newNode := &node{
		key: k,
		val: v,
	}

	newNode.next = curNode.next
	newNode.prev = curNode

	curNode.next.prev = newNode
	curNode.next = newNode

	nearestLeftNode := newNode.prev
	nearestRightNode := newNode.next
	curNode = newNode

	var levelIndex uint8 = 0
	for s.coinflip() == HEADS && s.height < s.maxHeight {
		if levelIndex >= s.height-1 {
			newHead := &node{key: NEG_INF}
			newTail := &node{key: POS_INF}

			head.up = newHead
			tail.up = newTail

			newHead.down = head
			newTail.down = tail

			s.head = newHead
			s.tail = newTail

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

		newColumnNode := &node{
			key: k,
			val: v,
		}

		curNode.up = newColumnNode
		newColumnNode.down = curNode

		newColumnNode.next = nearestRightNode
		nearestRightNode.prev = newColumnNode

		newColumnNode.prev = nearestLeftNode
		nearestLeftNode.next = newColumnNode

		curNode = newColumnNode
		head = head.up
		tail = tail.up

		levelIndex++
	}
}

func (s *SkipList) Get(k string) *string {
	curNode := s.head
	tail := s.tail

	for curNode != tail {
		if k == curNode.key {
			return &curNode.val
		} else if curNode.next != tail && curNode.next.key <= k {
			curNode = curNode.next
		} else if curNode.down != nil {
			curNode = curNode.down
			tail = tail.down
		} else {
			break
		}
	}

	return nil
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

		curNode := head

		for curNode != nil {
			fmt.Printf("(%s | %s) ", curNode.key, curNode.val)
			curNode = curNode.up
		}

		fmt.Printf("\n")

		head = head.next
	}
}
