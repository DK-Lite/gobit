gobit is Golang Package for bybit exchang api

## Installation
```cmd
go get github.com/DK-Lite/gobit
```

## project (go.mod)
add package
```cmd
require (
 ...
 github.com/DK-Lite/gobit v1.1.0
)

```

set go env
```cmd
go env -w GOPRIVATE="github.com/DK-Lite/*"
go mod tidy
```

## Get Started (Example)
```go
package main

import (
	"fmt"

	"github.com/DK-Lite/gobit/bybit"
)

func main() {
	client := bybit.NewClient(bybit.Options{
		AccessKey: "",
		SecretKey: "",
	})

	candle, err := client.GetCandle("BTCUSDT", 3, 10)
	if err != nil {
		return
	}

	fmt.Printf("Response: %+v", candle)
}
```

## API
```go
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
```

## Testing
```cmd
go test ./bybit -v
```
