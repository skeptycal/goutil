// Package http contains utilities for http requests.
package http

import (
	"io"
	"net/http"

	"github.com/skeptycal/goutil/repo/util2/gofile"
)

// DownloadURL - download content from a URL to <filename>
func DownloadURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f := gofile.Create(filename)
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
