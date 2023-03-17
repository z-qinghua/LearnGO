// @program:     LearnGo
// @file:        zaptest.go
// @create:      2022-09-23 12:11
// @description:

package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	logger.Warn("warning test")
}
