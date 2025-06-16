package server


import (
	"flag"
	"os"
	"time"
     w"vidoechat/pkg/webrtc"
	 "vidoechat/internal/handlers"
	 "github.com/gofiber/fiber/v3"
	  "github.com/gofiber/fiber/v3/middleware/cors"

    "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/contrib/websocket"
)

var(
	addr = flag.String("addr",":"+ os.Getenv("PORT"), "")
	cert = flag.String("cert", "","")
	key= flag.String("key", "", "")
)

func Run() error{
	flag.Parse()

	if *addr ==":"{
		*addr = ":8080"
	}


	engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{Views:engine})
	app.Use(logger.New())
	app.Use(cors.New())


	app.Get("/", handlers.Welcome)
	app.Get("/room/create",handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
   app.Get("/room/:uuid/websocket",websocket.New(handlers.RoomWebsocket, websocket.Config{
	HandShakeTimeout : 10*time.second,
   }))
   app.Get("/room/:uuid/chat", handlers.RoomChat)
   app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
   app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
   app/Get("/stream/:ssiud", handlers.Stream)
   app/Get("/stream/:ssiud/websocket",websocket.New(handlers.StreamWebSocket.Config{HandShakeTimeout: 10*time.Second,}))
   app/Get("/stream/:ssiud/chat/websocket", websocket.New(handlers.StreamChatWebsocket))
   app/Get("/stream/:ssiud/viewer/websocket", websocket.New(handlers.StreamViewWebsocket))
   app.Static("/","./assets")

  

   w.Rooms = make(map[string]*w.Room)
  w.Streams = make(map[string]* w.Room)
 
  go dispatchKeyFrames()

  if * cert != ""{
	return app.ListenTLS(*addr, *cert, *key)
  }
  return app.Listen(*addr)



}



   func dispatchKeyFrames(){
     for range time.NewTicker(time.Second *3).c{
		for _, room := range w.Rooms
		room.Peers.DispatchKeyFrames()
	 }
   }

