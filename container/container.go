package container

type Response struct {
	ResponseConfig interface{}
	ConfigType string
	NoOp bool
	CallBack func()
}