package authorization

import channel_iso8583 "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/iso8583"

type UrlResponse struct {
	InputURL           string `json:"inputURL"`
	ProcessedTimestamp uint64 `json:"processedTimestamp"`
	CreationDate       string `json:"creationDate"`
	ExpirationDate     string `json:"expirationDate"`
	ShortURL           string `json:"shortURL"`
}

type ErrorResponse struct {
	ProcessedTimestamp uint64 `json:"processedTimestamp"`
	ErrorMessage       string `json:"errorMessage"`
}

type AuthResponse struct {
	Message channel_iso8583.ChannelResponse
}
