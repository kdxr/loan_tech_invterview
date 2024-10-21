package services

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type TSocketClient struct {
	name string
	ws   *websocket.Conn
}

// var SocketClient = make(map[string]*websocket.Conn)
var SocketClient = []TSocketClient{}

func EmitSocketMessage(message interface{}) {

	for _, socket := range SocketClient {
		// fmt.Printf("Emit Socket Message To : %s", socket.name)

		// socket.ws.WriteMessage(websocket.TextMessage, []byte(message))
		socket.ws.WriteJSON(message)
	}
}

func InitSocket() {

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:name", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		// log.Println(c.Locals("allowed"))  // true
		// log.Println(c.Query("name")) // 1.0
		// log.Println(c.Cookies("session")) // ""

		name := c.Params("name")

		if name == "" {
			c.Close()
			return
		}

		// SocketClient[name] = c
		SocketClient = append(SocketClient, TSocketClient{
			name: name,
			ws:   c,
		})

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s (%d)", msg, mt)

			// if err = c.WriteMessage(mt, []byte("helloworld")); err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }

			// log.Println("Send Message back")
		}

	}))

	go func() {
		log.Fatal(app.Listen(":5001"))
	}()

}
