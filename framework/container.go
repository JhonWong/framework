package framework

// 服务容器
type Container interface {
	//绑定服务提供者provider,存在时替换
	Bind(provider ServiceProvider) error
	//判断是否已绑定
	IsBind(key string) bool

	//根据凭证获取服务
	Make(key string) (interface{}, error)
	//获取服务，不返回error，不存在提供者时panic
	MustMake(key string) interface{}
	//为不同参数提供不同实例
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type JwContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]interface{}
}
