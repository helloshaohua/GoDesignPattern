package demo6_2

import "testing"

func TestLogManage(t *testing.T) {
	logManage := NewLogManage(&FileLogManage{})
	logManage.Info()
	logManage.Error()

	logManage.Logger = &DBLogManage{}
	logManage.Info()
	logManage.Error()
}
