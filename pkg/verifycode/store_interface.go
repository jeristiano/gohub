package verifycode

type Store interface {
	Set(id string, value string) bool
	// 获取验证码
	Get(id string, clear bool) string

	// 检查验证码
	Verify(id, answer string, clear bool) bool
}
