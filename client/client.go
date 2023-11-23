package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Client struct {
	Origin     string
	ServerAddr string
	Socket     string
}

type Res struct {
	Name    string
	Message string
}

func NewClient(socket string) *Client {
	return &Client{
		Origin:     "http://localhost/",
		ServerAddr: "ws://localhost:3000",
	}
}

func (c *Client) OnPage() {
	res, err := http.Get("http://127.0.0.1:5500/index.html")
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func (c *Client) Dial(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ws, _, err := websocket.Dial(ctx, "ws://192.168.1.154:3000/echo", &websocket.DialOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer ws.CloseNow()

	// msg, err := json.Marshal(Res{Name: "Sean"})
	// if err != nil {
	// 	fmt.Println("failed to marshal")
	// }
	err = wsjson.Write(ctx, ws, "Hello from mobile")
	if err != nil {
		fmt.Println("failed to send")
	}

	ws.Close(websocket.StatusNormalClosure, "")
}

// func (c *Client) HandleAuth(ws *websocket.Conn) {
// 	fmt.Println("new incoming connection from client:", ws.RemoteAddr())
// 	go c.readLoop(ws)
// 	ws.Write([]byte("Authenticated"))
// }

// func (c *Client) Dial(endpoint string) {
// 	cfg, err := websocket.NewConfig(c.ServerAddr+endpoint, c.Origin)
// 	cfg.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o")
// 	if err != nil {
// 		log.Fatal("config error:\n", err)
// 	}
// 	ws, err := websocket.DialConfig(cfg)
// 	if err != nil {
// 		log.Fatal("connection to client failed:\n", err)
// 	} else {
// 		fmt.Printf("connected to socket server %s\n", ws.RemoteAddr())
// 	}
// 	go func(ws *websocket.Conn) {
// 		c.readLoop(ws)
// 	}(ws)
// }

// func (c *Client) readLoop(ws *websocket.Conn) {
// 	buf := make([]byte, 512)
// 	for {
// 		n, err := ws.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				log.Printf("disconnected from socket %s", ws.RemoteAddr())
// 				os.Exit(1)
// 			}
// 		}
// 		msg := buf[:n]
// 		config.Message.Set(fmt.Sprintf("Message :%s", msg[:n]))
// 		fmt.Printf("Message: %s\n", msg[:n])
// 	}

// }
