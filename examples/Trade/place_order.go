package main

import (
	"context"
	"fmt"

	bybit "github.com/suhostersky/bybit.go.api"
)

func PlaceOrder() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "15").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
