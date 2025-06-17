package handlers

import (
	"fmt"
	"os"

	w "vidoechat/pgk/webrtc"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v3"

	guuid "github.com/google/uuid"
)



func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}


func Room (c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == ""{
		c.Status(400)
		return nil
	}
	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION"{
		ws = "wss"
	}
	uuid,suuid,_ := createOrGetRoom(uuid)
	return c.Render("peer", fiber.Map{
		"RoomWebsocketAddr": fmt.Sprintf("%s://%s/room%s/websocket", ws, c.Hostname(), uuid),
		"RoomLink": fmt.Sprintf("%s://%s/room/%s",c.Protocol(), c.Hostname(), uuid),
		"ChatWebSocketAddr":fmt.Sprintf("%s://%s/room/%s/chat/websocket", ws, c.Hostname(), uuid),
		"ViewerWebSocketAddr": fmt.Sprintf("%s://%s/room/%s/viewer/websocket", ws, c.Hostname(), uuid),
		"StreamLink": fmt.Sprintf("%s://%s/steam/%s", c.Protocol(), c.Hostname(), suuid),
		"Type": "room",

	}, "layouts/main")
}

func RoomWebsocket(c *websocket.Conn){
	uuid := c.Params("uuid")
	if uuid == ""{
		return
	}

	_, _, room :=  createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}


func createOrGetRoom(uuid string)(string, string, *w.Room){

}

func RoomViewerWebsocket(c *websocket.Conn, p *w.Peers){
      uuid := c.Params("uuid")
	  if uuid == ""{
		return
	  }

	  w.RoomsLock.Lock()
	  if peer, ok := w.Rooms[uuid]; ok{
		w.RoomsLock.Unlock()
		roomViewerConn(c, peer.Peers)
		return
	  }
	 w.RoomsLock.Unlock()
}

func roomViewerConn(c *websocket.Conn, p *w.Peers){
	ticker :=time.NewTIcker(1*time.Second)
	defer ticker.Stop()
	defer c.CLose()

	for{
		select{
		case <- ticker.C:
			w, err := c.Conn.NexWriter(websocket.TextMessage)
			if err != nil{
				return
			}

			w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
		}
	}
}


type websocketMessage struct{
	Event string `json:"event"`
	Data string `json:"data"`
}