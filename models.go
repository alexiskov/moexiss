package moexiss

import "net/http"

type (
	MoexClient struct {
		Client *http.Client
	}
)

type (
	MoexHistoryData struct {
		Columns map[int]string `json:"columns"`
	}
)
