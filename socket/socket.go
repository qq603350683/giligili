package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketPool struct {
	num chan int
	connects map[string]*websocket.Conn
}

var upgrade = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var WebsocketConns *WebsocketPool

func NewWebsocketPool() *WebsocketPool {
	pool := &WebsocketPool{
		num: make(chan int, 1),
		connects: make(map[string]*websocket.Conn),
	}

	return pool
}

func main() {
	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world"))
	//})

	WebsocketConns = NewWebsocketPool()

	http.HandleFunc("/qqq", func (w http.ResponseWriter, r *http.Request) {
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		get := r.URL.Query()
		if len(get) == 0 {
			panic("登录对象为空")
		}

		token, ok := get["token"]
		if !ok {
			panic("token 不存在")
		}

		go func(token string, conn *websocket.Conn) {
			if token == "" {
				fmt.Println("token不能为空")
				return
			}

			WebsocketConns.num <- 1
			WebsocketConns.connects[token] = conn

			for {
				fmt.Println(WebsocketConns.num)
				mtype, content, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("客户端主动关闭链接")
					conn.Close()

					if _, ok := <- WebsocketConns.num; !ok {
						fmt.Println("WebsocketConns.num - 1 fail")
						return
					}

					fmt.Println("玩家数剩余")
					fmt.Println(len(WebsocketConns.num))

					return
				}
				switch mtype {
				case 1:
					//TextMessage
					conn.WriteMessage(mtype, content)
				case 2:
					// BinaryMessage
				case 8:
					// CloseMessage
					conn.Close()
				case 9:
					// PingMessage
				case 10:
					// PongMessage
				}
			}

			// 永远都不会执行到这里
		}(token[0], conn)
	})

	fmt.Println("启动....")

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if (err != nil) {
		panic(err)
	}

	fmt.Println("zzz")
}
