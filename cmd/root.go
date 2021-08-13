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

var (
	menus = make(map[*systray.MenuItem][]*systray.MenuItem)
)

func updateSubMenu(mi *systray.MenuItem, items []string) {
	ar := menus[mi]
	if len(ar) < len(items) {
		for i := len(ar); i < len(items); i++ {
			ar = append(ar, mi.AddSubMenuItem(items[i], items[i]))
		}
	} else if len(ar) > len(items) {
		for i := len(items); i < len(ar); i++ {
			ar[i].Hide()
		}
	}
	for i := range items {
		ar[i].Show()
		ar[i].SetTitle(items[i])
		ar[i].SetTooltip(items[i])
	}
	menus[mi] = ar
}

func onReady(fetcher *wakatime.ApiFetcher) func() {
	var gProjects, gLanguages *systray.MenuItem
	fetch := func(ctx context.Context) {
		d, summaries, err := fetcher.FetchTodaySummaries(ctx)
		if err != nil {
			return
		}
		systray.SetTitle(" " + formatDuration(d))

		var projects []string
		for _, s := range summaries.Projects {
			projects = append(projects, s.Name+" - "+s.Text)
			// gProjects.AddSubMenuItem(s.Name+" - "+s.Text, s.Text)
		}
		updateSubMenu(gProjects, projects)

		var languages []string
		for _, s := range summaries.Languages {
			languages = append(languages, s.Name+" - "+s.Text)
			// gLanguages.AddSubMenuItem(s.Name+" - "+s.Text, s.Text)
		}
		updateSubMenu(gLanguages, languages)
	}

	return func() {
		systray.SetTemplateIcon(assets.Icon(), assets.Icon())
		gProjects = systray.AddMenuItem("Projects", "Projects")
		gLanguages = systray.AddMenuItem("Languages", "Languages")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit wakago")
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
