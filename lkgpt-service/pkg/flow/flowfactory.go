package flow

// 类接口
type Class interface {
	Syspromt() string
	Presendgpt() bool
	Returnfromgpt() bool
	Getcontent() string
}

var (
	// 保存注册好的工厂信息
	flowByName = make(map[string]func() Class)
)

// 注册一个类生成工厂
func Register(name string, factory func() Class) {
	flowByName[name] = factory
}

// 根据名称创建对应的类
func Create(name string) Class {
	if f, ok := flowByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
