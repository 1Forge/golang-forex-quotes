package golang_forex_quotes

import io "github.com/graarh/golang-socketio"

const (
	LOGIN                = "login"
	SUBSCRIBE_TO         = "subscribe_to"
	UNSUBSCRIBE_FROM     = "unsubscribe_from"
	SUBSCRIBE_TO_ALL     = "subscribe_to_all"
	UNSUBSCRIBE_FROM_ALL = "unsubscribe_from_all"
	MESSAGE              = "message"
	FORCE_CLOSE          = "force_close"
	POST_LOGIN_SUCCESS   = "post_login_success"
	UPDATE               = "update"
)

type Quote struct {
	Symbol    string
	Bid       float32
	Ask       float32
	Price     float32
	Timestamp int
}

type SocketClient struct {
	ApiKey               string
	IO                   *io.Client
	connectCallback      func()
	disconnectCallback   func()
	messageCallback      func(string)
	updateCallback       func(Quote)
	loginSuccessCallback func()
}

type ForgeClient struct {
	apiKey       string
	restClient   RestClient
	socketClient SocketClient
}

type ConversionResult struct {
	Value     float32
	Text      string
	Timestamp int
}

type Quota struct {
	QuotaUsed       int `json:"quota_used"`
	QuotaLimit      int `json:"quota_limit"` //0 = Unlimited
	QuotaRemaining  int `json:"quota_remaining"`
	HoursUntilReset int `json:"hours_until_reset"`
}

type UnlimitedQuota struct {
	QuotaUsed       int    `json:"quota_used"`
	QuotaLimit      string `json:"quota_limit"`
	QuotaRemaining  string `json:"quota_remaining"`
	HoursUntilReset int    `json:"hours_until_reset"`
}

type MarketStatus struct {
	MarketIsOpen bool `json:"market_is_open"`
}

type RestError struct {
	Error   bool
	Message string
}

type RestClient struct {
	ApiKey string
}
