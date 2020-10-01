package main

import "fmt"

type node struct {
	val  int
	next *node
}

func f(n *node) *node {
	p := n
	var newHead *node
	for p != nil {
		if p.next == nil {
			p.next = newHead
			newHead = p
			break
		}
		tmp := p.next
		p.next = newHead
		newHead = p
		p = tmp
	}
	return newHead
}

func main() {
	head := &node{1, nil}
	tail := head
	var p *node
	for i := 2; i < 5; i++ {
		p = &node{i, nil}
		tail.next = p
		tail = tail.next
	}

	_h := f(head)

	for _h != nil {
		fmt.Println(_h.val)
		_h = _h.next
	}
}
