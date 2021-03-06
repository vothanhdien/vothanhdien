package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type weatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type openWeatherMapResponse struct {
	Weather []weatherInfo `json:"weather"`
}

func main() {
	fmt.Println("starting the application")
	f, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	descrip := readMeDesc() + generateExtention()

	if _, err = f.WriteString(descrip); err != nil {
		panic(err)
	}
}

func generateExtention() string {
	s := "\n### Extension\n"
	wi, err := getWeatherInfo()
	if err == nil {
		s += convertToString(*wi)
	}

	ltUpdate := fmt.Sprintf("\n\n**Last updated: %s**\n", getLocationTime())
	s += ltUpdate

	return s
}

func getWeatherInfo() (*weatherInfo, error) {
	resp, err := http.Get("https://samples.openweathermap.org/data/2.5/weather?id=1566083&appid=02565f166f92185daafde2f0d56d2f4e")
	if err != nil {
		fmt.Printf("The http request fail with err %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(data)
	dataStr := string(data)
	fmt.Println(dataStr)

	var iwmr openWeatherMapResponse
	json.Unmarshal(data, &iwmr)
	return &iwmr.Weather[0], nil
}

func convertToString(wi weatherInfo) string {
	return fmt.Sprintf("This is weather in Ho Chi Minh city where I live in : \n\n![icon.png](http://openweathermap.org/img/w/%s.png) *%s*\n\nDescription: %s\n", wi.Icon, wi.Main, wi.Description)
}

func readMeDesc() string {
	return "### Hi there 👋\n" +
		"I'm Võ Thành Điền\n" +
		"- 🔭 I’m currently working on **VNG Corp**\n" +
		"- 🌱 I’m currently learning **Golang**"
}

func getLocationTime() string {
	tz := "Asia/Ho_Chi_Minh"
	t := time.Now()
	loc, err := time.LoadLocation(tz)
	if err == nil {
		return t.In(loc).Format("2006-01-02 15:04:05")
	}
	return t.Format("2006-01-02 15:04:05")
}
