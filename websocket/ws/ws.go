package ws

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	// 协议升级
	Up = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WebSocketConn 封装的ws对象
type WebSocketConn struct {
	conn   *websocket.Conn
	input  chan []byte
	output chan []byte
	close  chan bool
	closed bool
	mutex  sync.Mutex
}

func NewWS(conn *websocket.Conn) (ws *WebSocketConn, err error) {
	ws = &WebSocketConn{conn: conn}
	ws.InitWS()
	return ws, nil
}

func (ws *WebSocketConn) InitWS() {
	ws.input = make(chan []byte, 1000)
	ws.output = make(chan []byte, 1000)
	ws.close = make(chan bool)
	_ = ws.WriteMessage([]byte("connection has been established"))
	go ws.ReadLoop()
	go ws.WriteLoop()
}

func (ws *WebSocketConn) Close() {
	ws.conn.Close()
	// 保证这个chan只关闭一次
	ws.mutex.Lock()
	if !ws.closed {
		close(ws.close)
		ws.closed = true
	}
	ws.mutex.Unlock()
}

func (ws *WebSocketConn) ReadMessage() (r []byte, err error) {
	select {
	case r = <-ws.input:
	case <-ws.close:
		err = errors.New("read failed, connection has been closed")
	}
	return
}

func (ws *WebSocketConn) WriteMessage(msg []byte) (err error) {
	select {
	case ws.output <- msg:
	case <-ws.close:
		err = errors.New("write failed, connection has been closed")
	}
	return
}

// 协程, 不停的读ws内的消息
func (ws *WebSocketConn) ReadLoop() {
	for {
		_, msg, err := ws.conn.ReadMessage()
		if err != nil {
			ws.Close()
			log.Println("read loop exit: ", err)
			return
		}
		// 客户端输入,放到input, 等待读取
		// 会阻塞,而且无法感知close conn,可以用select语法
		//ws.input <- msg
		select {
		case ws.input <- msg:
		case <-ws.close:
			ws.Close()
		}
	}
}

func (ws *WebSocketConn) WriteLoop() {
	for {
		var msg []byte
		select {
		case msg = <-ws.output:
		case <-ws.close:
			ws.Close()
			return
		}
		err := ws.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			ws.Close()
			log.Println("write loop exit: ", err)
		}
	}
}
