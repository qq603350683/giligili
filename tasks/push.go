package tasks

import (
	"giligili/socket"
	"log"
	"math/rand"
	"time"
)

func Stop() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(3) + 1
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
			i := rand.Intn(100)

			log.Println(i)

			if i > 20 || len(socket.WebsocketConns.Num) == 0 {
				Stop()

				c <- true

				continue
			}

			i = rand.Intn(l - 1)

			nickname := nicknames[i]

			msg := nickname + " is on line"

			log.Println(msg)

			Stop()

			c <- true
		}
	}
}
