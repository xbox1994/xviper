package reader

import (
	"context"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"net/url"
	"path"
	"strings"
)

type ConsulReader struct {
	ConfigUrl *url.URL
}

func (this *ConsulReader) Init() error {
	return nil
}

func (this *ConsulReader) Name() string {
	return Consul
}

func (this *ConsulReader) Read() error {
	viper.AddRemoteProvider(Consul, this.ConfigUrl.Host, this.ConfigUrl.Path)
	viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(this.ConfigUrl.Path)), constant.Point))
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Error.Printf("Fatal error config file: %s", err)
		return err
	}
	return nil
}

func (this *ConsulReader) GetWatchFunc(ctx context.Context) WatchFunc {
	return func(updatedValue chan string) {
		select {
		case <-ctx.Done():
			log.Info.Printf("consul watch exit")
			return
		default:
			var (
				err    error
				params map[string]interface{}
				plan   *watch.Plan
			)
			params = make(map[string]interface{})
			params["type"] = "key"
			params["key"] = this.ConfigUrl.Path
			plan, err = watch.Parse(params)
			if err != nil {
				log.Error.Printf("consul watch failed: %s", err)
				return
			}
			plan.Handler = func(index uint64, result interface{}) {
				if result == nil {
					log.Error.Printf("consul watch empty result: %s", err)
					return
				}
				v, ok := result.(*consulapi.KVPair)
				if !ok || v == nil {
					log.Error.Printf("consul watch invalid result: %s", err)
					return
				}
				updatedValue <- string(v.Value)
			}
			if err = plan.Run(this.ConfigUrl.Host); err != nil {
				log.Error.Printf("consul watch start failed: %s", err)
				return
			}
		}
	}
}

func (this *ConsulReader) Serialize() error {
	return Serialize(Consul, this.ConfigUrl.Path)
}

func (this *ConsulReader) Deserialize() error {
	return Deserialize(Consul, this.ConfigUrl.Path)
}
