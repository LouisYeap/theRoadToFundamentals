package user

//package 是go组织源码，源码的集合是代码复用的基础，go内部定义了很多的package，fmt，os，io
//每个源码文件开始都必须申明所属的package
//python中不需要申明package，python有默认采用了文件名的关系，通过文件名来访问
//代码规范
//go语言单元测试 go test是按照一定约定和组织驱动程序，在包目录中，所有有_test.go为后缀的源码文件都会按照go test运行到
//我们写的_test.go不用担心过多，不会被打包

func add(a, b int) int {
	return a + b
}
