package main

import (
	"context"
	"fmt"

	bybit "github.com/suhostersky/bybit.go.api"
)

func GetPositionList() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "settleCoin": "USDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPositionList(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
