package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//LongLat json
type LongLat struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

//Weather json
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

//Main json
type Main struct {
	Temp      float64 `json:"temp"`
	Pressure  float64 `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

//Wind json
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

//Rain json
type Rain struct {
	Hour float64 `json:"3h"`
}

//Clouds json
type Clouds struct {
	All int `json:"all"`
}

//Sys json
type Sys struct {
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

//JSON blah
type JSON struct {
	Coord   LongLat   `json:"coord"`
	Weather []Weather `json:"weather"`
	Base    string    `json:"base"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Rain    Rain      `json:"rain"`
	Clouds  Clouds    `json:"clouds"`
	Dt      int       `json:"Dt"`
	Sys     Sys       `json:"sys"`
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Cod     int       `json:"cod"`
}

//GetWeather func
func GetWeather(s string) JSON {
	// todo : Add the option to search city
	url := "http://api.openweathermap.org/data/2.5//weather?q=" + s + "&APPID=" + APIKey

	req, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	} else {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			//todo: fix this up
			fmt.Println("Error at the data reading stage", err)
		}
		var r JSON
		err = json.Unmarshal(data, &r)
		if err != nil {
			panic(err.Error())
		}
		return r
	}
}

//ToCelcius function
func ToCelcius(f float64) float64 {
	return f - 273.15
}

//FloatToString to convert a float number to a string
func FloatToString(input float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input, 'f', 3, 64)
}
