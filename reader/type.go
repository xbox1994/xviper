package reader

type Reader interface {
	Name() string
	Read() error
	GetWatchFunc() (WatchFunc, error)
	Serialize() error
}

type WatchFunc func(reread chan bool)
