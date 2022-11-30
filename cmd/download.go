package cmd

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Item",
	Long:  `Download Item`,
	Run:   download,
}

func download(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	checkIfArgsAreNumbers(args)

	url := fmt.Sprintf("https://nyaa.si/download/%s.torrent", args[0])
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to nyaa.si: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing response from nyaa.si: %s\n", err)
		os.Exit(1)
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error accessing user home directory: %s\n", err)
		os.Exit(1)
	}

	path := fmt.Sprintf("%s/Downloads/nyaa-dl/%s", homeDir, params["filename"])

	os.Mkdir(fmt.Sprintf("%s/Downloads/nyaa-dl/", homeDir), os.ModePerm)
	fd, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	_, err = io.Copy(fd, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("downloaded: %s\n", params["filename"])
}

func checkIfArgsAreNumbers(args []string) {
	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			fmt.Fprintf(os.Stderr, "%s is not a valid id\n", arg)
			os.Exit(1)
		}
	}
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
