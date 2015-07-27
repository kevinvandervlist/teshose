package container

type Response struct {
	ResponseConfig interface{}
	ConfigType string
	CallBack func()
}