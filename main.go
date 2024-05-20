package main

import (
	"fmt"
)

type Carro struct {
	Marca  string
	Modelo string
	Ano    string
}

func main() {
	carro1 := Carro{
		Marca:  "Ford",
		Modelo: "Mustang",
		Ano:    "1969",
	}
	carro2 := Carro{
		Marca:  "Fiat",
		Modelo: "Uno",
		Ano:    "1999",
	}
	carro3 := Carro{
		Marca:  "Chevrolet",
		Modelo: "Celta",
		Ano:    "2000",
	}
	fmt.Println(carro1, carro2, carro3)
}
