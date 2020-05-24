package demo6_1

import "testing"

func TestStrategy_Do(t *testing.T) {
	context := Context{}
	context.Strategy = &Strategy1{}
	context.Do()

	context.Strategy = &Strategy2{}
	context.Do()
}
