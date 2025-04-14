package types

import (
	"errors"
	"sync"
)

type Node struct {
	next, prev *Node
	list       *List
	Value      string
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

func (l *List) insertValue(v string, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

func (l *List) pushBack(v string) *Node {
	l.allocateSpace()
	return l.insertValue(v, l.root.prev)
}

func (l *List) pushFront(v string) *Node {
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

func (l *List) remove(n *Node) string {
	if n.list == l {
		l.removeNode(n)
	}

	return n.Value
}

func (l *List) pushBackList(other *List) {
	l.allocateSpace()
	for i, e := other.Length(), other.front(); i > 0; i, e = i-1, e.nextNode() {
		l.insertValue(e.Value, l.root.prev)
	}
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

func (l *List) Length() int {
	l.m.Lock()
	defer l.m.Unlock()
	return l.len
}

func (l *List) PushFront(val ...string) error {
	l.m.Lock()

	defer l.m.Unlock()

	if len(val) == 0 {
		return errors.New("No items to insert")
	} else if len(val) == 1 {
		l.pushFront(val[0])
	} else {
		newList := buildValueList(true, val...)
		l.pushFrontList(newList)
	}

	return nil
}

func (l *List) PushBack(val ...string) error {
	l.m.Lock()

	defer l.m.Unlock()

	if len(val) == 0 {
		return errors.New("No items to insert")
	} else if len(val) == 1 {
		l.pushBack(val[0])
	} else {
		newList := buildValueList(false, val...)
		l.pushBackList(newList)
	}

	return nil
}

func (l *List) Head() *Node {
	return l.front()
}

func (l *List) Tail() *Node {
	return l.back()
}

func (l *List) PopBack() string {
	l.m.Lock()
	defer l.m.Unlock()
	if l.Tail() == nil {
		return ""
	}

	return l.remove(l.Tail())
}

func (l *List) PopFront() string {
	l.m.Lock()
	defer l.m.Unlock()
	if l.Head() == nil {
		return ""
	}

	return l.remove(l.Head())
}
