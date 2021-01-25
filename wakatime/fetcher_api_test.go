package wakatime

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestApiFetcher_FetchTodayCodedTime(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: string("success"), wantErr: false},
	}
	ctx, cc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cc()
	fetcher := newFetcherFromEnv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetcher.FetchTodayCodedTime(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTodayCodedTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Today Coded Time: %v\n", got)
		})
	}
}

func newFetcherFromEnv() *ApiFetcher {
	apiKey := os.Getenv("WAKAGO_API_KEY")
	return NewApiFetcher(apiKey)
}
