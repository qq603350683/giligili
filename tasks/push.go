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

	tips := [...]string{
		"首次签到即可免费领取500钻石",
		"商店仅需500钻石购买2020飞行号",
		"在商店购买强化器可以强化子弹哦",
		"背包点击强化器即可强化子弹与技能",
		"子弹越强病毒消灭越快哦",
		"10月前7天签到每天领取500钻石",
		"商店购买新式战斗机可提高战斗力",
		"点击首页立即领取有惊喜！",
	}
	l := len(tips)

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

			tip := tips[i]

			//content := nickname + " is on line"

			fake_message := FakeMessage{
				Content:   tip,
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
