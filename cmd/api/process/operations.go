package process

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/authorization"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/capture"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/purchase"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/refund"
	channel_iso8583 "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583"
	"io/ioutil"

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
	response := purchase.BuildPurchaseOkResponse().Message
	fmt.Printf("Response=\n %v \n", string(response.ISO8583Message))
	c.JSON(http.StatusOK, response)
}

func (o Operations) AuthOkResponse(c *gin.Context) {
	response := authorization.BuildAuthOkResponse().Message
	fmt.Printf("Response=\n %v \n", string(response.ISO8583Message))
	c.JSON(http.StatusOK, response)
}

func (o Operations) CaptureOkResponse(c *gin.Context) {
	response := capture.BuildCaptureOkResponse().Message
	fmt.Printf("Response=\n %v \n", string(response.ISO8583Message))
	c.JSON(http.StatusOK, response)
}

func (o Operations) RefundOkResponse(c *gin.Context) {
	response := refund.BuildRefundOkResponse().Message
	fmt.Printf("Response=\n %v \n", string(response.ISO8583Message))
	c.JSON(http.StatusOK, response)
}

func (o Operations) GenericHandler(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	//var request transactions_models.Transaction
	var requestIso channel_iso8583.ChannelMessage

	err = json.Unmarshal(jsonData, &requestIso)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Request=\n %v \n", requestIso)
	evaluate := ""

	if requestIso.ISO8583Message.Fields[0].Content != nil {
		evaluate = requestIso.ISO8583Message.Fields[0].Content.(string)
	}

	switch evaluate {
	case "0200":
		o.PurchaseOkResponse(c)
		break
	case "0100":
		o.AuthOkResponse(c)
		break
	case "0220":
		o.CaptureOkResponse(c)
		break
	case "0420":
		o.RefundOkResponse(c)
		break
	default:
		o.PingHandler(c)
		break
	}

}
