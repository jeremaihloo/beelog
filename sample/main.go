package main

import "github.com/jeremaihloo/beelog"

func main() {
	logger := beelog.NewLogger(1000)
	logger.EnableFuncCallDepth(true)
	logger.SetLogger("console", "")
	logger.SetLogger("file", `{"filename":"test.log","maxsize":10485760,"daily":true,"maxdays":10}`)
	index := 0
	for index < 10000 {
		logger.Info("Beelog !")
		index++
	}
}
