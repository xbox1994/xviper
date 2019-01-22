package reader

type ConsulReader struct {

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

