package logsdk

// Options 日志可选配置
type Options struct {
	LogFormat    string       // 日志格式，仅json和text，若不传默认为json格式输出日志
	Module       string       // 日志模块，主要是使用模块
	ToFieldsFunc ToFieldsFunc // 注入函数，从ctx中获取字段整合fields
}
