package types

import (
	"errors"
	"fmt"
	"sync"
)

type Collection interface {
	Add(any) error
	Delete(any) error
	Update(any) error
	Find(any) (any, error)
	Count() uint8
	Clear() error
	Sort(func(a any, b any))
}

type List struct {
	Capacity uint8
	data     []string
	m        sync.Mutex
}

func (l *List) Add(value any) error {
	var s string
	switch v := value.(type) {
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return errors.New("bad type to add")
	}

	l.m.Lock()
	defer l.m.Unlock()
	l.data = append(l.data, s)

	return nil
}

type HashMap struct {
	Capacity uint8
	data     map[string]string
}

type Integer int

type UInt8 uint8

type String string
