package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/s-chernyavskiy/sakura/internal/sakura/command"
	"github.com/s-chernyavskiy/sakura/internal/sakura/config"
	"github.com/s-chernyavskiy/sakura/internal/sakura/db"
	"github.com/s-chernyavskiy/sakura/internal/sakura/errors"
)

type Clients struct {
	ConnectedClients int
	m                sync.Mutex
}

var (
	ConnectedClients Clients
	DB               = db.NewDB()
	dbCommand        = &command.DBCommand{}
)

func (c *Clients) Increase() {
	c.m.Lock()
	defer c.m.Unlock()

	c.ConnectedClients++
}

func (c *Clients) Decrease() {
	c.m.Lock()
	defer c.m.Unlock()

	c.ConnectedClients--
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		read, err := rw.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		strs := strings.Split(read, " ")

		if len(strs) == 0 {
			rw.Flush()
			continue
		}

		log.Printf("got command `%s`\n", read)

		rep := dbCommand.Execute(DB, strings.ToLower(strs[0]), strs[1:])

		if rep.Err == nil {
			rw.WriteString(rep.Rep.Reply())
		} else {
			err := rep.Err.Err()
			rw.WriteString(err.Error())
		}

		rw.Flush()
	}

	log.Println("Disconnected client from", conn.RemoteAddr())
	ConnectedClients.Decrease()
	log.Println(ConnectedClients.ConnectedClients, "connections are now open")
}

func Start(cfg config.AppConfig) {
	fmt.Println(cfg)
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		errors.LogErrorAndExit(err, 3)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			errors.LogError("Error on connection with", conn.RemoteAddr().String(), ":", err.Error())
			conn.Close()
			continue //  NOTE: skip malformed user
		}

		log.Println("Connected client on", conn.RemoteAddr())

		ConnectedClients.Increase()
		log.Println(ConnectedClients.ConnectedClients, "connections are now open")

		go handleConnection(conn)
	}
}
