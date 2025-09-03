package main

import (
	"context"
	"fmt"

	bybit "github.com/suhostersky/bybit.go.api"
)

func GetLongShortRatio() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "period": "5min"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetLongShortRatio(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
