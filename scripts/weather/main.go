package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var iconDict map[string]string = map[string]string{
	"01d": " ",
	"02d": " ",
	"03d": " ",
	"04d": " ",
	"09d": " ",
	"10d": " ",
	"11d": " ",
	"13d": "❄️ ",
	"50d": " ",
	"01n": " ",
	"02n": " ",
	"03n": " ",
	"04n": " ",
	"09n": " ",
	"10n": " ",
	"11n": " ",
	"13n": "❄️ ",
	"50n": " ",
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

var dataCache string = ".weather-cache"

func main() {

	currentTime := time.Now().Unix()

	data, err := ioutil.ReadFile(dataCache)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("File reading error", err)
			return
		}

		ioutil.WriteFile(dataCache, []byte("piedi;0"), 0777)
	}

	sdata := strings.Split(string(data), ";")

	lastCachedTime, _ := strconv.Atoi(sdata[0])

	if int64(lastCachedTime) < currentTime-300 {

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
			// fmt.Printf("Error: %s", err.Error())
			fmt.Println("requests number exceeded")
			return
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer 4c1ae3c1a2cbd1c3ff74c66b9305557a")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			fmt.Println("req limit exceded")
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Print("req limit exceded")
		}
		//fmt.Println(string(body))

		//var ret map[string]interface{}
		var ret respData
		err = json.Unmarshal(body, &ret)

		if err != nil {
			fmt.Print("req limit exceded")
		}

		temperature := ret.TmpData.Temp
		icon := iconDict[ret.Weather[0].Icon]
		retsrt := fmt.Sprintf("%s %v°C", icon, math.Round(temperature))

		fmt.Print(retsrt)

		ioutil.WriteFile(dataCache, []byte(fmt.Sprintf("%v;%v", retsrt, currentTime)), 0777)
		return
	}

	fmt.Print(sdata)
	return

}
