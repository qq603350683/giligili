package tasks

import (
	"giligili/constbase"
	"giligili/serializer"
	"giligili/socket"
	"giligili/util"
	"log"
	"math/rand"
	"time"
)

type FakeMessage struct {
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func Stop() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(20) + 1
	time.Sleep(time.Second * time.Duration(i))
}

func SendFakeMessage() {
	log.Println("开始假消息推送")
	time.Sleep(time.Second * 2)

	nicknames := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	l := len(nicknames)

	c := make(chan bool, 1)

	c <- true

	for {
		if _, ok := <- c; ok {
			rand.Seed(time.Now().Unix())
			i := rand.Intn(25)

			if i > 20 || len(socket.WebsocketConns.Num) == 0 {
				Stop()

				c <- true

				continue
			}

			i = rand.Intn(l - 1)

			nickname := nicknames[i]

			content := nickname + " is on line"

			fake_message := FakeMessage{
				Content:   content,
				CreatedAt: time.Now().Format(util.DATETIME),
			}

			res := serializer.JsonByte(constbase.CHAT_MESSAGE, "success", fake_message, "")

			for _, conn := range(socket.WebsocketConns.Connects) {

				conn.WriteMessage(constbase.WEBSOCKET_MESSAGE_TYPE_TEXT, res)
			}
			//log.Println(msg)

			Stop()

			c <- true
		}
	}
}
