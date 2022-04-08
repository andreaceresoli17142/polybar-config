package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var iconDict map[string]string = map[string]string{
	"01": "☀️",
	"02": "⛅",
	"03": "🌥️",
	"04": "☁️",
	"09": "🌧️",
	"10": "🌦️",
	"11": "⛈️",
	"13": "❄️",
	"50": "🌫️",
}

type weather struct {
	Icon string `json:"icon"`
}
type tmpData struct {
	Temp float64 `json:"temp"`
}

type respData struct {
	Weather []weather `json:"weather"`
	TmpData tmpData   `json:"main"`
}

func main() {

	resp, err := http.Get("https://ipinfo.io/loc")

	if err != nil {
		fmt.Print("computer offline")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err)
		return
	}

	coordSlice := strings.Split(string(body), ",")

	lat := coordSlice[0]
	lon := coordSlice[1][:len(coordSlice[1])-1]

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=4c1ae3c1a2cbd1c3ff74c66b9305557a&units=metric", lat, lon)

	//url := "https://id.paleo.bg.it/api/v2/user"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 4c1ae3c1a2cbd1c3ff74c66b9305557a")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))

	//var ret map[string]interface{}
	var ret respData
	json.Unmarshal(body, &ret)

	temperature := ret.TmpData.Temp
	icon := iconDict[ret.Weather[0].Icon[:len(ret.Weather[0].Icon)-1]]
	fmt.Printf("%s %v°C", icon, temperature)

}
