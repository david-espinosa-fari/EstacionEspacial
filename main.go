package main

import (
	"github.com/EstacionEspacial/Controller"
)

//Si desea especificar una Api descomente la variable que sea por favor
func main() {
	var name_api string /*= "ISSNOW"*/ //= "WHEREISS"

		CS := Controller.Station{
			Name: name_api,
		}
		CS.Conect()

	}


