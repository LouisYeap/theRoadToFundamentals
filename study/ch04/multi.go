package main

import "fmt"

//	定义接口

type MyWriter interface {

	// Write 定义一个写的方法
	Write(string) error
}
type MyCloser interface {

	// Close 定义一个关闭文件的方法
	Close() error
}

type writerCloser struct {
	MyWriter //interface也是一种类型，放入写文件的匿名实现。我想放一个写数据库的实现
}

// 写文件结构体
type fileWriter struct {
	filePath string
}

// 写数据库结构体
type databaseWriter struct {
	host string
}

func (fw *fileWriter) Write(s string) error {
	fmt.Println("写文件->", s)
	return nil
}

func (da *databaseWriter) Write(s string) error {
	fmt.Println("写数据库->", s)
	return nil
}

//	定时实现类，实现类不关心，只关心方法有没有，实现重写

func main() {

	//	结构体里面可以放interface，因为interface也是一种类型，具体实现

	//实例化
	var mw MyWriter = &writerCloser{
		//添加具体实现进来
		&databaseWriter{},
	}

	mw.Write("hello")

}
