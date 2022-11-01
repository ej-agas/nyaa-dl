package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ej-agas/nyaa-dl/domain"
	"github.com/ej-agas/nyaa-dl/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Item",
	Long:  `Search for Anime, Audio, Live Action, Pictures, Software`,
	Run:   run,
}

var (
	sortInput  *string
	orderInput *string
	filter     *string
)

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	nyaa := http.CreateNyaa("https://nyaa.si")
	query := domain.NewQuery(domain.NoFilter(), domain.AllCategories, args[0])
	res, err := nyaa.Search(*query)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items := http.SerializeToItems(res)

	sortItems(items)
	renderView(items)
}

func sortItems(items domain.Items) {
	switch strings.ToLower(*sortInput) {
	case "size":
		if strings.ToLower(*orderInput) == "asc" {
			sort.Sort(sort.Reverse(domain.SortBySize{Items: items}))
			break
		}

		sort.Sort(domain.SortBySize{Items: items})
	case "seeders":
		if strings.ToLower(*orderInput) == "asc" {
			sort.Sort(sort.Reverse(domain.SortBySeeders{Items: items}))
			break
		}

		sort.Sort(domain.SortBySeeders{Items: items})
	case "leechers":
		if strings.ToLower(*orderInput) == "asc" {
			sort.Sort(sort.Reverse(domain.SortByLeechers{Items: items}))
			break
		}

		sort.Sort(domain.SortByLeechers{Items: items})
	case "downloads":
		if strings.ToLower(*orderInput) == "asc" {
			sort.Sort(sort.Reverse(domain.SortByDownloads{Items: items}))
			break
		}

		sort.Sort(domain.SortByDownloads{Items: items})
	case "date":
		if strings.ToLower(*orderInput) == "asc" {
			sort.Sort(sort.Reverse(domain.SortByDate{Items: items}))
			break
		}

		sort.Sort(domain.SortByDate{Items: items})
	}
}

func renderView(items domain.Items) {
	itemsCount := len(items.Items)

	tw := table.NewWriter()
	tw.SetStyle(table.StyleBold)
	tw.Style().Options.SeparateRows = true
	tw.AppendHeader(
		table.Row{
			"Title",
			"Size",
			"Seeders",
			"Leechers",
			"Downloads",
			"Date Uploaded",
		})

	for i := 0; i < itemsCount; i++ {
		item := items.Items[i]
		time, err := item.PublishDateLocalTz()
		if err != nil {
			fmt.Println(err)
		}

		tw.AppendRow(
			table.Row{
				text.FgGreen.Sprint(item.Title),
				item.Size,
				text.FgHiGreen.Sprint(item.Seeders),
				text.FgRed.Sprint(item.Leechers),
				item.Downloads,
				time,
			})
	}

	fmt.Printf("%s", tw.Render())
}

func init() {
	rootCmd.AddCommand(searchCmd)

	filter = searchCmd.Flags().String(
		"filter",
		"",
		`Available Filters:
no-filter (Default)
no-remakes
trusted-only
	`)

	sortInput = searchCmd.Flags().String(
		"sort",
		"",
		`Sort By:
size
seeders
leechers
downloads
date (Default)
	`)

	orderInput = searchCmd.Flags().String(
		"order",
		"",
		`Order By:
desc (Default)
asc 
	`)
}
