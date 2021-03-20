package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyErr != err {
			fmt.Printf("fetch:reading %s %v\n", url, copyErr)
		}
		fmt.Printf("%v\n", resp.Status)
	}
}
