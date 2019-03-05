package Controller

import (
	"fmt"
	"github.com/EstacionEspacial/Models"
	"github.com/EstacionEspacial/Views"
	"time"
)
var location string
type ToGoogle interface {
	SetDataToGoogle(varLatitud float64, varLongitud float64)
}
type Station struct {
Name string
}
func (s *Station) conectIss() bool{
	SIss := &Models.StationIss{}
	var flag bool = SIss.Conect()
	var LatToGoogle float64 = SIss.GetLatitudIss()
	var LonToGoogle float64 = SIss.GetLongitudIss()
	Gm := Models.GoogleModel{}
	Gm.SetDataToGoogle(LatToGoogle, LonToGoogle)
	var holdValue, message_from_api string = Gm.GetDataFromGoogle()
	if location != holdValue && holdValue != ""{
		location = holdValue
		dv := &Views.View{}
		dv.SetView(location, message_from_api)
	}else {
		dv := &Views.View{}
		dv.SetHold()
	}
	return flag
}
func (s *Station) conectWhere() bool {
	sw := &Models.StationWhere{}
	var flag bool = sw.Conect() //obteniendo latitud y longitud
	var LatToGoogle float64 = sw.GetLatitud()
	var LonToGoogle float64 = sw.GetLongitud()
	Gm := Models.GoogleModel{}
	Gm.SetDataToGoogle(LatToGoogle, LonToGoogle)//pasando lat y lon
    var holdValue, message_from_api string = Gm.GetDataFromGoogle() //obteniendo nombre de location pais u oceano
	if location != holdValue{ //comprobacion que la locacion que viene no se esta mostrando
		location = holdValue
		dv := &Views.View{}
		dv.SetView(location, message_from_api)
	}else {
		dv := &Views.View{}
		dv.SetHold()
	}
	return flag  // el flag es para el caso que no se haya especificado api y esta de error salte a la la otra
}
func (s *Station)Conect() {
	switch s.Name {
	case "issnow":
	case "ISSNOW":
		fmt.Printf("\nConectando a la Api ISS NOW  \n")
		for {
			if !s.conectIss(){
				s.conectWhere()
			}
			time.Sleep(10 * time.Second)
		}

	case "whereiss":
	case "WHEREISS":
		fmt.Printf("\nConectando a la Api WHERE ISS  \n")
		for {
			if !s.conectWhere(){
			s.conectIss()
			}
			time.Sleep(10 * time.Second)
		}

	default:
		fmt.Printf("\nHemos detectado que NO se ha especificado una Api de igual forma lo resolveremos \n")
		for {
			if s.conectWhere() == false {
				s.conectIss()
			}
			time.Sleep(10 * time.Second)
		}

	}
}
