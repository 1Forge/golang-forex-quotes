/*
 * This library is provided without warranty under the MIT license
 * Created by Jacob Davis <jacob@1forge.com>
 */

package golang_forex_quotes

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"strings"
	"errors"
)

type Conversion struct {
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

type Quote struct {
	Symbol    string
	Bid       float32
	Ask       float32
	Price     float32
	Timestamp int
}

type ApiError struct {
	Error   bool
	Message string
}

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) Client {
	return Client{ApiKey: apiKey}
}

func fetch(query string, apiKey string) ([]byte, error) {
	response, e := http.Get("http://forex.1forge.com/1.0.3/" + query + "&api_key=" + apiKey)

	if e != nil {
		return nil, e
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func unableToUnmarshal(response []byte) error {
	apiError := ApiError{}

	e := json.Unmarshal(response, &apiError)

	if e != nil {
		return e
	}

	return errors.New("forex.1forge.com rejected your request: " + apiError.Message)
}

func (u UnlimitedQuota) toQuota() Quota {
	return Quota{
		QuotaUsed:       u.QuotaUsed,
		QuotaLimit:      0,
		QuotaRemaining:  0,
		HoursUntilReset: u.HoursUntilReset,
	}
}

func (c Client) GetQuota() (Quota, error) {
	result, e := fetch("quota?cache=false", c.ApiKey)

	quota := Quota{}

	if e != nil {
		return quota, e
	}

	//Able to unmarshal to Quota
	if json.Unmarshal(result, &quota) == nil {
		return quota, nil
	}

	unlimitedQuota := UnlimitedQuota{}

	//Unable to unmarshal to Quota, try UnlimitedQuota
	if json.Unmarshal(result, &unlimitedQuota) == nil {
		return unlimitedQuota.toQuota(), nil
	}

	//Unable to unmarshal at all
	return quota, unableToUnmarshal(result)
}

func (c Client) GetSymbols() ([]string, error) {
	result, e := fetch("symbols?cache=false", c.ApiKey)

	if e != nil {
		return nil, e
	}

	symbolList := []string{}

	if json.Unmarshal(result, &symbolList) != nil {
		return symbolList, unableToUnmarshal(result)
	}

	return symbolList, nil
}

func (c Client) GetQuotes(symbols []string) ([]Quote, error) {
	result, e := fetch("quotes?pairs="+strings.Join(symbols, ","), c.ApiKey)

	if e != nil {
		return nil, e
	}

	quotes := []Quote{}

	if json.Unmarshal(result, &quotes) != nil {
		return quotes, unableToUnmarshal(result)
	}

	return quotes, nil
}

func (c Client) Convert(from string, to string, quantity int) (Conversion, error) {
	result, e := fetch("convert?from="+from+"&to="+to+"&quantity="+strconv.Itoa(quantity), c.ApiKey)

	Conversion := Conversion{}

	if e != nil {
		return Conversion, e
	}

	if json.Unmarshal(result, &Conversion) != nil {
		return Conversion, unableToUnmarshal(result)
	}

	return Conversion, nil
}

func (c Client) MarketIsOpen() (bool, error) {
	result, e := fetch("market_status?cache=false", c.ApiKey)

	if e != nil {
		return false, e
	}

	marketStatus := MarketStatus{}

	if json.Unmarshal(result, &marketStatus) != nil {
		return false, unableToUnmarshal(result)
	}

	return marketStatus.MarketIsOpen, nil
}