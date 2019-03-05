package Views

import (
	"fmt"
)

type View struct {

}
func (v *View) SetView(location string, message string){
				fmt.Printf("\nTenemos la nueva posicion de ISS, se encuentra en %s", location)
				fmt.Printf(", le mostraremos cuando cambie de posicion. \n")
				fmt.Printf("Por favor sea paciente.")
				//fmt.Printf("\nMensaje de la Api: %s\n", message)
}
func (v *View) SetHold(){
	fmt.Printf(".")
}
