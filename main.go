package main

import (
	"encoding/json"
	"fmt"
	"github.com/getlantern/systray"
	"io/ioutil"
	"net/http"
	"time"
)

const apiUrl = "https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5"

var httpClient = &http.Client{Timeout: 10 * time.Second}

type curRate struct {
	Ccy string `json:"ccy"`
	Buy string `json:"buy"`
}

func main() {
	systray.Run(onReady, onClose)
}

func onReady() {
	systray.SetIcon(getIcon("assets/icon.png"))
	for {
		time.Sleep(30 * time.Minute)
		systray.SetTitle(checkRate())
	}
}

func onClose() {

}

func checkRate() string {
	req, err := httpClient.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error fetching data: %s", err)
	}
	defer req.Body.Close()
	var res []curRate
	if err := json.NewDecoder(req.Body).Decode(&res); err != nil {
		fmt.Println(err)
	}
	for _, v := range res {
		if v.Ccy == "USD" {
			return v.Buy
		}
	}
	return "NaN"
}

func getIcon(s string) []byte {
	r, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Printf("Error reading: %s", err)
	}
	return r
}
