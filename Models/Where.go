package Models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type StationWhere struct {
	Timestamp   int  `json:"timestamp"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Vurl string
}
var flag bool = false
func (La *StationWhere) GetLatitud() float64{
	return La.Latitude
}
func (Lo *StationWhere) GetLongitud() float64{
	return Lo.Longitude
}
func (SW *StationWhere) Conect()bool{
	SW.Vurl=`https://api.wheretheiss.at/v1/satellites/25544`
	response, err := http.Get(SW.Vurl)
	if err != nil {
		//log.Fatal(err)
		return flag
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//log.Fatal(err)
			return flag
	}
	json.Unmarshal(body, SW)
	if err != nil {
		//log.Fatal(err)
		return flag
	}
	flag = true
	return flag
}