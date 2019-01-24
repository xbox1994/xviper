package reader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/log"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"net/url"
	"path"
	"strings"
	"time"
)

type Etcdv3Reader struct {
	ConfigUrl *url.URL
	kv        clientv3.KV
	watcher   clientv3.Watcher
}

func (this *Etcdv3Reader) Init() error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{this.ConfigUrl.Host},
		DialTimeout: 5 * time.Second,
	})
	this.kv = clientv3.NewKV(client)
	this.watcher = clientv3.NewWatcher(client)
	if err != nil {
		log.Error.Println(fmt.Errorf("connect etcd server failed"))
		return err
	}
	return nil
}

func (this *Etcdv3Reader) Name() string {
	return Etcdv3
}

func (this *Etcdv3Reader) Read() error {
	response, e := this.kv.Get(context.TODO(), this.ConfigUrl.Path)
	if e != nil {
		log.Error.Println("unable to read etcd config")
		return e
	}
	if response == nil || len(response.Kvs) == 0 {
		e := "cannot found config in etcdv3"
		log.Error.Println(e)
		return errors.New(e)
	}
	viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(this.ConfigUrl.Path)), constant.Point))
	if e = viper.ReadConfig(bytes.NewReader(response.Kvs[0].Value)); e != nil {
		log.Error.Println("read config from etcd failed")
		return e
	}
	return nil
}

func (this *Etcdv3Reader) GetWatchFunc(ctx context.Context) WatchFunc {
	return func(reread chan bool) {
		select {
		case <-ctx.Done():
			log.Info.Println("etcdv3 watch exit")
			return
		default:
			watchRespChan := this.watcher.Watch(context.TODO(), this.ConfigUrl.Path)
			for watchResp := range watchRespChan {
				for _, event := range watchResp.Events {
					if event.Type == mvccpb.PUT {
						log.Info.Println("Config etcdv3 changed")
						reread <- true
					}
				}
			}
		}
	}
}

func (this *Etcdv3Reader) Serialize() error {
	return Serialize(Etcdv3, this.ConfigUrl.Path)
}

func (this *Etcdv3Reader) Deserialize() error {
	return Deserialize(Etcdv3, this.ConfigUrl.Path)
}
