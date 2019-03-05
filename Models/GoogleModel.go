package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type GoogleModel struct {
	Latitud  float64
	Longitud float64
	Vurl     string
	Idiomas string `json:"languages"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	Status Stat `json:"status"`
	Ocean Oce `json:"ocean"`
}
//{"status":{"message":"no country code found","value":15}}ContryName:
type Stat struct {
Mesagge string `json:"message"`
Value int `json:"value"`
}
type Oce struct {
	OceanName string `json:"name"`
}
func (Gm *GoogleModel)SetDataToGoogle(varLatitud float64, varLongitud float64){
	Gm.Latitud = varLatitud
	Gm.Longitud = varLongitud
}
func (Gm *GoogleModel) GetDataFromGoogle()(string, string){
//http://api.geonames.org/countryCodeJSON?formatted=true&lat=47.03&lng=10.2&username=demo&style=full
//http://api.geonames.org/oceanJSON?formatted=true&lat=40.78343&lng=-43.96625&username=demo&style=full
	var v_url=`http://api.geonames.org/countryCodeJSON?formatted=true&`
	var prox = fmt.Sprint(`lat=`,float64(Gm.Latitud),`&lng=`,float64(Gm.Longitud),`&username=david.espinosa`)
	var FullVurl = v_url+prox
	Gm.Vurl = FullVurl
	//fmt.Printf("%v\n",FullVurl)
	response, err := http.Get(Gm.Vurl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, Gm)
	if err != nil {
		log.Fatal(err)
	}
	if Gm.Status.Value ==15 {
		var nameOcean string = Gm.GetOcean()
		return nameOcean, Gm.Status.Mesagge
	}else{
		return Gm.CountryName, Gm.Status.Mesagge
	}
}
func (Gm* GoogleModel)GetOcean() string {
	var v_url=`http://api.geonames.org/oceanJSON?formatted=true&`
	var prox = fmt.Sprint(`lat=`,float64(Gm.Latitud),`&lng=`,float64(Gm.Longitud),`&username=david.espinosa`)
	var FullVurl = v_url+prox
	Gm.Vurl = FullVurl
	response, err := http.Get(Gm.Vurl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", string(body))
	json.Unmarshal(body, Gm)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Response Value: %s\n", Gm.Ocean.OceanName)
	return Gm.Ocean.OceanName
}
