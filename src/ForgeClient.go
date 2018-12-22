package golang_forex_quotes

func CreateForgeClient(apiKey string) *ForgeClient {
	return &ForgeClient{
		apiKey:       apiKey,
		restClient:   CreateRestClient(apiKey),
		socketClient: *CreateSocketClient(apiKey),
	}
}

// REST
func (c *ForgeClient) GetQuotes(symbols []string) ([]Quote, error) {
	return c.restClient.GetQuotes(symbols)
}

func (c *ForgeClient) GetQuota() (Quota, error) {
	return c.restClient.GetQuota()
}

func (c *ForgeClient) Convert(from string, to string, quantity int) (ConversionResult, error) {
	return c.restClient.Convert(from, to, quantity)
}

func (c *ForgeClient) GetMarketStatus() (MarketStatus, error) {
	return c.restClient.GetMarketStatus()
}

func (c *ForgeClient) GetSymbols() ([]string, error) {
	return c.restClient.GetSymbols()
}

// SOCKET
func (c *ForgeClient) Connect() {
	c.socketClient.Connect()
}

func (c *ForgeClient) Disconnect() {
	c.socketClient.Disconnect()
}

func (c *ForgeClient) SubscribeTo(symbols []string) {
	c.socketClient.SubscribeTo(symbols)
}

func (c *ForgeClient) SubscribeToAll() {
	c.socketClient.SubscribeToAll()
}

func (c *ForgeClient) UnsubscribeFrom(symbols []string) {
	c.socketClient.UnsubscribeFrom(symbols)
}

func (c *ForgeClient) UnsubscribeFromAll() {
	c.socketClient.UnsubscribeFromAll()
}

func (c *ForgeClient) OnConnection(callback func()) {
	c.socketClient.OnConnection(callback)
}

func (c *ForgeClient) OnDisconnection(callback func()) {
	c.socketClient.OnDisconnection(callback)
}

func (c *ForgeClient) OnUpdate(callback func(Quote)) {
	c.socketClient.OnUpdate(callback)
}

func (c *ForgeClient) OnMessage(callback func(string)) {
	c.socketClient.OnMessage(callback)
}

func (c *ForgeClient) OnLoginSuccess(callback func()) {
	c.socketClient.OnLoginSuccess(callback)
}
