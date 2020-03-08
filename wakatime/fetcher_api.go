package wakatime

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"time"
	"wakago/model"
)

var _ Fetcher = &ApiFetcher{}

type ApiFetcher struct {
	client *resty.Client
}

type summariesResponse struct {
	Data []model.SummaryData `json:"data"`
}

func (f *ApiFetcher) FetchTodayCodedTime(ctx context.Context) (time.Duration, error) {
	start := time.Now().Format("2006-01-02")
	end := start

	resp, err := f.client.R().
		SetQueryParam("start", start).
		SetQueryParam("end", end).
		SetResult(&summariesResponse{}).
		Get("https://wakatime.com/api/v1/users/current/summaries")
	if err != nil {
		return 0, err
	}
	if resp.IsError() {
		return 0, errors.New(resp.Status())
	}

	res := resp.Result().(*summariesResponse)
	if len(res.Data) == 0 {
		return 0, errors.New("invalid response")
	}

	ms := time.Duration(res.Data[0].GrandTotal.TotalSeconds * 1000)
	return ms * time.Millisecond, nil
}

func NewApiFetcher(apiKey string) *ApiFetcher {
	return &ApiFetcher{
		client: resty.New().SetBasicAuth(apiKey, ""),
	}
}
