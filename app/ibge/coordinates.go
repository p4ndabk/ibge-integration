package ibge

import (
	"io"
	"net/http"
)

type ResponseCoordinates struct {
	IbegeCode string `json:"ibge_code"`
}

var cache = make(map[string][]byte)

func GetCoordinatesIBGE(ibgeCode string) ([]byte, error) {
	if data, ok := cache[ibgeCode]; ok {
		return data, nil
	}

	url := "https://servicodados.ibge.gov.br/api/v3/malhas/municipios/" + ibgeCode + "?formato=application/vnd.geo+json"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	cache[ibgeCode] = body

	return body, nil
}
