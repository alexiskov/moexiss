package moexiss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func NewClient() *MoexClient {
	return &MoexClient{Client: &http.Client{}}
}

// Запрос к исс moex данных за указанную дату
// Дату передаем аргументом
// Вторым аргументом передаем длительность паузы между запросами страниц (если есть ограничение на количество запросов на какой-либо отрезок времени)
func (cli *MoexClient) GetStocksByDate(date time.Time, sleepDurationBetweenPageRequest time.Duration) ([]StockData, error) {
	stockData := []StockData{}

	dataShift := ""
	for {
		req, err := http.NewRequest(http.MethodGet, "https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/tqbr/securities.json?date="+date.Format("2006-01-02")+"&securitytypes=3&iss.meta=off"+dataShift, nil)
		if err != nil {
			return nil, fmt.Errorf("new request building error: %w", err)
		}

		resp, err := cli.Client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("request sending error: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("http response status code: %d", resp.StatusCode)
		}

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("moex response body reading to byte error: %w", err)
		}

		moexData := &MoexResponseData{History: MoexHistoryData{Columns: make([]string, 0, 24), Data: make([][]any, 0, 99)}}
		if err = json.Unmarshal(b, &moexData); err != nil {
			return nil, fmt.Errorf("moex response json unmarshaling error: %w", err)
		}

		if moexData.HistoryCursor.Data[0][0] < moexData.HistoryCursor.Data[0][1] {
			dataShift = "&start=" + strconv.Itoa(int(moexData.HistoryCursor.Data[0][0]+moexData.HistoryCursor.Data[0][2]))
		} else {
			break
		}

		for _, md := range moexData.History.Data {
			sd, err := ConvertToStockData(md)
			if err != nil {
				return nil, fmt.Errorf(" moex response to StockData{} convertation error: %w", err)
			}
			stockData = append(stockData, sd)
		}
		time.Sleep(sleepDurationBetweenPageRequest)
	}

	return stockData, nil
}

// Конвертер moex response data ([]any) ---> StockData{}
func ConvertToStockData(md []any) (StockData, error) {
	tradeDate, err := time.Parse("2006-01-02", md[1].(string))
	if err != nil {
		return StockData{}, fmt.Errorf("moexData time parsing error: %w", err)
	}

	var open float64
	if md[6] != nil {
		open = md[6].(float64)
	}

	var low float64
	if md[7] != nil {
		low = md[7].(float64)
	}

	var high float64
	if md[8] != nil {
		high = md[8].(float64)
	}

	var warPrice float64
	if md[10] != nil {
		warPrice = md[10].(float64)
	}

	var close float64
	if md[11] != nil {
		close = md[11].(float64)
	}

	var marketPrice2 float64
	if md[13] != nil {
		marketPrice2 = md[13].(float64)
	}

	var marketPrice3 float64
	if md[14] != nil {
		marketPrice3 = md[14].(float64)
	}

	var admitedQuote float64
	if md[15] != nil {
		admitedQuote = md[15].(float64)
	}

	var admitedValue float64
	if md[18] != nil {
		admitedValue = md[18].(float64)
	}

	var trendClsPR float64
	if md[22] != nil {
		trendClsPR = md[22].(float64)
	}

	return StockData{
		BOARDID:                 md[0].(string),
		TRADEDATE:               tradeDate,
		SHORTNAME:               md[2].(string),
		SECID:                   md[3].(string),
		NUMTRADES:               md[4].(float64),
		VALUE:                   md[5].(float64),
		OPEN:                    open,
		LOW:                     low,
		HIGH:                    high,
		LEGALCLOSEPRICE:         md[9].(float64),
		WAPRICE:                 warPrice,
		CLOSE:                   close,
		VOLUME:                  md[12].(float64),
		MARKETPRICE2:            marketPrice2,
		MARKETPRICE3:            marketPrice3,
		ADMITTEDQUOTE:           admitedQuote,
		MP2VALTRD:               md[16].(float64),
		MARKETPRICE3TRADESVALUE: md[17].(float64),
		ADMITTEDVALUE:           admitedValue,
		WAVAL:                   md[19].(float64),
		TRADINGSESSION:          int32(md[20].(float64)),
		CURRENCYID:              md[21].(string),
		TRENDCLSPR:              trendClsPR,
	}, nil
}
