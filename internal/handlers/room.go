package handlers

import (
	"fmt"
	"os"
	"time"


    w "vidoechat/pgk/webrtc"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/contrib/websocket"
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
	uuid,guuid,_ := createOrGetRoom(uuid)
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

}

type websocketMessage struct{
	Event string `json:"event"`
	Data string `json:"data"`
}