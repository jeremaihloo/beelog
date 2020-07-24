package main

import "github.com/jeremaihloo/beelog"

func main() {
	logger := beelog.NewLogger(1000)
	logger.Info("Beelog !")
}
