package cmd

import (
	"fmt"
	"os"

	"github.com/ej-agas/nyaa-dl/domain"
	"github.com/ej-agas/nyaa-dl/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Item",
	Long:  `Search for Anime, Audio, Live Action, Pictures, Software`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	nyaa := http.CreateNyaa("https://nyaa.si")
	query := domain.NewQuery(domain.NoFilter, domain.AllCategories, args[0])
	res, err := nyaa.Search(*query)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items := http.SerializeToItems(res)
	renderView(items)
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

	searchCmd.Flags().String(
		"filter",
		"",
		`Available Filters:
no-filter (Default)
no-remakes
trusted-only
	`)

	searchCmd.Flags().String(
		"sort",
		"",
		`Sort By:
title
size
seeders
leechers
downloads
date (Default)
	`)

	searchCmd.Flags().String(
		"order",
		"",
		`Order By:
desc (Default)
asc 
	`)
}
