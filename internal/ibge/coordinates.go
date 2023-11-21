package ibge

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type ResponseCoordinates struct {
	IbegeCode string `json:"ibge_code"`
}

var cache = make(map[string][]byte)

func CheckCoordinates(w http.ResponseWriter, r *http.Request) {
	ibegeCode := mux.Vars(r)["ibge_code"]

	if data, ok := cache[ibegeCode]; ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Cache: ", ibegeCode)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}

	fmt.Println("not Cache: ", ibegeCode)

	body, err := GetCoordinatesIBGE(ibegeCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cache[ibegeCode] = body

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetCoordinatesIBGE(ibgeCode string) ([]byte, error) {
	url := "https://servicodados.ibge.gov.br/api/v3/malhas/municipios/" + ibgeCode + "?formato=application/vnd.geo+json"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
