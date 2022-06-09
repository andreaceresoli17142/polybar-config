package main

import (
	"fmt"
	// "encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {

	resp, err := http.Get("https://ipinfo.io/loc")

	if err != nil {
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

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/air_pollution?lat=%s&lon=%s&appid=4c1ae3c1a2cbd1c3ff74c66b9305557a", lat, lon)

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

	// fmt.Println(string(body))

	pm10 := gjson.Get(string(body), "list.0.components.pm10").Float()

	/*
		  Good	   1 0-25
		  Fair	   2 25-50
		  Moderate	3 50-90
		  Poor	   4 90-180
		  Very Poor 5 >180
	*/

	retStr := "" 

	if pm10 <= 25 {
		retStr = "💤"
	} else if pm10 <= 50 {
		retStr = "✅"
	} else if pm10 <= 90 {
		retStr = "⚠️ "
	} else if pm10 <= 180 {
		retStr = "🔴"
	} else {
		retStr = "🔥"
	}

	fmt.Printf("%s %v",retStr, pm10)
}
