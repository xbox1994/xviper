package reader

import "net/url"

type ConsulReader struct {
	ConfigUrl *url.URL
}

func (this *ConsulReader) Init() error {
	return nil
}

func (this *ConsulReader) Name() string {
	panic("implement me")
}

func (this *ConsulReader) Read() error {
	panic("implement me")
}

func (this *ConsulReader) GetWatchFunc() WatchFunc {
	panic("implement me")
}

func (this *ConsulReader) Serialize() error {
	panic("implement me")
}

func (this *ConsulReader) Deserialize() error {
	panic("implement me")
}
