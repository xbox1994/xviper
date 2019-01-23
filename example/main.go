package main

import (
	"fmt"
	"github.com/xbox1994/xviper"
	"time"
)

func main() {
	// environment variable:
	// XVIPER_CONFIG_URL=etcdv3://10.13.89.40:2379/config.json
	// XVIPER_CONFIG_URL=consul://localhost:8500/config.json

	// config.json:
	// {
	//  "GIN_MODE":"1"
	// }

	e := xviper.Init(&xviper.Option{
		Strategy: &xviper.ReadFailedStrategy{
			Type:          xviper.LoadLast,
			RetryTimes:    5,
			RetryInterval: 2,
		},
	})
	if e != nil {
		panic(e)
	}
	for {
		fmt.Println("Get value from xviper: " + xviper.GetString("GIN_MODE"))
		time.Sleep(time.Second)
	}
}
