package main

import (
	"fmt"
	"github.com/xbox1994/xviper"
	"time"
)

func main() {
	// config.json for test:
	// {
	//  "GIN_MODE":"1"
	// }

	e := xviper.Init(&xviper.Option{
		// if read failed once, you can set how to handle
		Strategy: &xviper.ReadFailedStrategy{
			Type:          xviper.Retry, // Once/Retry/LoadLast
			RetryTimes:    5,            // 0 if use Once
			RetryInterval: 2,            // 0 if use Once
		},
		// use config url by passing ConfigUrl or environment variable:
		// XVIPER_CONFIG_URL=file://./config/config.json
		// XVIPER_CONFIG_URL=etcdv3://10.13.89.40:2379/config.json
		// XVIPER_CONFIG_URL=consul://localhost:8500/config.json
		ConfigUrl: "etcdv3://10.13.89.40:2379/config.json",
		// watch and reload config automatically if need watch
		NeedWatch: true,
	})
	xviper.SetRemoteWatchHandler(func(updatedValue string) {
		fmt.Println("updateValue: " + updatedValue)
	})

	if e != nil {
		panic(e)
	}
	//xviper.Reset()
	for {
		fmt.Println("Get value from xviper every second for test, you can update config to get latest value: " +
			xviper.GetString("GIN_MODE"))
		time.Sleep(time.Second)
	}
}
