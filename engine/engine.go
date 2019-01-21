package engine

import (
	"errors"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"github.com/xbox1994/xviper/parser"
	"github.com/xbox1994/xviper/reader"
	"os"
)

func Init() {
	configUrl, e := parser.Parse(os.Getenv(constant.UrlEnvVarName))
	if e != nil {
		log.Error.Println("parse config url failed")
		panic(e)
	}

	var r reader.Reader
	switch configUrl.Scheme {
	case constant.UrlPrefixFile:
		r = &reader.FileReader{}
	case constant.UrlPrefixEtcd:
		r = &reader.EtcdReader{}
	case constant.UrlPrefixConsul:
		r = &reader.ConsulReader{}
	default:
		panic(errors.New("impossible run"))
	}
	log.Info.Println("create reader success: " + r.Name())

	if e = r.Read(); e != nil {
		log.Error.Println("read failed")
		panic(e)
	}
	if e = r.Serialize(); e != nil {
		log.Error.Println("serialize failed")
		panic(e)
	}

	reread := make(chan bool)
	if watchFunc, e := r.GetWatchFunc(); e != nil {
		log.Error.Println("get watch func failed")
		panic(e)
	} else if watchFunc == nil {
		log.Error.Println("watch func is empty")
		panic(e)
	} else {
		go watchFunc(reread)
	}

	go func() {
		for {
			<-reread
			r.Read()
		}
	}()
}
