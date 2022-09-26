package bybit

import (
	"errors"
	"time"
)

func (c *Client) GetAPI() API {
	return c
}

type Client struct {
	auth APIAuth
}

func NewClient(opt Options) *Client {
	return &Client{
		auth: BybitAuth(opt.AccessKey, opt.SecretKey),
	}
}

func (Client) GetServerTime() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Second)
}

func (Client) GetCandle(code string, unit, amount int) ([]Candle, error) {
	res, err := GetCandle(CandleParam{Code: "BTCUSDT", Unit: unit, Amount: amount})
	if err != nil {
		return []Candle{}, errors.New("Get Candle Error")
	}

	return res, nil
}

func (c Client) GetCandlePlus(code string, unit, amount int) ([]Candle, error) {
	quotient := int(amount/200) + 1
	remainder := int(amount % 200)

	candles := []Candle{}

	for i := quotient; i >= 0; i-- {
		c, _ := c.GetCandle(code, unit, 200*i)
		candles = append(candles, c...)
	}

	return candles[200-remainder:], nil
}

func (c Client) GetBalance(coin string) (*Balance, error) {
	ret, err := GetBalance(c.auth, BalanceParam{Coin: coin})
	if err != nil {
		return nil, errors.New("Get Balance Error")
	}

	return ret, nil
}

func (c Client) GetPosition(code string) ([]Position, error) {
	res, err := GetPosition(c.auth, PositionParam{Code: "BTCUSDT"})
	if err != nil {
		return nil, errors.New("Get Position Error")
	}
	return res, nil
}

func (c Client) PlaceClear(code string, position ClearSide, price, amount float64, force bool) (*Order, error) {
	res, err := Place(c.auth, NewClearParam(code, position, price, amount, force))
	if err != nil {
		return nil, errors.New("Place Error")
	}
	return res, nil
}

func (c Client) PlaceBetting(code string, position BettingSide, price, amount float64, force bool) (*Order, error) {
	res, err := Place(c.auth, NewBettingParam(code, position, price, amount, force))
	if err != nil {
		return nil, errors.New("Place Error")
	}
	return res, nil
}

func (c Client) Replace(code, id string, price, amount float64) (*ReplaceOrder, error) {
	res, err := Replace(c.auth, ReplaceOrderParam{Code: code, UUID: id, Price: price, Amount: amount})
	if err != nil {
		return nil, errors.New("Replace Error")
	}

	return res, nil
}

func (c Client) Cancel(code, id string) (*CancelOrder, error) {
	return Cancel(c.auth, CancelOrderParam{UUID: id, Code: code})
}

func (c Client) Search(code string) ([]ActiveOrder, error) {
	return Search(c.auth, ActiveOrderParam{Code: code})
}
