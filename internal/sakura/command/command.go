package command

import (
	"github.com/s-chernyavskiy/sakura/internal/sakura/db"
	"github.com/s-chernyavskiy/sakura/internal/sakura/messages"
	"github.com/s-chernyavskiy/sakura/internal/sakura/protocol"
)

type CommandFunc func(*db.DB, []string) protocol.Message

type Command struct {
	ModifyKeySpace bool
	Fn             CommandFunc
}

func getCommand(cmd string) (*Command, protocol.Err) {
	if v, ok := commandTable[cmd]; ok {
		return &v, nil
	}

	return nil, &protocol.ErrUnknownCommand{}
}

var commandTable = map[string]Command{
	"get":    {ModifyKeySpace: false, Fn: messages.Get},
	"set":    {ModifyKeySpace: true, Fn: messages.Set},
	"exists": {ModifyKeySpace: false, Fn: messages.Exists},
}

type DBCommand struct{}

func (DBCommand) Execute(db *db.DB, cmd string, args []string) protocol.Message {
	command, err := getCommand(cmd)
	if err != nil {
		return protocol.Message{Rep: nil, Err: err}
	}

	return command.Fn(db, args)
}
