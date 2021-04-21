package server

import (
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/modules/chat/dao"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/modules/open/middleware"
	"zodream.cn/godream/utils/upload"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 256 * 2
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// 用户id
	userId uint

	// 房间类型 0 为单用户 1 为群
	roomType uint32

	// 房间的id 用户id 或 群id
	roomId uint
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, content, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		event, _, message := decodeMessage(string(content))
		fmt.Println("[chat event]" + event)
		hanleFuncMap := map[string]func(string){
			EVENT_AUTH: func(s string) {
				c.userId = middleware.JWTTokenUser(s)
				if c.userId < 1 {
					c.hub.unregister <- c
					c.conn.Close()
					return
				}
			},
			EVENT_MESSAGE_SEND_TEXT: func(s string) {
				var form models.MssageForm
				decodeJSON(s, &form)
				messageModel, err := dao.SendText(c.userId, form.ItemType, form.ItemId, form.Content)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				c.hub.broadcast <- messageModel
			},
			EVENT_MESSAGE_SEND_IMAGE: func(s string) {
				var form models.MssageForm
				err := decodeJSON(s, &form)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				file, url := configs.UploadRandomFileName(form.FileName)
				if err := upload.SaveBase64(form.File, file); err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				messageModel, err := dao.SendImage(c.userId, form.ItemType, form.ItemId, url)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				c.hub.broadcast <- messageModel
			},
			EVENT_MESSAGE_SEND_VIDEO: func(s string) {
				var form models.MssageForm
				decodeJSON(s, &form)
				file, url := configs.UploadRandomFileName(form.FileName)
				if err := upload.SaveBase64(form.File, file); err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				messageModel, err := dao.SendVideo(c.userId, form.ItemType, form.ItemId, url)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				c.hub.broadcast <- messageModel
			},
			EVENT_MESSAGE_SEND_AUDIO: func(s string) {
				var form models.MssageForm
				decodeJSON(s, &form)
				file, url := configs.UploadRandomFileName(form.FileName)
				if err := upload.SaveBase64(form.File, file); err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				messageModel, err := dao.SendVoice(c.userId, form.ItemType, form.ItemId, url)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				c.hub.broadcast <- messageModel
			},
			EVENT_MESSAGE_SEND_FILE: func(s string) {
				var form models.MssageForm
				decodeJSON(s, &form)
				file, url := configs.UploadRandomFileName(form.FileName)
				if err := upload.SaveBase64(form.File, file); err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				messageModel, err := dao.SendFile(c.userId, form.ItemType, form.ItemId, form.FileName, url)
				if err != nil {
					c.send <- encodeByte(EVENT_ERROR, MESSAGE_TYPE_STRING, []byte(err.Error()))
					return
				}
				c.hub.broadcast <- messageModel
			},
			EVENT_PROFILE: func(s string) {
				profile := dao.GetProfile(c.userId)
				c.send <- encodeJSON(EVENT_PROFILE, profile)
			},
			EVENT_HISTORY: func(s string) {
				data := dao.GetHistories(c.userId)
				c.send <- encodeJSON(EVENT_HISTORY, gin.H{
					"data": data,
				})
			},
			EVENT_FRIENDS: func(s string) {
				data := dao.GetFriendList(c.userId)
				c.send <- encodeJSON(EVENT_FRIENDS, gin.H{
					"data": data,
				})
			},
			EVENT_GROUPS: func(s string) {
				data := dao.GetGroupList(c.userId)
				c.send <- encodeJSON(EVENT_GROUPS, gin.H{
					"data": data,
				})
			},
			EVENT_MESSAGE: func(s string) {
				var query models.MssageQuery
				decodeJSON(s, &query)
				c.roomType = query.ItemType
				c.roomId = query.ItemId
				data := dao.GetMessageList(c.userId, query.StartTime, query.ItemType, query.ItemId)
				c.send <- encodeJSON(EVENT_MESSAGE, gin.H{
					"data": data,
				})
			},
		}
		if hanleFunc, ok := hanleFuncMap[event]; ok {
			hanleFunc(message)
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// 使用“&”分割获取房间号
			// 聊天内容不得包含&字符
			// msg[0]为房间号 msg[1]为打印内容
			// msg := strings.Split(string(message), "&")
			// if msg[0] == string(c.hub.roomID[c]) {
			// 	w.Write([]byte(msg[1]))
			// }
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	if msg[0] == string(c.hub.roomID[c]) {
			// 		w.Write(newline)
			// 		w.Write(<-c.send)
			// 	}
			// }
			if err := w.Close(); err != nil {
				log.Printf("error: %v", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, c *gin.Context) {
	// 获取redis连接(暂未使用)
	// pool := c.MustGet("test").(*redis.Pool)
	// redisConn := pool.Get()
	// defer redisConn.Close()
	// 将网络请求变为websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  maxMessageSize,
		WriteBufferSize: maxMessageSize,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), userId: 0}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
