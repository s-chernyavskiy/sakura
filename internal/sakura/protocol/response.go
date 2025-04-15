package protocol

import (
	"fmt"
	"strings"
)

type IntegerReply struct {
	Value int
}

func (r *IntegerReply) Reply() string {
	return fmt.Sprintf(":%d\r\n", r.Value)
}

func NewIntegerReply(value int) *IntegerReply {
	return &IntegerReply{Value: value}
}

type BooleanReply struct {
	Value bool
}

func (r *BooleanReply) Reply() string {
	return fmt.Sprintf(":%d\r\n", r.Value)
}

func NewBooleanReply(value bool) *BooleanReply {
	return &BooleanReply{Value: value}
}

type SimpleStringReply struct {
	Value string
}

func (r *SimpleStringReply) Reply() string {
	return fmt.Sprintf("+%s\r\n", r.Value)
}

func NewSimpleStringReply(value string) *SimpleStringReply {
	return &SimpleStringReply{Value: value}
}

type BulkStringReply struct {
	Value string
	Nil   bool
}

func (r *BulkStringReply) Reply() string {
	if r.Nil == true {
		return fmt.Sprintf("$-1\r\n")
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(r.Value), r.Value)
}

func NewBulkStringReply(isNil bool, value string) *BulkStringReply {
	return &BulkStringReply{Nil: isNil, Value: value}
}

type ArrayReply struct {
	Elems []Reply
}

func NewArrayReply(elems []Reply) *ArrayReply {
	return &ArrayReply{Elems: elems}
}

func (r *ArrayReply) Reply() string {
	length := len(r.Elems)
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("*%d\r\n", length))

	for _, re := range r.Elems {
		builder.WriteString(re.Reply() + "\r\n")
	}

	return builder.String()
}

type ErrorReply struct {
	Prefix string
	Err    string
}

func (r *ErrorReply) Error() string {
	return fmt.Sprintf("-%s %s\r\n", r.Prefix, r.Err)
}
