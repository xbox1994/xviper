package reader

type FileReader struct {
}

func (this *FileReader) Name() string {
	return "FileReader"
}

func (this *FileReader) Read() error {
	panic("implement me")
}

func (this *FileReader) GetWatchFunc() (WatchFunc, error) {
	panic("implement me")
}

func (this *FileReader) Serialize() error {
	panic("implement me")
}

