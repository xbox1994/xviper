package engine

import (
	"errors"
	"fmt"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"github.com/xbox1994/xviper/option"
	"github.com/xbox1994/xviper/parser"
	"github.com/xbox1994/xviper/reader"
	"github.com/xbox1994/xviper/strategy"
	"os"
	"time"
)

func Init(opt *option.Option) error {
	configUrl, e := parser.Parse(os.Getenv(constant.UrlEnvVarName))
	if e != nil {
		log.Error.Println("parse config url failed")
		return e
	}

	var r reader.Reader
	switch configUrl.Scheme {
	case parser.UrlPrefixFile:
		r = &reader.FileReader{ConfigUrl: configUrl}
	case parser.UrlPrefixEtcd:
		r = &reader.EtcdReader{}
	case parser.UrlPrefixConsul:
		r = &reader.ConsulReader{}
	default:
		panic(errors.New("impossible run"))
	}
	log.Info.Println("create reader success: " + r.Name())

	if e = r.Read(); e != nil {
		switch opt.Strategy.Type {
		case strategy.Once:
			log.Error.Println("read failed once with strategy: once, done")
			return e
		case strategy.Retry:
			log.Info.Println("read failed once with strategy: retry, try again")
			if e = retry(r.Read, opt.Strategy.RetryTimes, opt.Strategy.RetryInterval); e != nil {
				return e
			}
		case strategy.LoadLast:
			log.Info.Println("read failed once with strategy: loadlast, try again")
			if e = retry(r.Deserialize, opt.Strategy.RetryTimes, opt.Strategy.RetryInterval); e != nil {
				return e
			}
		default:
			panic(errors.New("impossible run"))
		}
	}
	if e = r.Serialize(); e != nil {
		log.Error.Println("serialize failed")
		return e
	}

	reread := make(chan bool)
	if watchFunc := r.GetWatchFunc(); watchFunc == nil {
		log.Info.Println("watch func is empty")
	} else {
		go watchFunc(reread)
		go func() {
			for {
				log.Info.Println("engine watch again")
				if <-reread == true {
					r.Read()
				}
			}
		}()
	}
	log.Info.Println("read successfully")
	return nil
}

type loadFunc func() error

func retry(l loadFunc, times int, interval int) error {
	var i int
	for i = 1; i < times; i++ {
		time.Sleep(time.Duration(interval) * time.Second)
		if e := l(); e != nil {
			log.Info.Println("read failed once, try again")
		} else {
			break
		}
	}
	if i == times {
		e := fmt.Errorf("read failed over max retry times %d, done", times)
		return e
	}
	return nil
}
