package bybit

/* BASE */
type BaseReponse struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	ExtInfo          string      `json:"ext_info"`
	Result           interface{} `json:"result"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}
type Time struct{}
type TimeOptions struct{}

/* CANDLE */
type CandleReponse struct {
	BaseReponse
	Result []Candle `json:"result"`
}

type Candle struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Period   string  `json:"period"`
	StartAt  int     `json:"start_at"`
	Volume   float64 `json:"volume"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Interval string  `json:"interval"`
	OpenTime int     `json:"open_time"`
	Turnover float64 `json:"turnover"`
}

type CandleParam struct {
	Code   string `url:"symbol"`
	Unit   int    `url:"interval"`
	Amount int    `url:"limit"`
	From   string `url:"from"`
}

/* BALANCE */
type BalanceResponse struct {
	BaseReponse
	Result BalanceResultData `json:"result"`
}

type BalanceResultData struct {
	BTC  Balance `json:"BTC"`
	ETH  Balance `json:"ETH"`
	EOS  Balance `json:"EOS"`
	XRP  Balance `json:"XRP"`
	USDT Balance `json:"USDT"`
}

type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type BalanceParam struct {
	Coin string
}

/* ASSET */
type AssetOptions struct {
	ApiKey    string `url:"api_key"`
	Coin      string `url:"coin"`
	TimeStamp int    `url:"timestamp"`
	Sign      string `url:"sign"`
}

/* POSITION */
type Position struct {
	UserId              int     `json:"user_id"`
	Symbol              string  `json:"symbol"`
	Side                string  `json:"side"`
	Size                float64 `json:"size"`
	PositionValue       float64 `json:"position_value"`
	EntryPrice          float64 `json:"entry_price"`
	LiqPrice            float64 `json:"liq_price"`
	BustPrice           float64 `json:"bust_price"`
	Leverage            float64 `json:"leverage"`
	AutoAddMargin       float64 `json:"auto_add_margin"`
	IsIsolated          bool    `json:"is_isolated"`
	PositionMargin      float64 `json:"position_margin"`
	OccClosingFee       float64 `json:"occ_closing_fee"`
	RealisedPnl         float64 `json:"realised_pnl"`
	CumRealisedPnl      float64 `json:"cum_realised_pnl"`
	FreeQry             float64 `json:"free_qty"`
	TpSlMode            string  `json:"tp_sl_mode"`
	DeleverageIndicator float64 `json:"deleverage_indicator"`
	UnrealisedPnl       float64 `json:"unrealised_pnl"`
	RiskId              int     `json:"risk_id"`
	TakeProfit          float64 `json:"take_profit"`
	StopLoss            float64 `json:"stop_loss"`
	TrailingStop        float64 `json:"trailing_stop"`
}
type PositionResponse struct {
	BaseReponse
	Result []Position `json:"result"`
}
type PositionParam struct {
	Code string `url:"symbol"`
}

/* ATIVE ORDER LIST */
type ActiveOrderResponse struct {
	BaseReponse
	Result []ActiveOrder `json:"result"`
}
type ActiveOrder struct {
	OrderID        string  `json:"order_id"`
	UserID         int     `json:"user_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	LastExecPrice  float64 `json:"last_exec_price"`
	CumExecQty     float64 `json:"cum_exec_qty"`
	CumExecValue   float64 `json:"cum_exec_value"`
	CumExecFee     float64 `json:"cum_exec_fee"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
}

type ActiveOrderParam struct {
	Code string `url:"symbol"`
}

/* CANCEL ORDER */
type CancelOrderResponse struct {
	BaseReponse
	Result CancelOrder `json:"result"`
}
type CancelOrder struct {
	UUID string `json:"order_id"`
}
type CancelOrderParam struct {
	UUID string `json:"order_id"`
	Code string `json:"symbol"`
}

/* PLACE BETTING & CLEAR  */
type BettingSide string
type ClearSide string

const (
	BETTING_LONG  BettingSide = "Buy"
	BETTING_SHORT BettingSide = "Sell"
	CLEAR_LONG    ClearSide   = "Sell"
	CLEAR_SHORT   ClearSide   = "Buy"
)

type OrderResponse struct {
	BaseReponse
	Result Order `json:"result"`
}
type Order struct {
	UUID string `json:"order_id"`
}
type OrderParam struct {
	Side           string  `json:"side"`
	Symbol         string  `json:"symbol"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
}

/* REPLACE ORDER */
type ReplaceOrderResponse struct {
	BaseReponse
	Result ReplaceOrder `json:"result"`
}
type ReplaceOrder struct {
	UUID string `json:"order_id"`
}
type ReplaceOrderParam struct {
	Code   string  `json:"symbol"`
	UUID   string  `json:"order_id"`
	Price  float64 `json:"p_r_price"`
	Amount float64 `json:"p_r_qty"`
}

func getType(force bool) string {
	if force {
		return "Market"
	}

	return "Limit"
}

// NewBettingParam()
func NewBettingParam(code string, side BettingSide, price float64, amount float64, force bool) OrderParam {
	return OrderParam{
		Side:           string(side),
		Symbol:         code,
		OrderType:      getType(force),
		Price:          price,
		Qty:            amount,
		TimeInForce:    "GoodTillCancel",
		ReduceOnly:     false,
		CloseOnTrigger: false,
	}
}

func NewClearParam(code string, side ClearSide, price float64, amount float64, force bool) OrderParam {
	return OrderParam{
		Side:           string(side),
		Symbol:         code,
		OrderType:      getType(force),
		Price:          price,
		Qty:            amount,
		TimeInForce:    "GoodTillCancel",
		ReduceOnly:     true,
		CloseOnTrigger: false,
	}
}
