package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"syscall/js"

	"time"

	"github.com/seanburman/game/config"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Client struct {
	Origin     string
	ServerAddr string
	Conn       *websocket.Conn
}

type Res struct {
	Name    string
	Message string
}

func NewClient() *Client {
	return &Client{
		Origin:     "http://localhost/",
		ServerAddr: "ws://localhost:3000",
	}
}

func (c *Client) Dial(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	u := fmt.Sprintf("ws://%s%s", config.Env().HOST, url)
	ws, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer ws.CloseNow()
	c.Conn = ws

	err = wsjson.Write(ctx, ws, "Hello there")
	if err != nil {
		fmt.Println("failed to send")
	}

	// TODO:  subscribe
	ws.Close(websocket.StatusNormalClosure, "")
}

func (c Client) GetHTML() string {
	u := fmt.Sprintf("%s:5500/index.html", config.Env().HOST)
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	fmt.Println(string(content))
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func (c Client) Handshake() bool {
	id := js.Global().Call("func")
	fmt.Println("WELL THEN")
	fmt.Printf("%s connected to server...", id)
	if id.Truthy() {
		config.Message.Set(fmt.Sprint(id) + "connected to server...")
		return true
	}
	log.Fatal("connection to server failed...")
	return false
}

// findConcatByteString searches 'b' for a []byte of provided length following the presence of 's',
// returning the desired bytes from b[index of s[0] + len(s):index of s[0]+len(s)+length]
//
// This is used for parsing codes appended to a shared secret for authentication
func findConcatByteString(b []byte, s string, length int) ([]byte, bool) {
	var buf = make([]byte, len(s)+length)

	for _, char := range b {
		fmt.Println("Handshake buf: ", string(buf))
		// check if buf contains s
		if len(buf) == len(s)+length {
			if string(buf[:len(s)]) == s {
				// return remainder of buf to length
				fmt.Println(string(buf))
				fmt.Println(string(buf[len(s):]))
				return buf[len(s):], true
			}
			buf = buf[1:]
			buf = append(buf, byte(char))
		}
	}
	return []byte{}, false
}

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
