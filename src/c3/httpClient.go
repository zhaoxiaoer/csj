package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://192.168.1.7:8080/csj/c3/servlettemplate", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", req)
	fmt.Printf("%s\n", req.Header.Get("User-Agent"))
	//	req.Header.Set("User-Agent", "zhao")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%#v\n", resp)
}
