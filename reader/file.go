package reader

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"github.com/xbox1994/xviper/parser"
	"net/url"
	"path"
	"strings"
)

type FileReader struct {
	ConfigUrl *url.URL
}

func (this *FileReader) Init() error {
	return nil
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
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return nil
}

func (this *FileReader) Serialize() error {
	return Serialize(parser.UrlPrefixFile, this.ConfigUrl.Path)
}

func (this *FileReader) Deserialize() error {
	return Deserialize(parser.UrlPrefixFile, this.ConfigUrl.Path)
}
