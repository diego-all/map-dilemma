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
}

func readConfigMetadata(configFile string) (string, [][]string, error) {
	jsonData, err := os.Open(configFile)
	if err != nil {
		return "", nil, err
	}
	defer jsonData.Close()

	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return "", nil, err
	}

	var tipos []models.Tipo
	if err := json.Unmarshal(bytes, &tipos); err != nil {
		return "", nil, err
	}

	var class string
	var classMetadata [][]string

	// Iterar sobre cada tipo y sus atributos
	for _, tipo := range tipos {
		class = tipo.Tipo
		fmt.Println("Clase:", tipo.Tipo)
		fmt.Println("Atributos:")

		// Slice temporal para mantener el orden de los atributos
		var atributosOrdenados []struct {
			Key   string
			Value string
		}

		// Decodificar los atributos manteniendo el orden
		rawMessage := json.RawMessage{}
		json.Unmarshal(bytes, &rawMessage)

		var tmp []map[string]json.RawMessage
		json.Unmarshal(rawMessage, &tmp)

		for _, t := range tmp {
			var atributos map[string]json.RawMessage
			json.Unmarshal(t["atributos"], &atributos)

			for key, value := range atributos {
				var atributo models.Atributo
				json.Unmarshal(value, &atributo)
				atributosOrdenados = append(atributosOrdenados, struct {
					Key   string
					Value string
				}{Key: key, Value: atributo.TipoDato})
			}
		}

		for _, atributo := range atributosOrdenados {
			fmt.Printf(" - %s: %s\n", atributo.Key, atributo.Value)
			classMetadata = append(classMetadata, []string{atributo.Key, atributo.Value})
		}

		// Solo procesar el primer tipo en el JSON
		break
	}

	return class, classMetadata, nil
}
