package xviper

type Option struct {
	Strategy  *ReadFailedStrategy
	NeedWatch bool
	ConfigUrl string
}
