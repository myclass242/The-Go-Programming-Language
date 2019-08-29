package main

import (
	"io"
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		_, err1 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "%v\n", err)
		//	os.Exit(1)
		//}
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err1)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
