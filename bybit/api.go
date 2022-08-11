package bybit

import (
	"fmt"
	"log"
	"time"
)

/* =============== Public API ===================================== */
func GetServerTime() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Second)
}

func GetCandlePlus(param CandleParam) ([]Candle, error) {
	quotient := int(param.Amount/200) + 1
	remainder := int(param.Amount % 200)

	candles := []Candle{}

	for i := quotient; i >= 0; i-- {
		c, _ := GetCandle(CandleParam{
			Code:   param.Code,
			Unit:   param.Unit,
			Amount: 200 * i,
		})

		candles = append(candles, c...)
	}

	return candles[200-remainder:], nil
}

func GetCandle(param CandleParam) ([]Candle, error) {
	var ret CandleReponse

	maxSecond := param.Unit * 60 * param.Amount
	param.From = fmt.Sprintf("%d", GetServerTime()-int64(maxSecond))

	c := NewWrapper("", "")
	if err := c.PublicRequest("/public/linear/kline", param, &ret); err != nil {
		return []Candle{}, err
	}

	return ret.Result, nil
}

/*=================== Wallet with Account ================================ */
func GetBalance(auth APIAuth, param BalanceParam) (*Balance, error) {
	var ret BalanceResponse

	c := NewWrapper(auth.Token())
	if err := c.PrivateRequest("/v2/private/wallet/balance", param, &ret); err != nil {
		return &Balance{}, err
	}

	var result Balance
	switch param.Coin {
	case "USDT":
		result = ret.Result.USDT
	case "BTC":
		result = ret.Result.BTC
	default:
		log.Println("This coin is not supported.")
	}

	return &result, nil
}

func GetPosition(auth APIAuth, param PositionParam) ([]Position, error) {
	var ret PositionResponse

	c := NewWrapper(auth.Token())
	if err := c.PrivateRequest("/private/linear/position/list", param, &ret); err != nil {
		return nil, err
	}

	return ret.Result, nil
}

/*=================== Active Order with Account ================================ */
func Place(auth APIAuth, param OrderParam) (*Order, error) {
	var ret OrderResponse

	c := NewWrapper(auth.Token())
	if err := c.PostRequest("/private/linear/order/create", param, &ret); err != nil {
		return nil, err
	}

	return &ret.Result, nil
}

func Replace(auth APIAuth, param ReplaceOrderParam) (*ReplaceOrder, error) {
	var ret ReplaceOrderResponse

	c := NewWrapper(auth.Token())
	if err := c.PostRequest("/private/linear/order/replace", param, &ret); err != nil {
		return nil, err
	}

	return &ret.Result, nil
}

func Cancel(auth APIAuth, param CancelOrderParam) (*CancelOrder, error) {
	var ret CancelOrderResponse

	c := NewWrapper(auth.Token())
	if err := c.PostRequest("/private/linear/order/cancel", param, &ret); err != nil {
		return nil, err
	}

	return &ret.Result, nil
}

func Search(auth APIAuth, param ActiveOrderParam) ([]ActiveOrder, error) {
	var ret ActiveOrderResponse

	c := NewWrapper(auth.Token())
	if err := c.PrivateRequest("/private/linear/order/search", param, &ret); err != nil {
		return nil, err
	}

	return ret.Result, nil
}
