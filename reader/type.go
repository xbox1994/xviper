package reader

type Reader interface {
	Name() string
	Read() error
	GetWatchFunc() WatchFunc
	Serialize() error
	Deserialize() error
}

type WatchFunc func(reread chan bool)
