package bybit

type API interface {
	// Public API
	GetServerTime() int64
	GetCandle(code string, unit, amount int) ([]Candle, error)
	GetCandlePlus(code string, unit, amount int) ([]Candle, error)

	// Private API with Auth
	GetBalance(coin string) (*Balance, error)
	GetPosition(code string) ([]Position, error)

	// Active Order with Auth
	PlaceClear(code string, position ClearSide, price, amount float64, force bool) (*Order, error)
	PlaceBetting(code string, position BettingSide, price, amount float64, force bool) (*Order, error)
	Replace(code, id string, price, amount float64) (*ReplaceOrder, error)
	Cancel(code, id string) (*CancelOrder, error)
	Search(code string) ([]ActiveOrder, error)
}
