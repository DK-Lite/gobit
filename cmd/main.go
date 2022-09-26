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
