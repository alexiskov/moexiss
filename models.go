package moexiss

import "net/http"

type (
	MoexClient struct {
		Client *http.Client
	}
)

type (
	MoexResponseData struct {
		History MoexHistoryData `json:"history"`
	}

	MoexHistoryData struct {
		Columns []string `json:"columns"`
	}
)
