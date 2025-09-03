package main

import (
	"context"
	"fmt"

	bybit "github.com/suhostersky/bybit.go.api"
)

func GetAccountInfo() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	accountResult, err := client.NewUtaBybitServiceNoParams().GetAccountInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
