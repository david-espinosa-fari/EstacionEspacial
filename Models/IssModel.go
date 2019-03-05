package Models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type StationIss struct {
	Timestamp   int  `json:"timestamp"`
	IssPosition LatLong `json:"iss_position"`
	Vurl string
}
type LatLong struct {
	Latitude   float64 `json:"latitude,string"`
	Longitude  float64 `json:"longitude,string"`
}
func (La *StationIss) GetLatitudIss() float64{
	return La.IssPosition.Latitude
}
func (Lo *StationIss) GetLongitudIss() float64{
	return Lo.IssPosition.Longitude
}
func (StationIss *StationIss) Conect()bool{
	StationIss.Vurl=`http://api.open-notify.org/iss-now.json`
	response, err := http.Get(StationIss.Vurl)
	if err != nil {
		flag = true
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		flag = true
	}
	//fmt.Printf("%v", string(body))
	json.Unmarshal(body, StationIss)
	if err != nil {
		flag = false
	}
	flag = true
	return flag
}

