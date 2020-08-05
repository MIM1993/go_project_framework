package global

var (
	// Global 全局配置变量
	Global GlobalStruct
)

type GlobalStruct struct {
	ListenAddressHTTP string
}