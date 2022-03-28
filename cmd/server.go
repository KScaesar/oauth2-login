package main

import (
	"github.com/KScaesar/goutils/xLog"
	"github.com/gin-gonic/gin"
)

func main() {
	err := gin.Default().Run(":5001")
	if err != nil {
		xLog.Err(err).Send()
		return
	}
}
