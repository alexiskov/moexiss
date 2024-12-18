package moexiss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewClient() *MoexClient {
	return &MoexClient{Client: &http.Client{}}
}

// Запрос к исс moex данных за указанную дату
// Дату передаем аргументом
func (cli *MoexClient) GetIssByDate(date time.Time) (*MoexHistoryData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/tqbr/securities.json?date="+date.Format("2006-01-02"), nil)
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("moex response body reading to byte error: %w", err)
	}

	moexData := &MoexHistoryData{Columns: make(map[int]string)}
	if err = json.Unmarshal(b, moexData); err != nil {
		return nil, fmt.Errorf("moex response parsing error: %w", err)
	}

	return moexData, nil
}
