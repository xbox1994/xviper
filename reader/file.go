package reader

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type FileReader struct {
	ConfigUrl *url.URL
}

func (this *FileReader) Name() string {
	return "FileReader"
}

func (this *FileReader) Read() error {
	viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(this.ConfigUrl.Path)), constant.Point))
	viper.SetConfigName(strings.TrimSuffix(path.Base(this.ConfigUrl.Path), path.Ext(path.Base(this.ConfigUrl.Path))))
	if this.ConfigUrl.Host == constant.Point {
		viper.AddConfigPath(constant.Point + path.Dir(this.ConfigUrl.Path))
	} else {
		viper.AddConfigPath(path.Dir(this.ConfigUrl.Path))
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Error.Println(fmt.Errorf("unable to read file config: %v", err))
		return err
	}
	return nil
}

func (this *FileReader) GetWatchFunc() WatchFunc {
	return func(reread chan bool) {
		for {
			reread <- true
			time.Sleep(5 * time.Second)
		}
	}
}

func (this *FileReader) Serialize() error {
	serializePath := "xviper_local" + this.ConfigUrl.Path
	dir := path.Dir(serializePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	return viper.WriteConfigAs(serializePath)
}

func (this *FileReader) Deserialize() error {
	file, e := os.Open("xviper_local" + this.ConfigUrl.Path)
	if e != nil {
		log.Error.Println("deserialize failed, not found file")
		return e
	}
	return viper.ReadConfig(file)
}
