package gtea

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/moqsien/goutils/pkgs/gtea/bar"
)

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url) // nolint:gosec
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("receiving status of %d for url: %s", resp.StatusCode, url)
	}
	return resp, nil
}

func TestDownload(url string) {
	resp, err := getResponse(url)
	if err != nil {
		fmt.Println("could not get response", err)
		os.Exit(1)
	}
	defer resp.Body.Close() // nolint:errcheck

	// Don't add TUI if the header doesn't include content size
	// it's impossible see progress without total
	if resp.ContentLength <= 0 {
		fmt.Println("can't parse content length, aborting download")
		os.Exit(1)
	}

	filename := filepath.Base(url)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("could not create file:", err)
		os.Exit(1)
	}
	defer file.Close() // nolint:errcheck

	title := fmt.Sprintf("[%s] ", filename)
	dbar := bar.NewDownloadBar(bar.WithDefaultGradient(), bar.WithTitle(title), bar.WithWidth(20))
	dbar.SetTotal(resp.ContentLength)

	go dbar.Copy(resp.Body, file)
	dbar.Run()
	// fmt.Println(err)
}
