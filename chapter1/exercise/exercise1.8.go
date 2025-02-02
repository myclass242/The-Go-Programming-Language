package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	const urlHead = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlHead) {
			url = urlHead + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Status)
		fmt.Printf("%s", b)
	}
}