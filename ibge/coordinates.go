package ibge

import (
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

type ResponseCoordinates struct {
	IbegeCode string `json:"ibge_code"`
}

func CheckCoordinates(w http.ResponseWriter, r *http.Request) {
	ibegeCode := mux.Vars(r)["ibge_code"]

	body, err := getCoordinatesIBGE(ibegeCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func getCoordinatesIBGE(ibgeCode string) ([]byte, error) {
	url := "https://servicodados.ibge.gov.br/api/v3/malhas/municipios/" + ibgeCode + "?formato=application/vnd.geo+json"
	  
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Leia o corpo da resposta da requisição
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

