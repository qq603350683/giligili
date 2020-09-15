package socket

import (
	"fmt"
	"giligili/constbase"
	"giligili/model"
	"giligili/routes"
	"giligili/serializer"
	"giligili/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"runtime"
	"time"
)

type WebsocketPool struct {
	Num chan int
	Connects map[int]*websocket.Conn
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
		Num: make(chan int, 1),
		Connects: make(map[int]*websocket.Conn),
	}

	return pool
}

// 发送消息
func SendMessage(u_id int, message []byte) {
	if conn, ok := WebsocketConns.Connects[u_id]; ok {
		conn.WriteMessage(constbase.WEBSOCKET_MESSAGE_TYPE_TEXT, message)
	}
}

func Run() {
	WebsocketConns = NewWebsocketPool()

	http.HandleFunc("/wechat-login", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		//033nrF1w3r9cZU2yN23w3xTGkN1nrF1d

		if code, ok := query["code"]; ok {
			service.WechantLogin(code[0])
		}
	})

	http.HandleFunc("/plan", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("websocket 错误: %s", err.Error())
			return
		}

		get := r.URL.Query()
		if len(get) == 0 {
			log.Printf("错误链接")
			return
		}

		token, ok := get["token"]
		if !ok {
			log.Printf("Token 不能为空")
			return
		}

		go func(token string, conn *websocket.Conn) {
			var msg []byte

			if token == "" {
				fmt.Println("token不能为空")
				return
			}

			u_id := model.GetUIDByToken(token)
			if u_id == 0 {
				msg = serializer.JsonByte(constbase.WEBSOCKET_CLOSE, "请关闭链接", nil, "")
				conn.WriteMessage(constbase.WEBSOCKET_MESSAGE_TYPE_TEXT, msg)
				return
			}

			log.Printf("登录成功: %d", u_id)

			// 查看是否重复登录， 重复登录就退出
			if c, ok := WebsocketConns.Connects[u_id]; ok {
				log.Printf("账户登陆被强制下线: %d", u_id)
				msg = serializer.JsonByte(constbase.WEBSOCKET_OFFLINE, "您已被强制下线", nil, "")
				c.WriteMessage(constbase.WEBSOCKET_MESSAGE_TYPE_TEXT, msg)

				// 睡眠0.5秒确保消息客户端收到断开消息后断开链接再覆盖 WebsocketConns.Connects
				time.Sleep(time.Second / 2)
				c.Close()
				time.Sleep(time.Second / 2)
			} else {
				WebsocketConns.Num <- 1
			}

			WebsocketConns.Connects[u_id] = conn

			user := model.GetUserInfo(u_id)

			model.UserInfo = user

			for {
				mtype, content, err := conn.ReadMessage()
				if err != nil {
					log.Printf("关闭链接: %d, 错误信息: %s", u_id, err.Error())
					delete(WebsocketConns.Connects, u_id)
					conn.Close()

					if _, ok := <- WebsocketConns.Num; !ok {
						log.Printf("WebsocketConns.Num - 1 fail")
						return
					}

					log.Printf("当前在线人数为: %d", len(WebsocketConns.Num))

					return
				}
				switch mtype {
				case constbase.WEBSOCKET_MESSAGE_TYPE_TEXT:
					//TextMessage
					str := routes.Socket(content)

					conn.WriteMessage(mtype, str)
				case constbase.WEBSOCKET_MESSAGE_TYPE_BINARY:
					// BinaryMessage
				case constbase.WEBSOCKET_MESSAGE_TYPE_CLOSE:
					// CloseMessage
					fmt.Printf("close....")
					conn.Close()
				case constbase.WEBSOCKET_MESSAGE_TYPE_PING:
					// PingMessage
				case constbase.WEBSOCKET_MESSAGE_TYPE_PONG:
					// PongMessage
				}
			}

			// 永远都不会执行到这里
		}(token[0], conn)
	})

	var err error

	// 获取当前操作系统
	sys_type := runtime.GOOS

	log.Printf("%s 启动....", sys_type)

	switch sys_type {
	case "linux":
		err = http.ListenAndServeTLS("0.0.0.0:9501", "4501984_www.yfwethink.com.pem", "4501984_www.yfwethink.com.key", nil)
	case "windows":
		err = http.ListenAndServe("0.0.0.0:9501", nil)
	}

	if (err != nil) {
		panic(err)
	}

}
