package moexiss

import (
	"net/http"
	"time"
)

type (
	MoexClient struct {
		Client *http.Client
	}
)

type (
	MoexResponseData struct {
		History       MoexHistoryData `json:"history"`
		HistoryCursor CursorData      `json:"history.cursor"`
	}

	MoexHistoryData struct {
		Columns []string `json:"columns"`
		Data    [][]any  `json:"data"`
	}

	CursorData struct {
		Columns []string  `json:"colmns"`
		Data    [][]int32 `json:"data"`
	}
)

type (
	StockData struct {
		BOARDID                 string
		TRADEDATE               time.Time
		SHORTNAME               string
		SECID                   string
		NUMTRADES               float64
		VALUE                   float64
		OPEN                    float64
		LOW                     float64
		HIGH                    float64
		LEGALCLOSEPRICE         float64
		WAPRICE                 float64
		CLOSE                   float64
		VOLUME                  float64
		MARKETPRICE2            float64
		MARKETPRICE3            float64
		ADMITTEDQUOTE           float64
		MP2VALTRD               float64
		MARKETPRICE3TRADESVALUE float64
		ADMITTEDVALUE           float64
		WAVAL                   float64
		TRADINGSESSION          int32
		CURRENCYID              string
		TRENDCLSPR              float64
	}
)
