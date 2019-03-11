package reader

import (
	"context"
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"os"
	"path"
)

const (
	File   = "file"
	Etcdv3 = "etcdv3"
	Consul = "consul"
)

type Reader interface {
	Init() error
	Name() string
	Read() error
	GetWatchFunc(context.Context) WatchFunc
	Serialize() error
	Deserialize() error
}

type WatchFunc func(reread chan string)

func Serialize(readerType string, configPath string) error {
	serializePath := constant.SerializeFolderName + constant.UrlSeparator + readerType + configPath
	dir := path.Dir(serializePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	return viper.WriteConfigAs(serializePath)
}

func Deserialize(readerType string, configPath string) error {
	file, e := os.Open(constant.SerializeFolderName + constant.UrlSeparator + readerType + configPath)
	if e != nil {
		log.Error.Println("deserialize failed, not found file")
		return e
	}
	return viper.ReadConfig(file)
}
