package protocol

import "fmt"

type IntegerReply struct {
	value int
}

func (rep *IntegerReply) Reply() string {
	return fmt.Sprintf(":%d\r\n", rep.value)
}

type SimpleStringReply struct {
	value string
}

func (rep *SimpleStringReply) Reply() string {
	return fmt.Sprintf("+%s\r\n", rep.value)
}
