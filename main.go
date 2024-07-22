package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"map-dilemma/models"
	"os"
)

func main() {

	configFile := "inputs/classes.json" // Ruta al archivo JSON
	class, classMetadata, err := readConfigMetadata(configFile)
	if err != nil {
		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Configuración leída: %+v\n %+v\n", class, classMetadata)

	// class, classMetadata, err := readConfigMetadata(configFile)
	// if err != nil {
	// 	fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
	// 	fmt.Println("la clase es:", class)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Configuración leída: %+v\n %+v\n", class, classMetadata)
}

// func readConfigMetadata(configFile string) (string, map[string]string, [][]string, error) {

func readConfigMetadata(configFile string) (string, map[string]string, error) {

	jsonData, err := os.Open(configFile)
	if err != nil {
		return "", nil, err
	}
	defer jsonData.Close()

	// fmt.Println("JSONDATA ES:", jsonData)

	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return "", nil, err
	}

	var tipos []models.Tipo
	if err := json.Unmarshal(bytes, &tipos); err != nil {
		return "", nil, err
	}

	// PROVISIONAL [Solo 1 Tipo del JSON]
	mapAtributos := make(map[string]string)
	var Class string // Declaración de la variable Class

	// Iterar sobre cada tipo y sus atributos
	for _, tipo := range tipos {
		Class = tipo.Tipo
		fmt.Println("Clase:", tipo.Tipo)
		fmt.Println("Atributos:")
		for nombreAtributo, atributo := range tipo.Atributos {

			fmt.Printf(" - %s: %s\n", nombreAtributo, atributo.TipoDato)

			// PROVISIONAL [Solo 1 Tipo del JSON]
			mapAtributos[nombreAtributo] = atributo.TipoDato
		}

		// PROVISIONAL [Solo 1 Tipo del JSON]
		oneType := true
		if oneType == true {
			break
		}
	}

	// PROVISIONAL [Solo 1 Tipo del JSON]
	fmt.Println("mapAtributos es: ", mapAtributos)

	return Class, mapAtributos, nil
}
