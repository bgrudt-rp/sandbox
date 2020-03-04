package main

import (
	"github.com/radiology-partners/rpx/cmd"
	"github.com/radiology-partners/rpx/pkg/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error("program error")
	}
}