package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zizimoos/gocoin/utils"
)

// var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {

	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	fmt.Printf("%s wants an upgrade \n", openPort)
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)
	// time.Sleep(20 * time.Second)
	// peer.inbox <- []byte("hello from 3000!!!")
}

func AddPeer(address, port, openPort string) {
	fmt.Printf("%s wants to connect to port %s\n", openPort, port)
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	p := initPeer(conn, address, port)
	sendNewestBlock(p)
	// time.Sleep(10 * time.Second)
	// peer.inbox <- []byte("hello from 4000!!!")
}
