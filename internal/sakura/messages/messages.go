package messages

import (
	"github.com/s-chernyavskiy/sakura/internal/sakura/db"
	"github.com/s-chernyavskiy/sakura/internal/sakura/protocol"
	"github.com/s-chernyavskiy/sakura/pkg/util"
)

func Get(d *db.DB, args []string) protocol.Message {
	if len(args) != 1 {
		return protocol.Message{Rep: nil, Err: &protocol.ErrInsufficientArgs{Cmd: "get"}}
	}

	val, err := d.Get(args[0])
	if err != nil {
		return protocol.Message{Rep: nil, Err: &protocol.ErrGeneric{Error: err}}
	}

	if val.Type != db.TypeString {
		return protocol.Message{Rep: nil, Err: &protocol.ErrWrongType{}}
	}

	return protocol.Message{Rep: protocol.NewBulkStringReply(false, util.ToString(val.Value)), Err: nil}
}

func Set(d *db.DB, args []string) protocol.Message {
	if len(args) != 2 {
		return protocol.Message{Rep: nil, Err: &protocol.ErrInsufficientArgs{Cmd: "set"}}
	}

	key := args[0]
	val := args[1]

	d.Set(key, db.NewDataNode(db.TypeString, -1, val))

	return protocol.Message{Rep: protocol.NewSimpleStringReply("OK"), Err: nil}
}

func Exists(d *db.DB, args []string) protocol.Message {
	if len(args) != 1 {
		return protocol.Message{Rep: nil, Err: &protocol.ErrInsufficientArgs{Cmd: "get"}}
	}
	found := d.Exists(args[0])

	return protocol.Message{Rep: protocol.NewBooleanReply(found), Err: nil}
}
