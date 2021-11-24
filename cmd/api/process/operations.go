package process

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/purchase"

	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"net/http"
)

type Operations struct{}

// PingHandler returns a successful pong answer to all HTTP requests.
func (o Operations) PingHandler(c *gin.Context) {
	if txn := nrgin.Transaction(c); txn != nil {
		txn.Ignore()
	}

	c.String(http.StatusOK, "pong")
}

func (o Operations) PurchaseOkResponse(c *gin.Context) {
	c.JSON(http.StatusOK, purchase.BuildPurchaseOkResponse().Message)
}