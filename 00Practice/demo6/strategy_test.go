// @program:     LearnGo
// @file:        strategy_test.go
// @create:      2023-03-19 21:11
// @description:

package strategy

import "testing"

func TestStrategy_Do(t *testing.T) {
	context := Context{}
	strategy1 := &Strategy1{}
	context.Strategy = strategy1
	context.Do()

	strategy2 := &Strategy2{}
	context.Strategy = strategy2
	context.Do()

}
