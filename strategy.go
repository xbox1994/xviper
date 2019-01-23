package xviper

const (
	Once = iota
	Retry
	LoadLast
)

type ReadFailedStrategy struct {
	Type          int
	RetryTimes    int
	RetryInterval int
}
