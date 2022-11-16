package utils

import (
	//"bufio"

	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Coords struct {
	Lat string
	Lng string
}

func Read(file string) []Coords {

	var dados []Coords

	csvFile, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		dados = append(dados, Coords{
			Lat: line[0],
			Lng: line[1],
		})
	}
	return dados
}
