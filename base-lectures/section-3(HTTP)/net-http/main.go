package main

import (
	"fmt"
	"http/httpclient/coincap"
	"log"
	"time"
)

func main() {

	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	// assets, err := coincapClient.GetAssets()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, asset := range assets {
	// 	fmt.Println(asset.Info())
	// }

	bitcoin, err := coincapClient.GetAsset("zilliqa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bitcoin.Info())
}
