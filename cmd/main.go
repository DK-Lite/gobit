package main

import "gobit/bybit"

func main() {
	c := bybit.NewClient("", "")
	c.GetCandlePlus()

}
