package chat

import (
	"net/http"
	"strings"
	"zodream/modules/chat/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

var events = websocket.Namespaces{
	"default": websocket.Events{
		websocket.OnRoomJoined: onRoomJoined,
		websocket.OnRoomLeft:   onRoomLeft,
	},
}

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	setupWebsocket(app)
}

func setupWebsocket(app iris.Party) {
	// create our echo websocket server
	ws := websocket.New(websocket.DefaultGorillaUpgrader, events)
	ws.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {

		return r.RemoteAddr[:strings.IndexByte(r.RemoteAddr, ':')]
	}
	app.Get("/ws", websocket.Handler(ws))
}

func onRoomJoined(ns *websocket.NSConn, msg websocket.Message) error {
	// the roomName here is the source.
	pageSource := string(msg.Room)

	// fire the "onNewVisit" client event
	// on each connection joined to this room (source page)
	// and notify of the new visit,
	// including this connection (see nil on first input arg).
	ns.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Room:      pageSource,
		Event:     "onNewVisit", // fire the "onNewVisit" client event.
		Body:      []byte("1"),
	})

	return nil
}

func onRoomLeft(ns *websocket.NSConn, msg websocket.Message) error {
	// the roomName here is the source.
	// pageV := v.Get(msg.Room)

	// fire the "onNewVisit" client event
	// on each connection joined to this room (source page)
	// and notify of the new, decremented by one, visits count.
	ns.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Room:      msg.Room,
		Event:     "onNewVisit",
		Body:      []byte("7"),
	})

	return nil
}
