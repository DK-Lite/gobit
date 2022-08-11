package main

import "github.com/DK-Lite/gobit/bybit"

func main() {
	c := bybit.NewClient("", "")
	c.GetCandlePlus()

}
