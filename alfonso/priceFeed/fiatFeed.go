package priceFeed

import (
	"net/http"
	"time"
)

/*
{
	"success":true,
	"terms":"https:\/\/currencylayer.com\/terms",
	"privacy":"https:\/\/currencylayer.com\/privacy",
	"timestamp":1504027454,
	"source":"USD",
	"quotes":{"USDPHP":51.080002}
}
*/

type fiatAPIReturn struct {
	Quotes map[string]float64
}

type fiatFeed struct {
	url    string
	client http.Client
}

// ensure that it implements priceFeed
var _ priceFeed = &fiatFeed{}

func newFiatFeed(url string) *fiatFeed {
	//log.Info("newFiatFeed: ", url)
	m := new(fiatFeed)
	m.url = url
	m.client = http.Client{Timeout: 10 * time.Second}

	return m
}

func (self *fiatFeed) getPrice() (float64, error) {
	var ret fiatAPIReturn
	err := getJSON(self.client, self.url, &ret)
	if err != nil {
		return 0, err
	}
	var pA float64
	for _, value := range ret.Quotes {
		//log.Info("value:", value)
		pA = value
	}

	return (1.0 / pA), nil
}
