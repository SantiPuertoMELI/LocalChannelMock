package authorization

import (
	"encoding/json"
	"github.com/mercadolibre/LocalChannelMock/cmd/api/util"
	channel_iso8583 "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583"
	iso8583_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583/models"
	"strconv"
	"time"
)

func BuildAuthOkResponse() AuthResponse {
	return setupPurchaseOkResponse()
}

func setupAuthOkFields() map[int]iso8583_models.ISO8583MessageField {
	authISO8583MessageFields := make(map[int]iso8583_models.ISO8583MessageField, 0)
	// MTI
	authISO8583MessageFields[0] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "0110",
	}
	// Processing Code
	authISO8583MessageFields[3] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "00",
	}
	// Transaction amount
	authISO8583MessageFields[4] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "000000554050",
	}
	// Transmition date and time MMDDhhmmss
	authISO8583MessageFields[7] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("010206150405"),
	}
	// System trace audit number
	authISO8583MessageFields[11] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  strconv.Itoa(util.GenerateRandonNumber(100000, 999999)),
	}
	// Local transaction time hhmmss
	authISO8583MessageFields[12] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("150405"),
	}
	// Local transaction date MMDD
	authISO8583MessageFields[13] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("0102"),
	}
	// Capture date MMDD
	authISO8583MessageFields[17] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  time.Now().Format("0102"),
	}
	// Merchant type / Giro del comercio / SIC Code
	authISO8583MessageFields[18] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "5411",
	}
	// Point of service entry mode
	authISO8583MessageFields[22] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "010",
	}
	// Point of service condition code
	authISO8583MessageFields[25] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "59",
	}
	// Acquiring institution identification code
	authISO8583MessageFields[32] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "12", // 12 = BBVA
	}
	// Retrieval reference number
	authISO8583MessageFields[37] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "7542",
	}
	// Response code
	authISO8583MessageFields[39] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "00",
	}
	// Card acceptor terminal identification
	authISO8583MessageFields[41] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "55555555",
	}
	// Additional data-retailer data
	authISO8583MessageFields[48] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "50NNNY2010000",
	}
	// Transaction currency code
	authISO8583MessageFields[49] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "484",
	}
	// Datos de la terminal
	authISO8583MessageFields[60] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "WALMTES1+0000000",
	}
	// Postal code
	authISO8583MessageFields[62] = iso8583_models.ISO8583MessageField{
		Encoding: iso8583_models.ISO8583FieldTypeString,
		Content:  "010     27000",
	}
	return authISO8583MessageFields
}

func setupPurchaseOkResponse() AuthResponse {
	authISO8583Message := iso8583_models.ISO8583Message{
		Packaging: "iso87B",
		Fields:    setupAuthOkFields(),
	}

	rawJson, err := json.Marshal(authISO8583Message)
	if err != nil {
		panic(err)
	}

	authChannelResponse := channel_iso8583.ChannelResponse{
		Success:        true,
		Profile:        "bbva-interredes",
		Elapsed:        1000,
		Error:          "",
		Detail:         "",
		ISO8583Message: rawJson,
		Filters:        nil,
	}

	return AuthResponse{
		Message: authChannelResponse,
	}
}
