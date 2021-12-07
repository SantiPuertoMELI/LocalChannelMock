package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/process"

	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}

func run(port string) error {
	// DefaultMeliRouter includes newrelic, datadog, attributes filter, jsonp and pprof middlewares.
	router := mlhandlers.DefaultMeliRouter()

	operations := process.Operations{}

	mapRoutes(router, operations)

	return router.Run(":" + port)
}

func mapRoutes(r *gin.Engine, operations process.Operations) {
	r.GET("/ping", operations.PingHandler)
	//r.POST("/v1/message", operations.PurchaseOkResponse)
	r.POST("/v1/message", operations.CaptureOkResponse)
}
