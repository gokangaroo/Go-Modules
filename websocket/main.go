package main

import (
	"log"
	"net/http"
	"websocket/ws"
)

func main() {
	// mux router
	http.HandleFunc("/ws", wsHandler)

	// server listen
	http.ListenAndServe(":9000", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// server upgrader
	conn, err := ws.Up.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 自己封装一下,利用chan来线性投递消息
	wsConn, _ := ws.NewWS(conn)
	// 模拟一下竞争
	//go func() {
	//	for {
	//		err = wsConn.WriteMessage([]byte("heartbeat1"))
	//		if err != nil {
	//			wsConn.Close()
	//			return
	//		}
	//	}
	//}()
	//go func() {
	//	for {
	//		err = wsConn.WriteMessage([]byte("heartbeat2"))
	//		if err != nil {
	//			wsConn.Close()
	//			return
	//		}
	//	}
	//}()
	for {
		data, err := wsConn.ReadMessage()
		if err != nil {
			wsConn.Close()
			log.Println(err)
			return
		}
		err = wsConn.WriteMessage(data)
		if err != nil {
			wsConn.Close()
			log.Println(err)
			return
		}
	}

	// 利用websocket处理收发了,但是自身的api并不是线程安全的.
	// 会panic: concurrent write to websocket connection
	//conn.WriteMessage(websocket.TextMessage, []byte("connection has been established"))
	//go func() {
	//	for {
	//		err = conn.WriteMessage(websocket.TextMessage, []byte("heartbeat1"))
	//		if err != nil {
	//			conn.Close()
	//			return
	//		}
	//	}
	//}()
	//go func() {
	//	for {
	//		err = conn.WriteMessage(websocket.TextMessage, []byte("heartbeat2"))
	//		if err != nil {
	//			conn.Close()
	//			return
	//		}
	//	}
	//}()
	//for {
	//	// text ,binary
	//	messageType, data, err := conn.ReadMessage()
	//	if err != nil {
	//		conn.Close()
	//		log.Println("read failed: ", err)
	//		return
	//	}
	//	err = conn.WriteMessage(messageType, data)
	//	if err != nil {
	//		conn.Close()
	//		log.Println("write failed: ", err)
	//	}
	//}
}
