package demo6_2

import "fmt"

// --------------------------------------------------------------------------------
// 实现一个日志记录器(相当于Context)
// --------------------------------------------------------------------------------

type LogManage struct {
	Logger
}

func NewLogManage(logger Logger) *LogManage {
	return &LogManage{Logger: logger}
}

// --------------------------------------------------------------------------------
// 定义一个抽象的日志
// --------------------------------------------------------------------------------

type Logger interface {
	Info()
	Error()
}

// --------------------------------------------------------------------------------
// 实现具体的日志
// --------------------------------------------------------------------------------

type FileLogManage struct{}

func (f *FileLogManage) Info() {
	fmt.Println("文件日志记录器，级别：Info")
}

func (f *FileLogManage) Error() {
	fmt.Println("文件日志记录器，级别：Error")
}

type DBLogManage struct{}

func (f *DBLogManage) Info() {
	fmt.Println("数据库日志记录器，级别：Info")
}

func (f *DBLogManage) Error() {
	fmt.Println("数据库日志记录器，级别：Error")
}
