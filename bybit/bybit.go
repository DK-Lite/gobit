package bybit

import (
	"errors"
	"time"
)

type API interface {
	// Public API
	GetServerTime() int64
	GetCandle(code string, unit, amount int) ([]Candle, error)
	GetCandlePlus(code string, unit, amount int) ([]Candle, error)

	// Private API with Auth
	GetBalance(coin string) (float64, error)
	GetPosition(code string) (float64, error)

	// Active Order with Auth
	PlaceClear(code string, position ClearSide, price, amount float64, force bool) (string, error)
	PlaceBetting(code string, position BettingSide, price, amount float64, force bool) (string, error)
	Replace(code, id string, price, amount float64) (string, error)
	Cancel(code, id string) (*CancelOrder, error)
	Search(code string) ([]ActiveOrder, error)
}

type BybitClient struct {
	auth APIAuth
}

func NewClient(accessKey, secretKey string) API {
	return &BybitClient{
		auth: BybitAuth(accessKey, secretKey),
	}
}

func (BybitClient) GetServerTime() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Second)
}

func (BybitClient) GetCandle(code string, unit, amount int) ([]Candle, error) {
	res, err := GetCandle(CandleParam{Code: "BTCUSDT", Unit: unit, Amount: amount})
	if err != nil {
		return []Candle{}, errors.New("Get Candle Error")
	}

	return res, nil
}

func (c BybitClient) GetCandlePlus(code string, unit, amount int) ([]Candle, error) {
	quotient := int(amount/200) + 1
	remainder := int(amount % 200)

	candles := []Candle{}

	for i := quotient; i >= 0; i-- {
		c, _ := c.GetCandle(code, unit, 200*i)
		candles = append(candles, c...)
	}

	return candles[200-remainder:], nil
}

func (bybit BybitClient) GetBalance(coin string) (float64, error) {
	ret, err := GetBalance(bybit.auth, BalanceParam{Coin: coin})
	if err != nil {
		return 0, errors.New("Get Balance Error")
	}

	return ret.Equity, nil
}

func (bybit BybitClient) GetPosition(code string) (float64, error) {
	res, err := GetPosition(bybit.auth, PositionParam{Code: "BTCUSDT"})
	if err != nil {
		return 0, errors.New("Get Position Error")
	}
	return res[0].Size - res[1].Size, nil
}

func (bybit BybitClient) PlaceClear(code string, position ClearSide, price, amount float64, force bool) (string, error) {
	res, err := Place(bybit.auth, NewClearParam(code, position, price, amount, force))
	if err != nil {
		return "", errors.New("Place Error")
	}
	return res.UUID, nil
}

func (bybit BybitClient) PlaceBetting(code string, position BettingSide, price, amount float64, force bool) (string, error) {
	res, err := Place(bybit.auth, NewBettingParam(code, position, price, amount, force))
	if err != nil {
		return "", errors.New("Place Error")
	}
	return res.UUID, nil
}

func (bybit BybitClient) Replace(code, id string, price, amount float64) (string, error) {
	res, err := Replace(bybit.auth, ReplaceOrderParam{Code: code, UUID: id, Price: price, Amount: amount})
	if err != nil {
		return "", errors.New("Replace Error")
	}

	return res.UUID, nil
}

func (bybit BybitClient) Cancel(code, id string) (*CancelOrder, error) {
	return Cancel(bybit.auth, CancelOrderParam{UUID: id, Code: code})
}

func (bybit BybitClient) Search(code string) ([]ActiveOrder, error) {
	return Search(bybit.auth, ActiveOrderParam{Code: code})
}
