package purchase

import (
	"encoding/json"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/util"
	channel_iso8583 "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583"
	iso8583_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583/models"
	"time"
)

func BuildPurchaseOkResponse() PurchaseResponse {
	return setupPurchaseOkResponse()
}

func setupPurchaseOkFields() map[int]iso8583_models.ISO8583MessageField {
	purchaseISO8583MessageFields := make(map[int]iso8583_models.ISO8583MessageField, 0)
	// MTI
	purchaseISO8583MessageFields[0] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "1200",
	}
	// Processing Code
	purchaseISO8583MessageFields[3] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "000000",
	}
	// Transaction amount
	purchaseISO8583MessageFields[4] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "000000150050",
	}
	// Transmition date and time MMDDhhmmss
	purchaseISO8583MessageFields[7] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("010206150405"),
	}
	// System trace audit number
	purchaseISO8583MessageFields[11] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  string(util.GenerateRandonNumber(100000, 999999)),
	}
	// Local transaction time hhmmss
	purchaseISO8583MessageFields[12] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("150405"),
	}
	// Local transaction date MMDD
	purchaseISO8583MessageFields[13] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("0102"),
	}
	// Settlement date MMDD
	purchaseISO8583MessageFields[15] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("0102"),
	}
	// Capture date MMDD
	purchaseISO8583MessageFields[17] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("0102"),
	}
	// Merchant type / Giro del comercio / SIC Code
	purchaseISO8583MessageFields[18] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "5411",
	}
	// Point of service entry mode
	purchaseISO8583MessageFields[22] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "010",
	}
	// Acquiring institution identification code
	purchaseISO8583MessageFields[32] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "12", // 12 = BBVA
	}
	// Retrieval reference number
	purchaseISO8583MessageFields[37] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "17564264204",
	}
	// Authorization identification response
	purchaseISO8583MessageFields[38] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "123456",
	}
	// Response code
	purchaseISO8583MessageFields[39] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "00",
	}
	// Card acceptor terminal identification
	purchaseISO8583MessageFields[41] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "55555555",
	}
	// Card acceptor name/location
	purchaseISO8583MessageFields[43] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "MERPAGO*COMERCIOAGRE\\Paseo del Ma\\ABC\\MX",
	}
	// Additional data-retailer data
	purchaseISO8583MessageFields[48] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "50NNNY2010000",
	}
	// Transaction currency code
	purchaseISO8583MessageFields[49] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "484",
	}
	// Datos de la terminal
	purchaseISO8583MessageFields[60] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "WALMTES1+0000000",
	}
	// Postal code
	purchaseISO8583MessageFields[62] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "010     27000",
	}
	return purchaseISO8583MessageFields
}

func setupPurchaseOkResponse() PurchaseResponse {
	purchaseISO8583Message := iso8583_models.ISO8583Message{
		Packaging: "iso87B",
		Fields:    setupPurchaseOkFields(),
	}

	rawJson, err := json.Marshal(purchaseISO8583Message)
	if err != nil {
		panic(err)
	}

	purchaseChannelResponse := channel_iso8583.ChannelResponse{
		Success:        true,
		Profile:        "bbva-interredes",
		Elapsed:        1200,
		Error:          "",
		Detail:         "",
		ISO8583Message: rawJson,
		Filters:        nil,
	}

	return PurchaseResponse{
		Message: purchaseChannelResponse,
	}
}
