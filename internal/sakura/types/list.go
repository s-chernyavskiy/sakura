package types

import (
	"errors"
	"sync"
)

type Node struct {
	next, prev *Node
	list       *List
	Value      any
}

func (n *Node) nextNode() *Node {
	if e := n.next; n.list != nil && n != &n.list.root {
		return e
	}

	return nil
}

func (n *Node) prevNode() *Node {
	if e := n.prev; n.list != nil && n != &n.list.root {
		return e
	}

	return nil
}

type List struct {
	root Node
	len  int
	m    sync.Mutex
}

func (l *List) init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0

	return l
}

func newList() *List {
	return new(List).init()
}

func (l *List) Length() int {
	return l.len
}

func (l *List) front() *Node {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *List) back() *Node {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

func (l *List) allocateSpace() {
	if l.root.next == nil {
		l.init()
	}
}

func (l *List) insert(n, at *Node) *Node {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++

	return n
}

func (l *List) insertValue(v any, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

func (l *List) pushBack(v any) *Node {
	l.allocateSpace()
	return l.insertValue(v, l.root.prev)
}

func (l *List) pushFront(v any) *Node {
	l.allocateSpace()
	return l.insertValue(v, &l.root)
}

func (l *List) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.len--
}

func (l *List) remove(n *Node) any {
	if n.list == l {
		l.removeNode(n)
	}

	return n.Value
}

func (l *List) pushFrontList(other *List) {
	l.allocateSpace()
	for i, e := other.Length(), other.back(); i > 0; i, e = i-1, e.prevNode() {
		l.insertValue(e.Value, &l.root)
	}
}

func buildValueList(front bool, val ...string) *List {
	l := newList()

	if front == true {
		for _, v := range val {
			l.pushFront(v)
		}

		return l
	}
	for _, v := range val {
		l.pushBack(v)
	}

	return l
}

func (list *List) PushFront(val ...string) error {
	list.m.Lock()

	defer list.m.Unlock()

	if len(val) == 0 {
		return errors.New("No items to insert")
	} else if len(val) == 1 {
		list.pushFront(val)
	} else {
		newList := buildValueList(true, val...)
		list.pushFrontList(newList)
	}

	return nil
}


