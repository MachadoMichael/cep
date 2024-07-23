package cep

import (
	"encoding/json"
	"io"
	"net/http"
)

type InfoCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetInfo(cep string) (InfoCEP, error) {

	url := "https://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.Get(url)
	if err != nil {
		return InfoCEP{}, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return InfoCEP{}, err
	}

	var data InfoCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		return InfoCEP{}, err
	}

	return data, nil

}
