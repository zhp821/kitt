package flow

type EnglishTextFlow struct {
}

func (c *EnglishTextFlow) NewTextpromt() string {
	return "生成一个关于提行李坐飞机场景的英语文章,直接生成内容，不要任何解释描述"
}

func (c *EnglishTextFlow) Syspromt() string {
	return `你是一个非常幽默的英语老师，你正在教{name}学习雅思考试中的一个检票登飞机下飞机场景的英语对话或文章,
	你先打声招呼,然后用英文简单介绍了学习内容，接着输出第一句英文，你依次介绍了这句的重要英文词汇，
	这些重要词汇怎么拼写，中文怎么解释,最后以Please repeat结束,让学生重复一次`
}
func (c *EnglishTextFlow) Presendgpt() bool {
	return true
}
func (c *EnglishTextFlow) Returnfromgpt() bool {
	return true
}
func (c *EnglishTextFlow) Getcontent() string {
	return " "
}

func init() {
	// 在启动时注册类1工厂
	Register("EnglishTextFlow", func() Class {
		return new(EnglishTextFlow)
	})
}
