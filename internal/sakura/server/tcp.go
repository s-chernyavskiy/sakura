package server

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/s-chernyavskiy/sakura/internal/sakura/config"
	"github.com/s-chernyavskiy/sakura/internal/sakura/errors"
)

type Clients struct {
	ConnectedClients int
	m                sync.Mutex
}

var ConnectedClients Clients

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
		read, _ := rw.ReadString('\n')
		rw.WriteString(read)
		rw.Flush()
	}
}

func Start(cfg config.AppConfig) {
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
