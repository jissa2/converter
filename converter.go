package main

import (
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	origem, err := os.Open("recebidos/teste.txt") //Abre arquivo Origem
	if err != nil {
		panic(err)
	}
	converte, err := os.Create("convertidos/teste_utf8.txt") //Cria Arquivo que vai receber a conversão
	if err != nil {
		panic(err)
	}

	executaconversao := charmap.ISO8859_1.NewDecoder().Reader(origem) // Converte

	io.Copy(converte, executaconversao)  //Copia conversão

	converte.Close()
	origem.Close()
}
