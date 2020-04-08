package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/getlantern/systray"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"wakago/assets"
	"wakago/wakatime"
)

var rootCmd = &cobra.Command{
	Use:  "wakago",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	fetcher := wakatime.NewApiFetcher(mustGetApiKey())

	systray.Run(onReady(fetcher), onExit)
	return nil
}

func onReady(fetcher *wakatime.ApiFetcher) func() {
	fetch := func(ctx context.Context) {
		d, err := fetcher.FetchTodayCodedTime(ctx)
		if err != nil {
			return
		}
		systray.SetTitle(" " + formatDuration(d))
	}

	return func() {
		systray.SetTemplateIcon(assets.Icon(), assets.Icon())
		mQuit := systray.AddMenuItem("Quit", "Quits this app")
		ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

		fetch(ctx)
		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					cancelFunc()
					os.Exit(0)
				}
			}
		}()

		for range time.NewTicker(30 * time.Second).C {
			fetch(ctx)
		}
	}
}

func formatDuration(d time.Duration) string {
	s := fmt.Sprintf("%v mins", int(d.Minutes())%60)
	if d.Hours() > 0 {
		s = fmt.Sprintf("%v hrs %s", int(d.Hours()), s)
	}
	return s
}

func onExit() {

}

func mustGetApiKey() string {
	k, err := wakatime.GetApiKey()
	if err != nil {
		log.Err(err).Msg("get api_key failed")
		os.Exit(1)
	}
	return k
}

func Execute() error {
	return rootCmd.Execute()
}
