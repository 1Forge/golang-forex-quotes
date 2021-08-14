package golang_forex_quotes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CreateRestClient(apiKey string) RestClient {
	return RestClient{
		ApiKey: apiKey,
	}
}

func fetch(query string, apiKey string) ([]byte, error) {
	if strings.Count(query, "") > 7683 {
		// println("String count", strings.Count(query, ""))
		// err := errors.New("No more than 865 pairs or 1730 curriencies")
		err := errors.New("No more than 949 pairs or 1898 curriencies")
		return nil, err
	} else {
		response, e := http.Get("https://api.1forge.com/" + query + "&api_key=" + apiKey)

		if e != nil {
			return nil, e
		}

		defer response.Body.Close()

		return ioutil.ReadAll(response.Body)
	}
}

func unableToUnmarshal(response []byte) error {
	apiError := RestError{}

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

func (c RestClient) GetQuota() (Quota, error) {
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

func (c RestClient) GetSymbols() ([]string, error) {
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

func (c RestClient) GetQuotes(symbols []string) ([]Quote, error) {
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

func (c RestClient) Convert(from string, to string, quantity int) (ConversionResult, error) {
	result, e := fetch("convert?from="+from+"&to="+to+"&quantity="+strconv.Itoa(quantity), c.ApiKey)

	conversion := ConversionResult{}

	if e != nil {
		return conversion, e
	}

	if json.Unmarshal(result, &conversion) != nil {
		return conversion, unableToUnmarshal(result)
	}

	return conversion, nil
}

func (c RestClient) GetMarketStatus() (MarketStatus, error) {
	result, e := fetch("market_status?cache=false", c.ApiKey)

	marketStatus := MarketStatus{}

	if e != nil {
		return marketStatus, e
	}

	if json.Unmarshal(result, &marketStatus) != nil {
		return marketStatus, unableToUnmarshal(result)
	}

	return marketStatus, nil
}
