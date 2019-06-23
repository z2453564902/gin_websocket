package middle

import (
	"encoding/json"
	"net/http"

	ws "service/app/common/websocket"
	"service/app/model"
	"service/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//WsHandle .
func WsHandle(c *gin.Context) {
	// change the reqest to websocket model
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//判断用户是否登录, 不登录不能连socket
	var (
		cookie string
		redis  []byte
		users  model.Users
	)
	cookie, _ = c.Cookie(config.TokenName)
	redis, _ = model.Redis.Get(cookie).Bytes()
	json.Unmarshal(redis, &users)
	if users.ID < 1 {
		jsonMessage, _ := json.Marshal(&ws.Message{Content: "用户信息失效,请重新登录"})
		conn.WriteMessage(ws.NotLogin, jsonMessage)
		conn.Close()
	}
	// websocket connect
	client := &ws.Client{
		ID:     users.ID,
		Socket: conn,
		Send:   make(chan []byte),
	}
	ws.Manager.Register <- client
	go client.Read()
	go client.Write()
}
