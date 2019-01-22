package xviper

import (
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/option"
	"time"
)

import (
	"errors"
	"fmt"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"github.com/xbox1994/xviper/parser"
	"github.com/xbox1994/xviper/reader"
	"github.com/xbox1994/xviper/strategy"
	"os"
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
	case parser.UrlPrefixEtcdv3:
		r = &reader.Etcdv3Reader{ConfigUrl: configUrl}
	case parser.UrlPrefixConsul:
		r = &reader.ConsulReader{ConfigUrl: configUrl}
	default:
		panic(errors.New("impossible run"))
	}
	log.Info.Println("create reader success: " + r.Name())

	if e = r.Init(); e != nil {
		log.Error.Println("init failed")
		return e
	}
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
				log.Info.Println("xviper get change, reread")
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

func Get(key string) interface{}                      { return viper.Get(key) }
func GetString(key string) string                     { return viper.GetString(key) }
func GetBool(key string) bool                         { return viper.GetBool(key) }
func GetInt(key string) int                           { return viper.GetInt(key) }
func GetInt32(key string) int32                       { return viper.GetInt32(key) }
func GetInt64(key string) int64                       { return viper.GetInt64(key) }
func GetFloat64(key string) float64                   { return viper.GetFloat64(key) }
func GetTime(key string) time.Time                    { return viper.GetTime(key) }
func GetDuration(key string) time.Duration            { return viper.GetDuration(key) }
func GetStringSlice(key string) []string              { return viper.GetStringSlice(key) }
func GetStringMap(key string) map[string]interface{}  { return viper.GetStringMap(key) }
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}
func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }
func GetViper() *viper.Viper         { return viper.GetViper() }
