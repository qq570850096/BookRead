package sysinit

// init 函数只会执行一次,并且运行在 main 函数之前
func init() {
	sysinit()
	dbinit("w")
}
