package moexiss

import "net/http"

type (
	MoexClient struct {
		Client *http.Client
	}
)

type (
	MoexHistoryData struct {
		Columns MoexColumnsData `json:"columns"`
	}

	MoexColumnsData struct {
		Columns map[int]string
	}
)
