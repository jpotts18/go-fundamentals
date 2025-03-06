package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Latency int
	Url string
	StatusCode int
}

func GetLatencyWithUrl(url string) (result Result, err error) {
	start := time.Now()
	resp, err := http.Get(url)
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		return Result{}, errors.New("Url Failed")
	}
	result = Result{
		Latency: int(elapsed.Milliseconds()),
		Url: url,
		StatusCode: resp.StatusCode,
	}
	return result, nil
}

func main() {

	result, err := GetLatencyWithUrl("https://mangovoice.com")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Url:", result.Url)
	fmt.Println("Status Code:", result.StatusCode)
	fmt.Println("Latency:", result.Latency)
}
