package golang_forex_quotes

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/sacOO7/gowebsocket"
	// io "github.com/graarh/golang-socketio"
	// "github.com/graarh/golang-socketio/transport"
)

func CreateSocketClient(apiKey string) *SocketClient {
	c := SocketClient{
		ApiKey: apiKey,
	}

	return &c
}

func (c *SocketClient) Connect() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("wss://sockets.1forge.com/socket")
	// socket.EnableLogging()

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Received connect error - ", err)
	}

	socket.OnConnected = func(socket gowebsocket.Socket) {
		// log.Println("Connected to server")
		socket.SendText(LOGIN + "|" + c.ApiKey)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		parts := strings.SplitN(message, "|", 2)
		switch parts[0] {
		case POST_LOGIN_SUCCESS:
			c.handlePostLoginSuccess()
		case UPDATE:
			var quotes Quote
			json.Unmarshal([]byte(parts[1]), &quotes)
			// fmt.Printf("%+v", quotes)
			c.handleUpdate(quotes)
		case MESSAGE:
			c.handleMessage(message)
		default:
			log.Println(parts)
			// log.Println("Received message - " + message)
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	c.socket = &socket

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}

func (c *SocketClient) OnMessage(callback func(string)) {
	c.messageCallback = callback
}

func (c *SocketClient) OnConnection(callback func()) {
	c.connectCallback = callback
}

func (c *SocketClient) OnDisconnection(callback func()) {
	c.disconnectCallback = callback
}

func (c *SocketClient) OnUpdate(callback func(Quote)) {
	c.updateCallback = callback
}

func (c *SocketClient) OnLoginSuccess(callback func()) {
	c.loginSuccessCallback = callback
}

func (c *SocketClient) SubscribeTo(symbols []string) {
	for _, symbol := range symbols {
		// log.Println("Subscribing: ", symbol)
		c.socket.SendText(SUBSCRIBE_TO + "|" + symbol)
	}
}

func (c *SocketClient) SubscribeToAll() {
	c.socket.SendText(SUBSCRIBE_TO_ALL + "|" + "")
}

func (c *SocketClient) UnsubscribeFrom(symbols []string) {
	for _, symbol := range symbols {
		c.socket.SendText(UNSUBSCRIBE_FROM + "|" + symbol)
	}
}

func (c *SocketClient) UnsubscribeFromAll() {
	c.socket.SendText(UNSUBSCRIBE_FROM_ALL + "|" + "")
}

// func (c *SocketClient) handleLoginRequest(h *io.Channel) {
// 	c.login()
// }

func (c *SocketClient) handlePostLoginSuccess() {
	if c.loginSuccessCallback != nil {
		c.loginSuccessCallback()
	}
}

func (c *SocketClient) handleMessage(m string) {
	if c.messageCallback != nil {
		c.messageCallback(m)
	}
}

func (c *SocketClient) handleDisconnect() {
	if c.disconnectCallback != nil {
		c.disconnectCallback()
	}
}

func (c *SocketClient) handleConnect() {
	if c.connectCallback != nil {
		c.connectCallback()
	}
}

func (c *SocketClient) handleUpdate(q Quote) {
	if c.updateCallback != nil {
		c.updateCallback(q)
	}
}

func (c *SocketClient) Disconnect() {
	// c.IO.Close()
}
