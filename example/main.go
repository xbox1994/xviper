package main

import (
	"fmt"
	"github.com/xbox1994/xviper"
	"github.com/xbox1994/xviper/option"
	"github.com/xbox1994/xviper/strategy"
	"time"
)

func main() {
	e := xviper.Init(&option.Option{
		Strategy: &strategy.ReadFailedStrategy{
			Type:          strategy.LoadLast,
			RetryTimes:    5,
			RetryInterval: 2,
		},
	})
	fmt.Println(e)
	for {
		fmt.Println(xviper.Get("GIN_MODE"))
		time.Sleep(1 * time.Second)
	}
}
