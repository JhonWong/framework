package framework

type NewInstance func(...interface{}) (interface{}, error)

// 一个服务提供者需要实现的接口
type ServiceProvider interface {
	//注册服务
	Register(Container) NewInstance
	//实例化服务
	Boot(Container) error
	//是否延迟实例化
	IsDefer() bool
	//获取参数
	Params(Container) []interface{}
	//服务凭证
	Name() string
}
