package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//Coord struct
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

//CityInfo struct
type CityInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Coord   Coord  `json:"coord"`
}

func readCities() []CityInfo {
	filePath := "assets/citylist.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	var cities []CityInfo
	jsn := json.Unmarshal(file, &cities)
	if jsn != nil {
		log.Print("Error in json handling: ", jsn)
	}
	return cities
}

//Cities func
func Cities(str string) JSON {
	ci := readCities()
	str = strings.ToLower(str)
	found := false
	var data JSON
	for x := range ci {
		if strings.Contains(str, strings.ToLower(ci[x].Name)) {
			found = true
			data = GetWeather(str)
		}
	}
	if found == false {
		fmt.Println("Error, city entered not available/valid")

	}
	return data
}
