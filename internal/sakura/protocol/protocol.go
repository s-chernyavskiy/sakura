package protocol

type Reply interface {
	Reply() string
}

type Message struct {
	Rep Reply
	Error error
}
