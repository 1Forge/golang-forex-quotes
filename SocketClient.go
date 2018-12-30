package golang_forex_quotes

import (
	"log"

	io "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func CreateSocketClient(apiKey string) *SocketClient {
	c := SocketClient{
		ApiKey: apiKey,
	}

	return &c
}

func (c *SocketClient) Connect() {
	IO, e := io.Dial(
		io.GetUrl("socket.forex.1forge.com", 3000, true),
		transport.GetDefaultWebsocketTransport(),
	)

	if e != nil {
		log.Fatal(e)
	}

	err := IO.On(LOGIN, c.handleLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	err = IO.On(MESSAGE, c.handleMessage)
	if err != nil {
		log.Fatal(err)
	}

	err = IO.On(POST_LOGIN_SUCCESS, c.handlePostLoginSuccess)
	if err != nil {
		log.Fatal(err)
	}

	err = IO.On(UPDATE, c.handleUpdate)
	if err != nil {
		log.Fatal(err)
	}

	err = IO.On(io.OnDisconnection, c.handleDisconnect)
	if err != nil {
		log.Fatal(err)
	}

	err = IO.On(io.OnConnection, c.handleConnect)
	if err != nil {
		log.Fatal(err)
	}

	c.IO = IO
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

func (c *SocketClient) login() {
	e := c.IO.Emit(LOGIN, c.ApiKey)
	if e != nil {
		log.Fatal(e)
	}
}

func (c *SocketClient) SubscribeTo(symbols []string) {
	for _, symbol := range symbols {
		e := c.IO.Emit(SUBSCRIBE_TO, symbol)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func (c *SocketClient) SubscribeToAll() {
	e := c.IO.Emit(SUBSCRIBE_TO_ALL, "")
	if e != nil {
		log.Fatal(e)
	}
}

func (c *SocketClient) UnsubscribeFrom(symbols []string) {
	for _, symbol := range symbols {
		e := c.IO.Emit(UNSUBSCRIBE_FROM, symbol)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func (c *SocketClient) UnsubscribeFromAll() {
	e := c.IO.Emit(UNSUBSCRIBE_FROM_ALL, "")
	if e != nil {
		log.Fatal(e)
	}
}

func (c *SocketClient) handleLoginRequest(h *io.Channel) {
	c.login()
}

func (c *SocketClient) handlePostLoginSuccess(h *io.Channel) {
	if c.loginSuccessCallback != nil {
		c.loginSuccessCallback()
	}
}

func (c *SocketClient) handleMessage(h *io.Channel, m string) {
	if c.messageCallback != nil {
		c.messageCallback(m)
	}
}

func (c *SocketClient) handleDisconnect(h *io.Channel) {
	if c.disconnectCallback != nil {
		c.disconnectCallback()
	}
}

func (c *SocketClient) handleConnect(h *io.Channel) {
	if c.connectCallback != nil {
		c.connectCallback()
	}
}

func (c *SocketClient) handleUpdate(h *io.Channel, q Quote) {
	if c.updateCallback != nil {
		c.updateCallback(q)
	}
}

func (c *SocketClient) Disconnect() {
	c.IO.Close()
}
